import axios from 'axios'
import type {
  AuthResponse,
  CreateReleaseRequest,
  NamespaceRecord,
  ProjectConfig,
  ProjectDependency,
  ProjectRequest,
  ReleaseRecord,
  Result,
  UpsertUserRequest,
} from '../types/release'

const http = axios.create({
  baseURL: import.meta.env.BASE_URL,
  timeout: 60000,
  withCredentials: true,
})

async function get<T>(url: string): Promise<T> {
  const response = await http.get<Result<T>>(url)
  return response.data.data
}

async function send<T>(method: 'post' | 'put' | 'delete', url: string, payload?: unknown): Promise<T> {
  const response = await http.request<Result<T>>({ method, url, data: payload })
  return response.data.data
}

let namespaces: NamespaceRecord[] = []
let projects: ProjectConfig[] = []

function namespaceByName(name: string): NamespaceRecord {
  const namespace = namespaces.find((item) => item.name === name)
  if (!namespace) throw new Error(`Unknown namespace: ${name}`)
  return namespace
}

function projectByName(name: string, namespaceName?: string): ProjectConfig {
  const project = projects.find((item) => item.name === name
    && (!namespaceName || item.namespaceName === namespaceName))
  if (!project) throw new Error(`Unknown project: ${namespaceName ? `${namespaceName}/` : ''}${name}`)
  return project
}

type ProjectApi = Omit<ProjectConfig, 'namespace' | 'namespaceName' | 'projectName' | 'remoteName' | 'fetch' | 'defaultXmlPath' | 'dependencies' | 'versions' | 'namespaceInfo'> & {
  namespace: NamespaceRecord
  releases?: ReleaseRecord[]
}

function normalizeProject(project: ProjectApi): ProjectConfig {
  const releases = (project.releases ?? []).map(normalizeRelease)
  const defaultRelease = releases.find((release) => release.version === 'default' || release.version === 'dev')
  return {
    ...project,
    namespace: project.namespace.name,
    namespaceInfo: project.namespace,
    namespaceName: project.namespace.name,
    projectName: project.name,
    releases,
    versions: releases.map((release) => release.version),
    dependencies: defaultRelease?.dependencies ?? [],
    baselineRef: defaultRelease?.version ?? 'default',
    remoteName: 'origin',
    fetch: 'https://github.com',
  }
}

function normalizeRelease(release: ReleaseRecord): ReleaseRecord {
  return {
    ...release,
    taskId: String(release.id),
    projectName: projects.find((project) => project.id === release.projectId)?.projectName,
    releaseName: release.version,
    dependencies: (release.dependencies ?? []).map((dependency) => {
      const groups = dependency.groups as unknown
      return {
        ...dependency,
        groups: Array.isArray(groups)
          ? groups as string[]
          : typeof groups === 'string'
            ? groups.split(',').filter(Boolean)
            : [],
      }
    }),
    branchResults: (release.branchResults ?? []).map((item) => ({
      ...item,
      repoName: item.repoName ?? item.repoUrl,
    })),
  }
}

export async function refreshProjects(): Promise<void> {
  if (namespaces.length === 0) await listNamespaces()
  const result = await Promise.all(namespaces.map(async (namespace) => {
    const data = await get<{ items: ProjectApi[] }>(`/api/namespaces/${namespace.id}/projects`)
    return (data.items ?? []).map(normalizeProject)
  }))
  projects = result.flat()
}

export async function createRelease(payload: CreateReleaseRequest): Promise<ReleaseRecord> {
  const project = payload.projectId
    ? projectByID(payload.projectId)
    : projectByName(payload.projectName ?? '', payload.namespace)
  const baseline = payload.baselineId
    ? project.releases.find((release) => release.id === payload.baselineId)
    : project.releases.find((release) => release.version === payload.baselineRef)
  const body = {
    projectId: project.id,
    releaseName: payload.releaseName,
    baselineId: payload.baselineId ?? baseline?.id,
    dependencies: payload.dependencies,
    commitSummary: payload.commitSummary,
    commitDescription: payload.commitDescription,
  }
  const response = normalizeRelease(await send<ReleaseRecord>('post', `/api/projects/${project.id}/releases`, body))
  response.projectName = project.projectName
  return response
}

export async function upsertUser(payload: UpsertUserRequest): Promise<AuthResponse> {
  const response = await send<AuthResponse>('post', '/api/users', {
    ...payload,
    gitProviderType: payload.gitProviderType ?? 1,
  })
  return normalizeAuth(response)
}

export async function getCurrentUser(): Promise<AuthResponse> {
  return normalizeAuth(await get<AuthResponse>('/api/users/me'))
}

export async function createProject(payload: ProjectRequest): Promise<ProjectConfig> {
  const id = payload.namespaceId ?? namespaceByName(payload.namespace ?? '').id
  const response = await send<ProjectApi>('post', `/api/namespaces/${id}/projects`, {
    namespaceId: id,
    projectName: payload.projectName,
    dependencies: payload.dependencies,
  })
  await refreshProjects()
  return normalizeProject(response)
}

