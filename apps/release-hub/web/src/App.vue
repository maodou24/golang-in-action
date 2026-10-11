<!-- eslint-disable vue/no-v-model-argument -->
<template>
  <a-layout class="shell">
    <a-spin
      v-if="projectCommitSaving"
      class="operation-loading"
      size="large"
      tip="Processing..."
    />
    <a-layout-sider
      :collapsed="collapsed"
      @update:collapsed="handleSiderCollapsedChange"
      collapsible
      :trigger="null"
      theme="dark"
      :width="220"
      :style="siderStyle"
    >
      <div class="brand">Release Hub</div>
      <a-menu
        theme="dark"
        mode="inline"
        :selected-keys="selectedMenuKeys"
        :open-keys="openKeys"
        @openChange="handleOpenChange"
      >
        <a-sub-menu key="project-management">
          <template #icon><AppstoreOutlined /></template>
          <template #title>项目管理</template>

          <a-menu-item
            key="project-create"
            @click="selectSection('project-create')"
          >
            新建项目
          </a-menu-item>
          <a-menu-item
            key="project-config"
            @click="selectSection('project-config')"
          >
            配置项目
          </a-menu-item>
        </a-sub-menu>

        <a-sub-menu key="branch-management">
          <template #icon><BranchesOutlined /></template>
          <template #title>分支管理</template>

          <a-menu-item
            key="branch-create"
            @click="selectSection('branch-create')"
          >
            创建分支
          </a-menu-item>
          <a-menu-item
            key="branch-config"
            @click="selectSection('branch-config')"
          >
            配置分支
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>

    <div class="main-header-fixed">
      <a-button type="text" class="sider-toggle-btn" @click="toggleSider">
        <template #icon>
          <MenuUnfoldOutlined v-if="collapsed" />
          <MenuFoldOutlined v-else />
        </template>
      </a-button>
      <div class="main-header-title">Repo Manifest 管理台</div>
      <div class="header-spacer" />
      <NamespaceSettings @changed="onNamespacesChanged" />
      <a-dropdown v-if="savedUser" :trigger="['click']">
        <a-button type="text" class="account-button">
          <a-avatar class="account-avatar">{{ accountInitial }}</a-avatar>
          <span class="account-login">{{ savedUser.githubLogin }}</span>
        </a-button>
        <template #overlay>
          <a-menu>
            <a-menu-item key="account" disabled>
              <div class="account-menu-user">{{ savedUser.displayName || savedUser.githubLogin }}</div>
              <div class="account-menu-login">@{{ savedUser.githubLogin }}</div>
              <div v-if="savedUser.email" class="account-menu-email">{{ savedUser.email }}</div>
            </a-menu-item>
            <a-menu-divider />
            <a-menu-item key="token" @click="settingsOpen = true">
              <GithubOutlined /> 更新 Access Token
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
      <a-button v-else type="primary" class="login-button" @click="settingsOpen = true">
        <template #icon><GithubOutlined /></template>
        连接 GitHub
      </a-button>
    </div>

    <a-layout :style="mainLayoutStyle">
      <a-layout-content class="content">
        <a-row :gutter="24">
          <a-col :xs="24" :xl="24">
            <a-card class="intro-card" :bordered="false">
              <template v-if="activeSection === 'project-create'">
                <h1>Create Project</h1>
                <p>
                  新建项目配置，设置项目名称、Remote
                  URL，并维护依赖库清单，用于初始化项目目录和 default.xml。
                </p>
              </template>
              <template v-else-if="activeSection === 'project-config'">
                <h1>Project Structure Management</h1>
                <p>
                  选择已有项目后更新 Remote URL
                  和依赖库配置，支持持续维护项目结构。
                </p>
              </template>
              <template v-else-if="activeSection === 'branch-create'">
                <h1>Release Branch Orchestration</h1>
                <p>
                  在一个页面内完成 release 参数录入、manifest 变更提交和 GitHub
                  分支同步。
                </p>
              </template>
              <template v-else>
                <h1>Branch Configuration</h1>
                <p>
                  维护分支模板、命名规则、baseline 策略和 Manifest
                  项目映射，为后续自动创建分支提供配置基础。
                </p>
              </template>
            </a-card>

            <a-card
              v-if="activeSection === 'project-create'"
              class="config-card"
              title="新建项目"
              :bordered="false"
            >
              <a-form layout="vertical">
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="Namespace">
                      <a-auto-complete
                        v-model:value="projectCreateForm.namespace"
                        :options="rootNamespaceOptions"
                        placeholder="选择或输入 Namespace"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="项目名称">
                      <a-input
                        :value="projectCreateForm.projectName"
                        @update:value="
                          (value: string) => (projectCreateForm.projectName = value)
                        "
                      />
                    </a-form-item>
                  </a-col>
                </a-row>

                <a-form-item label="Dependencies">
                  <a-table
                    class="dependencies-table"
                    size="small"
                    :pagination="false"
                    :data-source="projectCreateDependencies"
                    :columns="projectDependencyColumns"
                    row-key="id"
                  >
                    <template #footer>
                      <a-button
                        type="dashed"
                        @click="addProjectCreateDependency"
                        >新增依赖</a-button
                      >
                    </template>
                    <template #bodyCell="{ column, record }">
                      <template v-if="column.key === 'repoUrl'">
                        <a-input
                          :value="record.repoUrl"
                          placeholder="Git 仓库链接"
                          @update:value="
                            (value: string) => updateDependencyRepo(record, value)
                          "
                        />
                      </template>
                      <template v-if="column.key === 'path'">
                        <a-input
                          v-model:value="record.path"
                          placeholder="Path"
                        />
                      </template>
                      <template v-if="column.key === 'name'">
                        <a-input
                          v-model:value="record.name"
                          placeholder="Name"
                        />
                      </template>
                      <template v-if="column.key === 'branchName'">
                        <a-input
                          v-model:value="record.branchName"
                          placeholder="Branch"
                        />
                      </template>
                      <template v-if="column.key === 'groups'">
                        <a-select
                          :value="record.groups"
                          mode="tags"
                          :options="groupOptions"
                          placeholder="选择或输入组"
                          @update:value="(value: string[]) => (record.groups = value)"
                        />
                      </template>
                      <template v-if="column.key === 'action'">
                        <a-button
                          danger
                          type="link"
                          @click="removeProjectCreateDependency(record.id)"
                          >删除</a-button
                        >
                      </template>
                    </template>
                  </a-table>
                </a-form-item>

                <a-space>
                  <a-button type="primary" @click="saveNewProject"
                    >新增项目</a-button
                  >
                </a-space>
              </a-form>
            </a-card>

            <a-card
              v-else-if="activeSection === 'project-config'"
              class="config-card"
              title="配置项目"
              :bordered="false"
            >
              <a-form layout="vertical">
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="Namespace">
                      <a-select
                        v-model:value="projectUpdateForm.namespace"
                        :options="rootNamespaceOptions"
                        placeholder="Select namespace"
                        @change="handleProjectNamespaceChange"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="项目名称">
                      <a-select
                        :value="projectUpdateForm.projectName"
                        @update:value="
                          (value: string) => (projectUpdateForm.projectName = value)
                        "
                        show-search
                        placeholder="下拉选择项目"
                        :options="projectUpdateOptions"
                        option-filter-prop="label"
                        @change="onProjectConfigSelect"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>

                <a-form-item label="Dependencies">
                  <a-table
                    class="dependencies-table"
                    size="small"
                    :pagination="false"
                    :data-source="projectUpdateDependencies"
                    :columns="projectDependencyColumns"
                    row-key="id"
                  >
                    <template #footer>
                      <a-button
                        type="dashed"
                        @click="addProjectUpdateDependency"
                        >新增依赖</a-button
                      >
                    </template>
                    <template #bodyCell="{ column, record }">
                      <template v-if="column.key === 'repoUrl'">
                        <a-input
                          :value="record.repoUrl"
                          placeholder="Git 仓库链接"
                          @update:value="
                            (value: string) => updateDependencyRepo(record, value)
                          "
                        />
                      </template>
                      <template v-if="column.key === 'path'">
                        <a-input
                          v-model:value="record.path"
                          placeholder="Path"
                        />
                      </template>
                      <template v-if="column.key === 'name'">
                        <a-input
                          v-model:value="record.name"
                          placeholder="Name"
                        />
                      </template>
                      <template v-if="column.key === 'branchName'">
                        <a-input
                          v-model:value="record.branchName"
                          placeholder="Branch"
                        />
                      </template>
                      <template v-if="column.key === 'groups'">
                        <a-select
                          :value="record.groups"
                          mode="tags"
                          :options="groupOptions"
                          placeholder="选择或输入组"
                          @update:value="(value: string[]) => (record.groups = value)"
                        />
                      </template>
                      <template v-if="column.key === 'action'">
                        <a-button
                          danger
                          type="link"
                          @click="removeProjectUpdateDependency(record.id)"
                          >删除</a-button
                        >
                      </template>
                    </template>
                  </a-table>
                </a-form-item>

                <a-space>
                  <a-button type="primary" @click="saveProjectConfig"
                    >保存配置</a-button
                  >
                  <a-button
                    danger
                    :disabled="!projectUpdateForm.projectName"
                    @click="deleteProjectConfig"
                    >删除项目</a-button
                  >
                </a-space>
              </a-form>
            </a-card>

            <ReleaseForm
              v-else-if="activeSection === 'branch-create'"
              :key="namespaceRevision"
              @created="onCreated"
            />

            <a-card
              v-else
              class="config-card"
              title="分支配置"
              :bordered="false"
            >
              <a-form layout="vertical">
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="Namespace">
                      <a-select
                        v-model:value="branchConfigForm.namespace"
                        :options="namespaceOptions"
                        @change="handleBranchNamespaceChange"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="项目">
                      <a-select
                        v-model:value="branchConfigForm.projectName"
                        :options="branchProjectOptions"
                        @change="handleBranchConfigProjectChange"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="发布版本">
                      <a-select
                        v-model:value="branchConfigForm.version"
                        :options="
                          branchVersions.map((version) => ({
                            label: version,
                            value: version,
                          }))
                        "
                        @change="loadBranchConfigVersion"
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-row :gutter="16">
                  <a-col :xs="24" :md="12">
                    <a-form-item label="基线分支">
                      <a-input
                        :value="branchConfigForm.baselineRef || '-'"
                        disabled
                      />
                    </a-form-item>
                  </a-col>
                </a-row>
                <a-form-item label="Dependencies">
                  <a-table
                    :data-source="branchConfigDependencies"
                    :columns="projectDependencyColumns"
                    row-key="id"
                    :pagination="false"
                  >
                    <template #footer>
                      <a-button type="dashed" @click="addBranchConfigDependency"
                        >新增依赖</a-button
                      >
                    </template>
                    <template #bodyCell="{ column, record }">
                      <a-input
                        v-if="column.key === 'repoUrl'"
                        v-model:value="record.repoUrl"
                      />
                      <a-input
                        v-else-if="column.key === 'path'"
                        v-model:value="record.path"
                      />
                      <a-input
                        v-else-if="column.key === 'name'"
                        v-model:value="record.name"
                      />
                      <a-input
                        v-else-if="column.key === 'branchName'"
                        v-model:value="record.branchName"
                        :disabled="isDefaultBranchVersion"
                      />
                      <a-select
                        v-else-if="column.key === 'groups'"
                        :value="record.groups"
                        mode="tags"
                        :options="groupOptions"
                        placeholder="选择或输入组"
                        @update:value="(value: string[]) => (record.groups = value)"
                      />
                      <a-space v-else-if="column.key === 'action'">
                        <a-button
                          type="link"
                          size="small"
                          :disabled="
                            isDefaultBranchVersion ||
                            !branchConfigForm.version ||
                            record.branchName.trim() === branchConfigForm.version.trim()
                          "
                          @click="replaceBranchConfigDependency(record)"
                          >Replace</a-button
                        >
                        <a-button
                          danger
                          type="link"
                          @click="removeBranchConfigDependency(record.id)"
                          >删除</a-button
                        >
                      </a-space>
                    </template>
                  </a-table>
                </a-form-item>
                <a-space>
                  <a-button
                    type="primary"
                    :disabled="isDefaultBranchVersion"
                    @click="saveBranchConfig"
                    >保存配置</a-button
                  >
                  <a-button
                    danger
                    :disabled="
                      !branchConfigForm.projectName ||
                      !branchConfigForm.version ||
                      isDefaultBranchVersion
                    "
                    @click="deleteBranchConfig"
                    >删除版本</a-button
                  >
                </a-space>
              </a-form>
            </a-card>
          </a-col>
        </a-row>

        <a-card
          v-if="activeSection === 'branch-create' && task"
          class="result-card"
          title="Latest Task"
          :bordered="false"
        >
          <a-descriptions :column="1" bordered size="small">
            <a-descriptions-item label="Task ID">{{
              task.taskId
            }}</a-descriptions-item>
            <a-descriptions-item label="Status">{{
              task.status
            }}</a-descriptions-item>
            <a-descriptions-item label="Project">{{
              task.projectName
            }}</a-descriptions-item>
            <a-descriptions-item label="Release">{{
              task.releaseName
            }}</a-descriptions-item>
            <a-descriptions-item label="Message">{{
              task.message
            }}</a-descriptions-item>
          </a-descriptions>

          <a-table
            class="result-table"
            :pagination="false"
            :data-source="task.branchResults"
            :columns="columns"
            row-key="repoName"
          />
        </a-card>
      </a-layout-content>

      <a-layout-footer class="footer">
        Release Hub · Manifest Driven Release Management
      </a-layout-footer>
    </a-layout>

    <a-modal
      :open="projectCommitModalOpen"
      :title="projectCommitMode === 'delete' ? '确认删除项目' : '提交信息'"
      ok-text="提交"
      cancel-text="取消"
      :confirm-loading="projectCommitSaving"
      @update:open="handleProjectCommitOpenChange"
      @ok="confirmProjectCommit"
      @cancel="cancelProjectCommit"
    >
      <a-form layout="vertical">
        <a-alert
          v-if="projectCommitMode === 'delete'"
          type="warning"
          show-icon
          message="将删除该项目的整个目录，包括 default.xml 和 releases 文件夹。"
          style="margin-bottom: 16px"
        />
        <a-form-item label="Summary" required>
          <a-input
            :value="projectCommitForm.summary"
            @update:value="(value: string) => (projectCommitForm.summary = value)"
          />
        </a-form-item>
        <a-form-item label="Description">
          <a-textarea
            :value="projectCommitForm.description"
            :rows="4"
            placeholder="补充提交描述"
            @update:value="(value: string) => (projectCommitForm.description = value)"
          />
        </a-form-item>
      </a-form>
    </a-modal>

    <a-drawer
      :open="settingsOpen"
      title="GitHub 账户"
      placement="right"
      :width="420"
      @update:open="handleSettingsOpenChange"
    >
      <a-alert
        class="settings-alert"
        type="info"
        show-icon
        message="输入 GitHub Access Token 和提交邮箱完成认证。验证成功后将显示当前 GitHub 用户名。"
      />

      <a-form layout="vertical" :model="userForm" @finish="saveUserSettings">
        <a-form-item
          label="GitHub Access Token"
          name="accessToken"
          :rules="
            savedUser
              ? []
              : [{ required: true, message: '请输入 GitHub Access Token' }]
          "
        >
          <a-input-password
            :value="userForm.accessToken"
            placeholder="github_pat_xxx"
            @update:value="(value: string) => (userForm.accessToken = value)"
          />
        </a-form-item>

        <a-form-item
          label="Email"
          name="email"
          :rules="[
            { required: true, message: '请输入 Email' },
            { type: 'email', message: '请输入有效的 Email' },
          ]"
        >
          <a-input
            :value="userForm.email"
            type="email"
            placeholder="name@jci.com"
            @update:value="(value: string) => (userForm.email = value)"
          />
        </a-form-item>

        <a-space>
          <a-button type="primary" html-type="submit" :loading="settingsSaving"
            >{{ savedUser ? "更新 Token" : "连接 GitHub" }}</a-button
          >
          <a-button @click="settingsOpen = false">Cancel</a-button>
        </a-space>
      </a-form>

      <a-card v-if="savedUser" class="saved-user-card" size="small">
        <a-space align="center">
          <a-avatar class="account-avatar">{{ accountInitial }}</a-avatar>
          <a-descriptions :column="1" size="small">
            <a-descriptions-item label="GitHub Login">{{ savedUser.githubLogin }}</a-descriptions-item>
            <a-descriptions-item v-if="savedUser.email" label="Email">{{ savedUser.email }}</a-descriptions-item>
          </a-descriptions>
        </a-space>
      </a-card>
    </a-drawer>
  </a-layout>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import { message, notification } from "ant-design-vue";
