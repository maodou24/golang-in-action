export interface Result<T> {
  code: number
  msg: string
  data: T
  cost?: string
  request_id?: string
}

export interface NamespaceRecord {
  id: number
  name: string
  createdAt?: string
  updatedAt?: string
}

export interface ProjectDependency {
  id?: number
  projectId?: number
  releaseId?: number
  repoUrl: string
  path: string
  name: string
  branchName: string
  groups?: string[]
}

export interface RepoBranchResult {
  id?: number
  releaseId?: number
  repoUrl: string
  repoName?: string
  branchName: string
  result: string
  message: string
}

export interface ReleaseRecord {
  id: number
  projectId: number
  version: string
  baselineId?: number
  branchName: string
  status: string
  message?: string
  manifestCommitSha?: string
  dependencies: ProjectDependency[]
  branchResults?: RepoBranchResult[]
  taskId?: string
  projectName?: string
  releaseName?: string
}

export interface ProjectConfig {
  id: number
  namespaceId: number
  name: string
  namespace: string
  namespaceInfo?: NamespaceRecord
  releases: ReleaseRecord[]
  dependencies: ProjectDependency[]
  versions: string[]
  namespaceName: string
  projectName: string
  baselineRef?: string
  remoteName: string
  fetch: string
  defaultXmlPath?: string
}

export interface ProjectRequest {
  namespaceId?: number
  projectId?: number
  projectName: string
  dependencies: ProjectDependency[]
  commitSummary?: string
  commitDescription?: string
  namespace?: string
}

export interface CreateReleaseRequest {
  projectId?: number
  releaseName: string
  baselineId?: number
  dependencies?: ProjectDependency[]
  commitSummary: string
  commitDescription: string
  namespace: string
  projectName: string
  baselineRef: string
  taskId?: string
  targetBranchName: string
}

export type ReleaseTask = ReleaseRecord

export interface UpsertUserRequest {
  accessToken: string
  email: string
  gitProviderType?: number
}

export interface UserProfile {
  publicId: string
  name: string
  login: string
  email: string
  gitProviderType: number
  githubLogin: string
  displayName: string
}

export interface AuthResponse {
  authenticated: boolean
  user: UserProfile
  userId: string
}