export async function updateProject(projectName: string | number, payload: ProjectRequest): Promise<ProjectConfig> {
  const projectID = typeof projectName === 'number'
    ? projectName
    : projectByName(projectName, payload.namespace).id
  const response = await send<ProjectApi>('put', `/api/projects/${projectID}`, {
    projectId: projectID,
    namespaceId: payload.namespaceId ?? namespaceByName(payload.namespace ?? '').id,
    projectName: payload.projectName,
    dependencies: payload.dependencies,
  })
  await refreshProjects()
  return normalizeProject(response)
}

export async function updateProjectVersion(namespaceName: string | number, projectName: string, version: string, payload: Partial<CreateReleaseRequest>): Promise<ReleaseRecord> {
  const releaseID = typeof namespaceName === 'number'
    ? namespaceName
    : getReleaseID(namespaceName, projectName, version)
  const project = typeof namespaceName === 'number' ? undefined : projectByName(projectName, namespaceName)
  const currentRelease = project?.releases.find((item) => item.id === releaseID)
  const baselineID = payload.baselineId ?? (payload.baselineRef && project
    ? project.releases.find((item) => item.version === payload.baselineRef)?.id
    : currentRelease?.baselineId)
  const response = await send<ReleaseRecord>('put', `/api/releases/${releaseID}`, {
    id: releaseID,
    projectId: payload.projectId ?? project?.id,
    releaseName: payload.releaseName ?? version,
    baselineId: baselineID,
    dependencies: payload.dependencies ?? [],
  })
  return normalizeRelease(response)
}

export async function deleteProjectVersion(namespaceName: string | number, projectName?: string, version?: string, _payload?: unknown): Promise<void> {
  const releaseID = typeof namespaceName === 'number'
    ? namespaceName
    : getReleaseID(namespaceName, projectName ?? '', version ?? '')
  await send<void>('delete', `/api/releases/${releaseID}`)
}

export async function deleteProject(namespaceName: string | number, projectName?: string, _payload?: unknown): Promise<void> {
  const projectID = typeof namespaceName === 'number'
    ? namespaceName
    : projectByName(projectName ?? '', namespaceName).id
  await send<void>('delete', `/api/projects/${projectID}`)
  projects = projects.filter((project) => project.id !== projectID)
}

export async function deleteNamespace(namespaceName: string | number, _payload?: unknown): Promise<void> {
  const namespaceID = typeof namespaceName === 'number' ? namespaceName : namespaceByName(namespaceName).id
  await send<void>('delete', `/api/namespaces/${namespaceID}`)
  namespaces = namespaces.filter((namespace) => namespace.id !== namespaceID)
  projects = projects.filter((project) => project.namespaceId !== namespaceID)
}

export async function listProjects(namespaceName: string): Promise<ProjectConfig[]> {
  const namespace = namespaceByName(namespaceName)
  const data = await get<{ items: ProjectApi[] }>(`/api/namespaces/${namespace.id}/projects`)
  const selected = (data.items ?? []).map(normalizeProject)
  projects = [...projects.filter((project) => project.namespaceId !== namespace.id), ...selected]
  return selected
}

export async function listNamespaces(): Promise<string[]> {
  namespaces = await get<NamespaceRecord[]>('/api/namespaces')
  return namespaces.map((namespace) => namespace.name)
}

export async function getProjectVersion(namespaceName: string, projectName: string, version: string): Promise<ProjectConfig> {
  const project = projectByName(projectName, namespaceName)
  if (!project.releases.some((item) => item.version === version)) await listProjectVersions(namespaceName, projectName)
  const release = project.releases.find((item) => item.version === version)
  if (!release) throw new Error(`Unknown release: ${version}`)
  const detail = normalizeRelease(await get<ReleaseRecord>(`/api/releases/${release.id}`))
  const baselineVersion = detail.baselineId
    ? project.releases.find((item) => item.id === detail.baselineId)?.version ?? String(detail.baselineId)
    : ''
  return {
    ...project,
    baselineRef: baselineVersion,
    dependencies: detail.dependencies,
  }
}

export async function listProjectVersions(namespaceName: string, projectName: string): Promise<string[]> {
  const project = projectByName(projectName, namespaceName)
  const data = await get<{ items: ReleaseRecord[] }>(`/api/projects/${project.id}/releases`)
  project.releases = (data.items ?? []).map(normalizeRelease)
  project.versions = project.releases.map((release) => release.version)
  return project.versions
}

function projectByID(id: number): ProjectConfig {
  const project = projects.find((item) => item.id === id)
  if (!project) throw new Error(`Unknown project ID: ${id}`)
  return project
}

function getReleaseID(namespaceName: string, projectName: string, version: string): number {
  const release = projectByName(projectName, namespaceName).releases.find((item) => item.version === version)
  if (!release) throw new Error(`Unknown release: ${version}`)
  return release.id
}

export function resolveNamespaceID(name: string): number {
  return namespaceByName(name).id
}

export function resolveProject(name: string, namespaceName?: string): ProjectConfig {
  return projectByName(name, namespaceName)
}

function normalizeAuth(auth: AuthResponse): AuthResponse {
  return {
    ...auth,
    userId: auth.user.publicId,
    user: {
      ...auth.user,
      githubLogin: auth.user.login,
      displayName: auth.user.name,
    },
  }
}

export function groupsAsArray(groups: ProjectDependency['groups'] | string): string[] {
  if (Array.isArray(groups)) return groups
  return groups ? groups.split(',').filter(Boolean) : []
}