import {
  refreshProjects,
  resolveNamespaceID,
  resolveProject,
  createProject,
  deleteProject,
  deleteProjectVersion,
  getCurrentUser,
  getProjectVersion,
  groupsAsArray,
  listNamespaces,
  listProjects,
  listProjectVersions,
  updateProject,
  updateProjectVersion,
  upsertUser,
} from "./api/client";
import type {
  ProjectConfig,
  ProjectRequest,
  ReleaseTask,
  UpsertUserRequest,
  UserProfile,
} from "./types/release";
import ReleaseForm from "./components/ReleaseForm.vue";
import NamespaceSettings from "./components/NamespaceSettings.vue";
import {
  AppstoreOutlined,
  BranchesOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  GithubOutlined,
} from "@ant-design/icons-vue";

const collapsed = ref(false);
const settingsOpen = ref(false);
const settingsSaving = ref(false);
const namespaceRevision = ref(0);
const task = ref<ReleaseTask | null>(null);
const savedUser = ref<UserProfile | null>(null);
const accountInitial = computed(() => savedUser.value?.githubLogin.charAt(0).toUpperCase() ?? "G");
type ActiveSection =
  | "project-create"
  | "project-config"
  | "branch-create"
  | "branch-config";

const activeSection = ref<ActiveSection>(readActiveSection());
const openKeys = ref<string[]>(["project-management", "branch-management"]);
const projectCommitModalOpen = ref(false);
const projectCommitSaving = ref(false);
const projectCommitMode = ref<
  "create" | "update" | "delete" | "version" | "version-delete"
