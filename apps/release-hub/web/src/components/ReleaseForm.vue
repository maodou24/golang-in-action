<!-- eslint-disable vue/no-v-model-argument -->
<template>
  <a-card title="Create Release Branch" :bordered="false" class="form-card">
    <a-spin :spinning="loading" tip="Processing...">
    <a-form layout="vertical" :model="formState" @finish="onSubmit">
      <a-row :gutter="16">
        <a-col :xs="24" :md="12">
          <a-form-item
            label="Release Branch Name"
            name="targetBranchName"
            :rules="[nonEmptyRule('请输入分支名称')]"
          >
            <a-input v-model:value="formState.targetBranchName" placeholder="3.0.0" />
          </a-form-item>
        </a-col>
      </a-row>

      <a-row :gutter="16">
        <a-col :xs="24" :md="12">
          <a-form-item label="Namespace" name="namespace">
            <a-select v-model:value="formState.namespace" :options="namespaceOptions" @change="handleNamespaceChange" />
          </a-form-item>
        </a-col>
      </a-row>

      <a-row :gutter="16">
        <a-col :xs="24" :md="12">
          <a-form-item
            label="Project"
            name="projectName"
            :rules="[nonEmptyRule('请选择项目')]"
          >
            <a-select
              v-model:value="formState.projectName"
              show-search
              placeholder="支持搜索配置项目"
              :options="projectOptions"
              option-filter-prop="label"
              @change="handleProjectChange"
            />
          </a-form-item>
        </a-col>
      </a-row>

      <a-row :gutter="16">
        <a-col :xs="24" :md="12">
          <a-form-item
            label="Baseline"
            name="baselineRef"
            :rules="[nonEmptyRule('请选择 baseline')]"
          >
            <a-select
              v-model:value="formState.baselineRef"
              placeholder="支持引用当前其他分支，可选 default"
              :options="baselineOptions"
              @change="handleBaselineChange"
            />
          </a-form-item>
        </a-col>
      </a-row>

      <a-form-item label="Dependencies">
        <a-table
          class="dependencies-table"
          size="small"
          :pagination="false"
          :data-source="dependencyRows"
          :columns="dependencyColumns"
          row-key="id"
        >
          <template #footer>
            <a-button type="dashed" @click="addDependency">新增依赖</a-button>
          </template>
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'repoUrl'">
              <a-input v-model:value="record.repoUrl" />
            </template>
            <template v-if="column.key === 'path'">
              <a-input v-model:value="record.path" />
            </template>
            <template v-if="column.key === 'name'">
              <a-input v-model:value="record.name" />
            </template>
            <template v-if="column.key === 'branchName'">
              <a-input v-model:value="record.branchName" placeholder="Branch Name" />
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
              <a-space>
                <a-button
                  v-if="formState.targetBranchName.trim()"
                  type="link"
                  size="small"
                  :disabled="record.branchName.trim() === formState.targetBranchName.trim()"
                  @click="replaceBranch(record)"
                >
                  Replace
                </a-button>
                <a-button danger type="link" @click="removeDependency(record.id)">删除</a-button>
              </a-space>
            </template>
          </template>
        </a-table>
      </a-form-item>

      <a-space>
        <a-button type="primary" html-type="submit" :loading="loading"
          >Create</a-button
        >
      </a-space>
    </a-form>
    </a-spin>

    <a-modal
      :open="commitModalOpen"
      title="提交信息"
      ok-text="提交"
      cancel-text="取消"
      :confirm-loading="loading"
      @update:open="handleCommitModalOpenChange"
      @ok="confirmSubmit"
      @cancel="cancelSubmit"
    >
      <a-form layout="vertical">
        <a-form-item label="Summary" required>
          <a-input
            :value="commitForm.summary"
            @update:value="(value: string) => (commitForm.summary = value)"
          />
        </a-form-item>
        <a-form-item label="Description">
          <a-textarea
            :value="commitForm.description"
            :rows="4"
            placeholder="补充提交描述"
            @update:value="(value: string) => (commitForm.description = value)"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </a-card>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref } from "vue";
import { message, notification } from "ant-design-vue";
import {
  createRelease,
  getProjectVersion,
  groupsAsArray,
  listNamespaces,
  listProjectVersions,
  listProjects,
  resolveProject,
} from "../api/client";
import type {
  CreateReleaseRequest,
  ProjectDependency,
  ProjectConfig,
  ReleaseTask,
} from "../types/release";

