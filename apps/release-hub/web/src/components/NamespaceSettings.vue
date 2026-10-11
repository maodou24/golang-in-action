<script setup lang="ts">
import { computed, ref } from 'vue'
import { message } from 'ant-design-vue'
import axios from 'axios'
import { createNamespace, deleteNamespace, listNamespaceRecords, updateNamespace } from '../api/client'
import type { NamespaceRecord } from '../types/release'

const emit = defineEmits<{
  changed: [change: { oldName?: string; newName?: string }]
}>()
const open = ref(false)
const loading = ref(false)
const saving = ref(false)
const rows = ref<NamespaceRecord[]>([])
const search = ref('')
const name = ref('')
const editing = ref<NamespaceRecord | null>(null)
const filtered = computed(() => rows.value.filter((row) =>
  row.name.toLowerCase().includes(search.value.trim().toLowerCase())))
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: 'Namespace', dataIndex: 'name', key: 'name' },
  { title: '操作', key: 'actions', width: 150 },
]

function report(error: unknown) {
  const text = axios.isAxiosError<{ msg?: string }>(error)
    ? error.response?.data?.msg ?? error.message
    : error instanceof Error ? error.message : '操作失败'
  message.error(text)
}

async function load() {
  loading.value = true
  try { rows.value = await listNamespaceRecords() }
  catch (error) { report(error) }
  finally { loading.value = false }
}

function reset() { editing.value = null; name.value = '' }

async function show() {
  open.value = true
  reset()
  await load()
}

async function save() {
  const value = name.value.trim()
  if (!value) { message.warning('请输入 Namespace 名称'); return }
  saving.value = true
  try {
    const oldName = editing.value?.name
    const result = editing.value
      ? await updateNamespace(editing.value.id, value)
      : await createNamespace(value)
    emit('changed', { oldName, newName: result.name })
    message.success(editing.value ? 'Namespace 已修改' : 'Namespace 已创建')
    reset()
    await load()
  } catch (error) { report(error) }
  finally { saving.value = false }
}

async function remove(row: NamespaceRecord) {
  saving.value = true
  try {
    await deleteNamespace(row.id)
    if (editing.value?.id === row.id) reset()
    emit('changed', { oldName: row.name })
    message.success('Namespace 已删除')
    await load()
  } catch (error) { report(error) }
  finally { saving.value = false }
}
</script>

<template>
  <a-button type="text" @click="show">Namespace 配置</a-button>
  <a-drawer v-model:open="open" title="Namespace 基础配置" :width="640" :style="{ maxWidth: '100vw' }">
    <a-alert type="info" show-icon message="Namespace 用于项目分类。修改名称不会改变 ID 或项目绑定；仅无项目的 Namespace 可以删除。" />
    <a-form layout="vertical" class="namespace-form" @finish="save">
      <a-form-item :label="editing ? `修改 Namespace #${editing.id}` : '新增 Namespace'">
        <a-input v-model:value="name" placeholder="例如 platform、mobile" :disabled="saving" />
      </a-form-item>
      <a-space>
        <a-button type="primary" html-type="submit" :loading="saving">{{ editing ? '保存修改' : '新增' }}</a-button>
        <a-button v-if="editing" :disabled="saving" @click="reset">取消修改</a-button>
      </a-space>
    </a-form>
    <a-space class="namespace-toolbar">
      <a-input-search v-model:value="search" placeholder="搜索 Namespace" allow-clear />
      <a-button :loading="loading" :disabled="saving" @click="load">刷新列表</a-button>
    </a-space>
    <a-table :data-source="filtered" :columns="columns" row-key="id" :loading="loading" :pagination="{ pageSize: 10 }">
      <template #bodyCell="{ column, record }">
        <a-space v-if="column.key === 'actions'">
          <a-button type="link" :disabled="saving" @click="editing = record; name = record.name">修改</a-button>
          <a-popconfirm title="确认删除？仅空 Namespace 可以删除。" @confirm="remove(record)">
            <a-button type="link" danger :disabled="saving">删除</a-button>
          </a-popconfirm>
        </a-space>
      </template>
    </a-table>
  </a-drawer>
</template>

<style scoped>
.namespace-form { margin: 24px 0; }
.namespace-toolbar { display: flex; margin-bottom: 16px; }
</style>