>("create");
const projectCommitForm = reactive({
  summary: "",
  description: "",
});

interface ProjectDependencyRow {
  id: number;
  repoUrl: string;
  path: string;
  name: string;
  branchName: string;
  groups: string[];
}

interface ProjectTemplate {
  value: string;
  label: string;
  namespace: string;
  projectName: string;
  versions: string[];
  dependencies: Array<{
    repoUrl: string;
    path: string;
    name: string;
    branchName: string;
    groups?: string[];
  }>;
}

const projectUpdateTemplates = ref<ProjectTemplate[]>([]);
const branchConfigTemplates = ref<ProjectTemplate[]>([]);
const rootNamespaces = ref<string[]>([]);

const projectCreateForm = ref({
  namespace: "",
  projectName: "",
});

const projectUpdateForm = ref({
  namespace: "",
  projectName: "",
});

const rootNamespaceOptions = computed(() =>
  rootNamespaces.value.map((value) => ({ label: value, value })),
);
const projectUpdateOptions = computed(() =>
  projectUpdateTemplates.value
    .map((item) => ({
      label: `${item.namespace}/${item.projectName}`,
      value: item.projectName,
    })),
);

const branchConfigForm = ref({
  namespace: "",
  projectName: "",
  version: "",
  baselineRef: "",
});
const branchConfigDependencies = ref<ProjectDependencyRow[]>([]);
const nextBranchConfigDependencyID = ref(1);