interface DependencyRow {
  id: number;
  repoUrl: string;
  path: string;
  name: string;
  branchName: string;
  groups: string[];
}

export default defineComponent({
  name: "ReleaseForm",
  emits: ["created"],
  setup(_, { emit }) {
    const loading = ref(false);

    const defaultState: CreateReleaseRequest = {
      projectName: "",
      namespace: "",
      releaseName: "",
      baselineRef: "",
      targetBranchName: "",
      commitSummary: "",
      commitDescription: "",
    };

    const formState = reactive<CreateReleaseRequest>({ ...defaultState });
    const commitModalOpen = ref(false);
    const commitForm = reactive({
      summary: "",
      description: "",
    });
    const baselineOptions = ref<Array<{ label: string; value: string }>>([]);
    const dependencyRows = ref<DependencyRow[]>([]);
    const nextDependencyID = ref(1);

    function nonEmptyRule(messageText: string) {
      return {
        validator: (_rule: unknown, value: unknown) => {
          if (typeof value === "string" && value.trim()) {
            return Promise.resolve();
          }
          return Promise.reject(new Error(messageText));
        },
      };
    }

    const projects = ref<ProjectConfig[]>([]);
    const namespaces = ref<string[]>([]);
    const namespaceOptions = computed(() => namespaces.value.map((value) => ({ label: value, value })));
    const projectOptions = computed(() =>
      projects.value.map((p) => ({
        label: `${p.namespace}/${p.projectName}`,
        value: p.projectName,
      })),
    );

    const dependencyColumns = [
      { title: "Repository", dataIndex: "repoUrl", key: "repoUrl", width: "28%" },
      { title: "Path", dataIndex: "path", key: "path", width: "14%" },
      { title: "Name", dataIndex: "name", key: "name", width: "18%" },
      {
        title: "Branch Name",
        dataIndex: "branchName",
        key: "branchName",
        width: "14%",
      },
      { title: "Groups", dataIndex: "groups", key: "groups", width: "26%" },
      { title: "操作", dataIndex: "action", key: "action", width: 90 },
    ];

    const groupOptions = computed(() => {
      const groups = dependencyRows.value.reduce<string[]>(
        (allGroups, dependency) => allGroups.concat(dependency.groups),
        [],
      );
      return [...new Set(groups)].sort().map((value) => ({ label: value, value }));
    });

    function setDependencies(rows: ProjectDependency[]) {
      dependencyRows.value = rows.map((d, index) => ({
        id: index + 1,
        repoUrl: d.repoUrl,
        path: d.path,
        name: d.name,
        branchName: d.branchName,
        groups: groupsAsArray(d.groups),
      }));
      nextDependencyID.value = dependencyRows.value.length + 1;
    }

    async function loadBaselineDependencies(version: string) {
      if (!formState.namespace || !formState.projectName || !version) {
        dependencyRows.value = [];
        nextDependencyID.value = 1;
        return;
      }
      try {
        const project = await getProjectVersion(
          formState.namespace,
          formState.projectName,
          version,
        );
        setDependencies(project.dependencies);
      } catch (err: any) {
        dependencyRows.value = [];
        nextDependencyID.value = 1;
        showError(err, "Failed to load baseline manifest");
      }
    }

    function handleBaselineChange(value: string) {
      formState.baselineRef = value;
      void loadBaselineDependencies(value);
    }

    function replaceBranch(row: DependencyRow) {
      const branchName = formState.targetBranchName.trim();
      if (branchName) {
        row.branchName = branchName;
      }
    }

    async function handleProjectChange(value: string) {
      const selected = projects.value.find(
        (p) => p.namespace === formState.namespace && p.projectName === value,
      );
      if (!selected) {
        baselineOptions.value = [{ label: "default", value: "default" }];
        dependencyRows.value = [];
        nextDependencyID.value = 1;
        formState.baselineRef = "default";
        return;
      }

      try {
        const versions = await listProjectVersions(selected.namespace, selected.projectName);
        const orderedVersions = ["default", ...versions.filter((version) => version !== "default")];
        baselineOptions.value = orderedVersions.map((version) => ({
          label: `${selected.namespace}/${selected.projectName}/${version}`,
          value: version,
        }));
        formState.baselineRef = orderedVersions.includes("default") ? "default" : "";
      } catch (err: any) {
        baselineOptions.value = [];
        formState.baselineRef = "";
        showError(err, "Failed to load release versions");
      }
      await loadBaselineDependencies(formState.baselineRef);
    }

    async function handleNamespaceChange(value: string) {
      formState.namespace = value;
      formState.projectName = "";
      formState.baselineRef = "";
      baselineOptions.value = [];
      dependencyRows.value = [];
      nextDependencyID.value = 1;
      try {
        projects.value = await listProjects(value);
      } catch (err: any) {
        projects.value = [];
        showError(err, "Failed to load projects");
      }
    }

    function reset() {
      Object.assign(formState, defaultState);
      baselineOptions.value = [];
      dependencyRows.value = [];
      nextDependencyID.value = 1;
      commitForm.summary = "";
      commitForm.description = "";
    }

    function addDependency() {
      dependencyRows.value.push({
        id: nextDependencyID.value,
        repoUrl: "",
        path: "",
        name: "",
        branchName: "",
        groups: [],
      });
      nextDependencyID.value += 1;
    }

    function removeDependency(id: number) {
      dependencyRows.value = dependencyRows.value.filter(
        (row) => row.id !== id,
      );
    }

    function onSubmit() {
      if (!formState.targetBranchName.trim()) {
        message.warning("请输入 Release Branch Name");
        return;
      }
      if (!formState.projectName.trim()) {
        message.warning("请选择 Project");
        return;
      }
      if (!formState.baselineRef.trim()) {
        message.warning("请选择 Baseline");
        return;
      }
      commitForm.summary = "";
      commitForm.description = "";
      commitModalOpen.value = true;
    }

    function cancelSubmit() {
      commitModalOpen.value = false;
    }

    function handleCommitModalOpenChange(value: boolean) {
      commitModalOpen.value = value;
    }

    async function confirmSubmit() {
      if (!commitForm.summary.trim()) {
        message.warning("请输入提交 summary");
        return;
      }

      try {
        loading.value = true;
        commitModalOpen.value = false;
        const payload: CreateReleaseRequest = {
          ...formState,
          projectId: resolveProject(formState.projectName, formState.namespace).id,
          dependencies: dependencyRows.value.map(({ id, ...dependency }) => dependency),
          releaseName: formState.targetBranchName,
          projectName: formState.projectName,
          commitSummary: commitForm.summary.trim(),
          commitDescription: commitForm.description.trim(),
        };
        const result: ReleaseTask = await createRelease(payload);
        showSuccess(`创建成功: ${result.taskId}`);
        emit("created", result);
      } catch (err: any) {
        showError(err, "创建失败");
      } finally {
        loading.value = false;
      }
    }

    function getBackendError(err: any, fallback: string) {
      const data = err?.response?.data;
      if (typeof data === "string" && data.trim()) {
        return data;
      }
      return data?.error ?? data?.message ?? err?.message ?? fallback;
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
        duration: 2,
        placement: "top",
        top: 72,
      });
    }

    listNamespaces()
      .then((items) => {
        namespaces.value = items;
      })
      .catch((err: any) => {
        showError(err, "Failed to load namespaces");
      });

    return {
      loading,
      formState,
      projectOptions,
      baselineOptions,
      dependencyRows,
      dependencyColumns,
      groupOptions,
      nonEmptyRule,
      addDependency,
      removeDependency,
      handleProjectChange,
      handleBaselineChange,
      handleNamespaceChange,
      replaceBranch,
      namespaceOptions,
      onSubmit,
      cancelSubmit,
      handleCommitModalOpenChange,
      confirmSubmit,
      commitModalOpen,
      commitForm,
      reset,
    };
  },
});
</script>

<style scoped>
.form-card {
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.06);
}

.form-card :deep(.ant-card-head) {
  border-bottom: 1px solid #eef2f7;
}

.form-card :deep(.ant-card-head-title) {
  font-weight: 600;
}

.form-card :deep(.ant-card-body) {
  padding: 24px;
}

.dependencies-table :deep(.ant-table-cell) {
  vertical-align: middle;
}

.dependencies-toolbar {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 10px;
}
</style>