const projectCreateDependencies = ref<ProjectDependencyRow[]>([]);
const projectUpdateDependencies = ref<ProjectDependencyRow[]>([]);
const nextCreateDependencyID = ref(1);
const nextUpdateDependencyID = ref(1);

const projectDependencyColumns = [
  { title: "Repository", dataIndex: "repoUrl", key: "repoUrl", width: "28%" },
  { title: "Path", dataIndex: "path", key: "path", width: "14%" },
  { title: "Name", dataIndex: "name", key: "name", width: "18%" },
  { title: "Branch", dataIndex: "branchName", key: "branchName", width: "14%" },
  { title: "Groups", dataIndex: "groups", key: "groups", width: "26%" },
  { title: "操作", dataIndex: "action", key: "action", width: 90 },
];

const groupOptions = computed(() => {
  const dependencies = [
    ...projectCreateDependencies.value,
    ...projectUpdateDependencies.value,
    ...branchConfigDependencies.value,
  ];
  const groups = dependencies.reduce<string[]>(
    (allGroups, dependency) => allGroups.concat(dependency.groups ?? []),
    [],
  );
  return [...new Set(groups)].sort().map((value) => ({ label: value, value }));
});

const userForm = ref<UpsertUserRequest>({
  accessToken: "",
  email: "",
});

const selectedMenuKeys = computed(() => [activeSection.value]);

const columns = [
  { title: "Repo", dataIndex: "repoName", key: "repoName" },
  { title: "Branch", dataIndex: "branchName", key: "branchName" },
  { title: "Result", dataIndex: "result", key: "result" },
  { title: "Message", dataIndex: "message", key: "message" },
];

const siderStyle = computed(() => ({
  position: "fixed",
  left: "0",
  top: "56px",
  height: "calc(100vh - 56px)",
  overflow: "auto",
  zIndex: 1100,
}));

const mainLayoutStyle = computed(() => ({
  marginTop: "56px",
  marginLeft: collapsed.value ? "80px" : "220px",
  minHeight: "calc(100vh - 56px)",
  transition: "margin-left 0.2s ease",
}));

function toggleSider() {
  collapsed.value = !collapsed.value;
}

function handleSiderCollapsedChange(value: boolean) {
  collapsed.value = value;
}

function handleOpenChange(nextOpenKeys: string[]) {
  openKeys.value = nextOpenKeys;
}

function selectSection(section: ActiveSection) {
  activeSection.value = section;
  localStorage.setItem("release_hub_active_section", section);
  const group = section.startsWith("project-")
    ? "project-management"
    : "branch-management";
  if (!openKeys.value.includes(group)) {
    openKeys.value = [...openKeys.value, group];
  }
}

function readActiveSection(): ActiveSection {
  const saved = localStorage.getItem("release_hub_active_section");
  if (
    saved === "project-create" ||
    saved === "project-config" ||
    saved === "branch-create" ||
    saved === "branch-config"
  ) {
    return saved;
  }
  return "project-create";
}

function handleProjectCommitOpenChange(value: boolean) {
  projectCommitModalOpen.value = value;
}

function handleSettingsOpenChange(value: boolean) {
  settingsOpen.value = value;
}

function addProjectCreateDependency() {
  projectCreateDependencies.value.push({
    id: nextCreateDependencyID.value,
    repoUrl: "",
    path: "",
    name: "",
    branchName: "",
    groups: [],
  });
  nextCreateDependencyID.value += 1;
}

function removeProjectCreateDependency(id: number) {
  projectCreateDependencies.value = projectCreateDependencies.value.filter(
    (item) => item.id !== id,
  );
}

function addProjectUpdateDependency() {
  projectUpdateDependencies.value.push({
    id: nextUpdateDependencyID.value,
    repoUrl: "",
    path: "",
    name: "",
    branchName: "",
    groups: [],
  });
  nextUpdateDependencyID.value += 1;
}

function removeProjectUpdateDependency(id: number) {
  projectUpdateDependencies.value = projectUpdateDependencies.value.filter(
    (item) => item.id !== id,
  );
}

const branchProjectOptions = computed(() =>
  branchConfigTemplates.value
    .map((item) => ({
      label: `${item.namespace}/${item.projectName}`,
      value: item.projectName,
    })),
);

const namespaceOptions = computed(() =>
  rootNamespaces.value.map((value) => ({ label: value, value })),
);

const branchVersions = ref<string[]>([]);
const isDefaultBranchVersion = computed(() =>
  ["dev", "default"].includes(branchConfigForm.value.version.trim().toLowerCase()),
);

async function loadBranchConfigVersion(version: string) {
  branchConfigForm.value.version = version;
  branchConfigForm.value.baselineRef = "";
  if (!branchConfigForm.value.projectName) {
    branchConfigDependencies.value = [];
    return;
  }
  try {
    const project = await getProjectVersion(
      branchConfigForm.value.namespace,
      branchConfigForm.value.projectName,
      version,
    );
    branchConfigForm.value.baselineRef = project.baselineRef ?? "";
    branchConfigDependencies.value = project.dependencies.map((dep, index) => ({
      id: index + 1,
      ...dep,
      groups: groupsAsArray(dep.groups),
    }));
  } catch (err: any) {
    branchConfigForm.value.baselineRef = "";
    branchConfigDependencies.value = [];
    showError(err, "Failed to load release version");
  }
  nextBranchConfigDependencyID.value =
    branchConfigDependencies.value.length + 1;
}

function toProjectTemplates(items: ProjectConfig[]): ProjectTemplate[] {
  return items.map((project) => ({
    value: project.projectName,
    label: `${project.namespace}/${project.projectName}`,
    namespace: project.namespace,
    projectName: project.projectName,
    versions: project.versions ?? [],
    dependencies: project.dependencies ?? [],
  }));
}

async function handleProjectNamespaceChange(value: string) {
  projectUpdateForm.value.namespace = value;
  projectUpdateForm.value.projectName = "";
  projectUpdateDependencies.value = [];
  nextUpdateDependencyID.value = 1;
  try {
    projectUpdateTemplates.value = toProjectTemplates(await listProjects(value));
  } catch (err: any) {
    projectUpdateTemplates.value = [];
    showError(err, "Failed to load projects");
  }
}

async function handleBranchConfigProjectChange(value: string) {
  branchConfigForm.value.projectName = value;
  branchVersions.value = [];
  branchConfigForm.value.version = "";
  branchConfigForm.value.baselineRef = "";
  branchConfigDependencies.value = [];
  try {
    const versions = await listProjectVersions(branchConfigForm.value.namespace, value);
    branchVersions.value = versions;
    branchConfigForm.value.version = versions.includes("default") ? "default" : "";
    if (branchConfigForm.value.version) {
      await loadBranchConfigVersion(branchConfigForm.value.version);
    }
  } catch (err: any) {
    showError(err, "Failed to load release versions");
  }
}

async function refreshBranchConfig() {
  const { namespace, projectName, version } = branchConfigForm.value;
  if (!namespace || !projectName) {
    return;
  }

  try {
    const versions = await listProjectVersions(namespace, projectName);
    branchVersions.value = versions;
    const nextVersion = versions.includes(version)
      ? version
      : versions.includes("default")
        ? "default"
        : "";
    branchConfigForm.value.version = nextVersion;
    if (nextVersion) {
      await loadBranchConfigVersion(nextVersion);
    } else {
      branchConfigDependencies.value = [];
      nextBranchConfigDependencyID.value = 1;
    }
  } catch (err: any) {
    branchVersions.value = [];
    branchConfigForm.value.version = "";
    branchConfigForm.value.baselineRef = "";
    branchConfigDependencies.value = [];
    nextBranchConfigDependencyID.value = 1;
    showError(err, "Failed to refresh release versions");
  }
}

async function handleBranchNamespaceChange(value: string) {
  branchConfigForm.value.namespace = value;
  branchConfigForm.value.projectName = "";
  branchConfigForm.value.version = "";
  branchConfigForm.value.baselineRef = "";
  branchConfigDependencies.value = [];
  branchVersions.value = [];
  try {
    branchConfigTemplates.value = toProjectTemplates(await listProjects(value));
  } catch (err: any) {
    branchConfigTemplates.value = [];
    showError(err, "Failed to load projects");
  }
}

function addBranchConfigDependency() {
  branchConfigDependencies.value.push({
    id: nextBranchConfigDependencyID.value++,
    repoUrl: "",
    path: "",
    name: "",
    branchName: "",
    groups: [],
  });
}

function removeBranchConfigDependency(id: number) {
  branchConfigDependencies.value = branchConfigDependencies.value.filter(
    (item) => item.id !== id,
  );
}

function replaceBranchConfigDependency(record: ProjectDependencyRow) {
  const branchName = branchConfigForm.value.version.trim();
  if (branchName) {
    record.branchName = branchName;
  }
}

function resetBranchConfig() {
  branchConfigForm.value = {
    namespace: "",
    projectName: "",
    version: "",
    baselineRef: "",
  };
  branchConfigDependencies.value = [];
  nextBranchConfigDependencyID.value = 1;
}

function saveBranchConfig() {
  if (isDefaultBranchVersion.value) {
    message.warning("dev/default 只能在项目管理中修改");
    return;
  }
  if (!branchConfigForm.value.projectName || !branchConfigForm.value.version) {
    message.warning("请选择项目和版本");
    return;
  }
  projectCommitMode.value = "version";
  projectCommitForm.summary = "";
  projectCommitForm.description = "";
  projectCommitForm.summary = `chore(manifest): update ${branchConfigForm.value.projectName}/${branchConfigForm.value.version}`;
  projectCommitForm.description = "updated by release-hub";
  projectCommitModalOpen.value = true;
}

function deleteBranchConfig() {
  if (
    !branchConfigForm.value.projectName ||
    !branchConfigForm.value.version ||
    isDefaultBranchVersion.value
  ) {
    message.warning("默认分支只能在项目管理中删除");
    return;
  }
  projectCommitMode.value = "version-delete";
  projectCommitForm.summary = "";
  projectCommitForm.description = "";
  projectCommitForm.summary = `chore(manifest): delete ${branchConfigForm.value.projectName}/${branchConfigForm.value.version}`;
  projectCommitForm.description = "delete release manifest and remote branches";
  projectCommitModalOpen.value = true;
}

function parseGitHubRepo(value: string) {
  let repo = value
    .trim()
    .replace(/\.git\/?$/, "")
    .replace(/\/$/, "");
  repo = repo.replace(/^https?:\/\/github\.com\//, "");
  repo = repo.replace(/^ssh:\/\/git@github\.com\//, "");
  repo = repo.replace(/^git@github\.com:/, "");
  const parts = repo.split("/").filter(Boolean);
  if (parts.length !== 2) {
    return null;
  }
  return { name: parts.join("/"), path: parts[1] };
}

function repoPath(value: string) {
  return parseGitHubRepo(value)?.path ?? "";
}

function updateDependencyRepo(record: ProjectDependencyRow, value: string) {
  record.repoUrl = value;
  const parsed = parseGitHubRepo(value);
  if (!parsed) {
    return;
  }
  record.path = parsed.path ?? "";
  record.name = parsed.name;
}

async function onProjectConfigSelect(value: string) {
  projectUpdateForm.value.projectName = value;
  if (!value) {
    projectUpdateDependencies.value = [];
    nextUpdateDependencyID.value = 1;
    return;
  }
  try {
    const project = await getProjectVersion(
      projectUpdateForm.value.namespace,
      value,
      "default",
    );
    projectUpdateDependencies.value = project.dependencies.map((dep, index) => ({
      id: index + 1,
      repoUrl: dep.repoUrl,
      path: dep.path || repoPath(dep.repoUrl),
      name: dep.name || dep.repoUrl,
      branchName: dep.branchName,
      groups: groupsAsArray(dep.groups),
    }));
    nextUpdateDependencyID.value = projectUpdateDependencies.value.length + 1;
  } catch (err: any) {
    projectUpdateDependencies.value = [];
    nextUpdateDependencyID.value = 1;
    showError(err, "Failed to load project dependencies");
  }
}

function resetProjectCreateForm() {
  projectCreateForm.value = {
    namespace: "",
    projectName: "",
  };
  projectCreateDependencies.value = [];
  nextCreateDependencyID.value = 1;
}

function resetProjectUpdateForm() {
  projectUpdateForm.value = {
    namespace: "",
    projectName: "",
  };
  projectUpdateDependencies.value = [];
  nextUpdateDependencyID.value = 1;
}

function saveNewProject() {
  projectCommitMode.value = "create";
  projectCommitForm.summary = "";
  projectCommitForm.description = "";
  projectCommitModalOpen.value = true;
}

function saveProjectConfig() {
  projectCommitMode.value = "update";
  projectCommitForm.summary = "";
  projectCommitForm.description = "";
  projectCommitModalOpen.value = true;
}

function deleteProjectConfig() {
  projectCommitMode.value = "delete";
  projectCommitForm.summary = `chore(project): delete ${projectUpdateForm.value.projectName}`;
  projectCommitForm.description = "delete project directory and manifests";
  projectCommitModalOpen.value = true;
}

function cancelProjectCommit() {
  projectCommitModalOpen.value = false;
}

async function confirmProjectCommit() {
  if (!projectCommitForm.summary.trim()) {
    message.warning("请输入提交 summary");
    return;
  }

  const refreshBranchConfigAfterRequest =
    projectCommitMode.value === "version" ||
    projectCommitMode.value === "version-delete";

  try {
    projectCommitSaving.value = true;
    projectCommitModalOpen.value = false;
    if (projectCommitMode.value === "create") {
      const payload: ProjectRequest = {
        namespace: projectCreateForm.value.namespace,
        namespaceName: projectCreateForm.value.namespace,
        projectName: projectCreateForm.value.projectName,
        dependencies: projectCreateDependencies.value.map((d) => ({
          repoUrl: d.repoUrl,
          path: d.path,
          name: d.name,
          branchName: d.branchName,
          groups: d.groups,
        })),
        commitSummary: projectCommitForm.summary.trim(),
        commitDescription: projectCommitForm.description.trim(),
      };
      await createProject(payload);
      showSuccess("新增项目成功");
      resetProjectCreateForm();
    } else if (projectCommitMode.value === "update") {
      const payload: ProjectRequest = {
        namespace: projectUpdateForm.value.namespace,
        namespaceId: resolveNamespaceID(projectUpdateForm.value.namespace),
        projectId: resolveProject(projectUpdateForm.value.projectName, projectUpdateForm.value.namespace).id,
        projectName: projectUpdateForm.value.projectName,
        dependencies: projectUpdateDependencies.value.map((d) => ({
          repoUrl: d.repoUrl,
          path: d.path,
          name: d.name,
          branchName: d.branchName,
          groups: d.groups,
        })),
        commitSummary: projectCommitForm.summary.trim(),
        commitDescription: projectCommitForm.description.trim(),
      };
      await updateProject(projectUpdateForm.value.projectName, payload);
      showSuccess("更新项目成功");
    } else if (projectCommitMode.value === "version") {
      const payload: ProjectRequest = {
        namespace: branchConfigForm.value.namespace,
        namespaceId: resolveNamespaceID(branchConfigForm.value.namespace),
        projectId: resolveProject(branchConfigForm.value.projectName, branchConfigForm.value.namespace).id,
        projectName: branchConfigForm.value.projectName,
        dependencies: branchConfigDependencies.value.map((d) => ({
          repoUrl: d.repoUrl,
          path: d.path,
          name: d.name,
          branchName: d.branchName,
          groups: d.groups,
        })),
        commitSummary: projectCommitForm.summary.trim(),
        commitDescription: projectCommitForm.description.trim(),
      };
      await updateProjectVersion(
        branchConfigForm.value.namespace,
        branchConfigForm.value.projectName,
        branchConfigForm.value.version,
        payload,
      );
      showSuccess("版本 manifest 更新成功");
    } else if (projectCommitMode.value === "version-delete") {
      await deleteProjectVersion(
        branchConfigForm.value.namespace,
        branchConfigForm.value.projectName,
        branchConfigForm.value.version,
        {
          commitSummary: projectCommitForm.summary.trim(),
          commitDescription: projectCommitForm.description.trim(),
        },
      );
      showSuccess("发布版本和远程分支已删除");
    } else {
      const projectName = projectUpdateForm.value.projectName;
	  const namespace = projectUpdateForm.value.namespace;
      await deleteProject(projectUpdateForm.value.namespace, projectName, {
        commitSummary: projectCommitForm.summary.trim(),
        commitDescription: projectCommitForm.description.trim(),
      });
      showSuccess(`项目已删除: ${projectName}`);
    projectUpdateForm.value.projectName = "";
    projectUpdateDependencies.value = [];
    nextUpdateDependencyID.value = 1;
    projectUpdateTemplates.value = toProjectTemplates(await listProjects(namespace));
    if (branchConfigForm.value.namespace === namespace) {
    if (branchConfigForm.value.projectName === projectName) {
      branchConfigForm.value.projectName = "";
      branchConfigForm.value.version = "";
      branchConfigForm.value.baselineRef = "";
      branchConfigDependencies.value = [];
      branchVersions.value = [];
    }
    branchConfigTemplates.value = toProjectTemplates(await listProjects(namespace));
    }
    await loadNamespaces();
    }
    if (projectUpdateForm.value.namespace) {
      await handleProjectNamespaceChange(projectUpdateForm.value.namespace);
    }
    if (branchConfigForm.value.namespace) {
      await handleBranchNamespaceChange(branchConfigForm.value.namespace);
    }
  } catch (err: any) {
    showError(err, "提交失败");
  } finally {
    projectCommitSaving.value = false;
    if (refreshBranchConfigAfterRequest) {
      await refreshBranchConfig();
    }
  }
}

function getBackendError(err: any, fallback: string) {
  const data = err?.response?.data;
  if (typeof data === "string" && data.trim()) {
    return data;
  }
  return data?.msg ?? data?.error ?? data?.message ?? err?.message ?? fallback;
}

function showError(err: any, fallback: string) {
  notification.error({
    message: "操作失败",
    description: getBackendError(err, fallback),
    duration: 10,
    placement: "top",
    top: 72,
  });
}

function showSuccess(description: string) {
  notification.success({
    message: "操作成功",
    description,
    duration: 3,
    placement: "top",
    top: 72,
  });
}

async function loadCurrentUser() {
  try {
    const auth = await getCurrentUser();
    if (!auth.authenticated) {
      throw new Error("user is not authenticated");
    }
    savedUser.value = auth.user;
  } catch (err: any) {
    if (err?.response?.status !== 401) {
      showError(err, "Failed to load current user");
    }
    savedUser.value = null;
    localStorage.removeItem("release_hub_user_id");
  }
}

function onCreated(newTask: ReleaseTask) {
  task.value = newTask;
}

async function saveUserSettings() {
  try {
    if (!userForm.value.accessToken.trim() && savedUser.value) {
      settingsOpen.value = false;
      return;
    }
    settingsSaving.value = true;
    const auth = await upsertUser(userForm.value);
    if (!auth.authenticated) {
      throw new Error("GitHub authentication failed");
    }
    savedUser.value = auth.user;
    userForm.value.accessToken = "";
    settingsOpen.value = false;
    showSuccess(`用户配置已保存: ${auth.user.githubLogin}`);
  } catch (err: any) {
    showError(err, "保存用户配置失败");
  } finally {
    settingsSaving.value = false;
  }
}

async function loadNamespaces() {
  try {
    rootNamespaces.value = await listNamespaces();
  } catch (err: any) {
    rootNamespaces.value = [];
    showError(err, "Failed to load namespaces");
  }
}

async function onNamespacesChanged(change: { oldName?: string; newName?: string }) {
  if (change.oldName && projectCreateForm.value.namespace === change.oldName) {
    projectCreateForm.value.namespace = change.newName ?? "";
  }
  // Discard selections backed by the old names; all relationships still use stable IDs.
  resetProjectUpdateForm();
  resetBranchConfig();
  projectUpdateTemplates.value = [];
  branchConfigTemplates.value = [];
  namespaceRevision.value += 1;
  try {
    rootNamespaces.value = await listNamespaces();
    await refreshProjects();
  } catch (err: unknown) {
    showError(err, "刷新 Namespace 配置失败");
  }
}

void loadNamespaces();
void refreshProjects().catch((err: any) => showError(err, "Failed to load projects"));
loadCurrentUser();
</script>

<style scoped>
.shell {
  min-height: 100vh;
  background: #f5f7fb;
}

.operation-loading {
  position: fixed;
  inset: 50% auto auto 50%;
  transform: translate(-50%, -50%);
  z-index: 3000;
}

.brand {
  height: 56px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  color: #fff;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: 0.4px;
}

.main-header-fixed {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 56px;
  background: #fff;
  z-index: 2000;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.sider-toggle-btn {
  margin-right: 8px;
  font-size: 16px;
}

.main-header-title {
  font-size: 17px;
  font-weight: 600;
  color: #333;
  letter-spacing: 0.5px;
}

.header-spacer {
  flex: 1;
}

.account-button {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  color: #1f2937;
  font-weight: 600;
}

.account-avatar {
  background: #24292f;
  color: #fff;
}

.account-login {
  max-width: 180px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.login-button {
  font-weight: 600;
}

.account-menu-user {
  font-weight: 600;
}

.account-menu-login {
  color: #6b7280;
  font-size: 12px;
}

.account-menu-email {
  max-width: 240px;
  overflow: hidden;
  color: #6b7280;
  font-size: 12px;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.content {
  margin: 24px;
  min-height: 360px;
}

.intro-card,
.result-card,
.config-card {
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.06);
}

.intro-card {
  margin-bottom: 20px;
}

.intro-card :deep(.ant-card-body) {
  padding: 24px 28px;
}

.intro-card h1 {
  margin: 0 0 8px;
  font-size: 30px;
  color: #111827;
}

.intro-card p {
  margin: 0;
  color: #475569;
  line-height: 1.7;
}

.result-card {
  margin-top: 24px;
}

.config-card :deep(.ant-card-body) {
  padding: 24px;
}

.dependencies-toolbar {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 10px;
}

.dependencies-table :deep(.ant-table-cell) {
  vertical-align: middle;
}

.result-table {
  margin-top: 16px;
}

.footer {
  text-align: center;
  color: #94a3b8;
  font-size: 12px;
  background: transparent;
}

.settings-alert {
  margin-bottom: 16px;
}

.saved-user-card {
  margin-top: 20px;
}

@media (max-width: 992px) {
  .content {
    margin: 16px;
  }

  .intro-card h1 {
    font-size: 24px;
  }
}
</style>
