<template>
  <div class="flex flex-col overflow-auto">
    <!-- 按钮条 -->
    <div class="table-bar flex justify-between items-center mb-3">
      <div>
        <ElButton icon="Plus" @click="form.toAdd">新增</ElButton>
        <ElButton icon="Delete"@click="form.toDelete">删除</ElButton>
      </div>
      <div>
        <ElButton icon="Refresh" round @click="table.request"></ElButton>
        <ElButton icon="Search" round @click="search.show = !search.show"></ElButton>
      </div>
    </div>
    <!-- 列表 -->
    <AgelTable class="flex-1" v-bind="table" v-model:page="table.page"> </AgelTable>
    <!-- 弹窗表单 -->
    <ElDialog v-model="form.show" :title="form.title" width="800px" top="10vh">
      <ElForm :ref="(v) => (form.ref = v)" :model="form.model" label-width="80px">
        <AgelFormDesc 
          :items="form.items.filter(item => item.prop !== 'config' || form.state !== 'edit')" 
          :view-model="form.state === 'view'"
        ></AgelFormDesc>
      </ElForm>
      <template #footer>
        <ElButton v-if="form.state === 'add'" type="primary" @click="form.submit">提交</ElButton>
        <ElButton v-if="form.state === 'edit'" type="primary" @click="form.submitEdit">提交</ElButton>
      </template>
    </ElDialog>
  </div>
</template>

<script lang="jsx" setup>
import { reactive, ref, nextTick } from 'vue'
import http from '@/api'
import yaml from 'js-yaml'

const form = reactive({
  test: false,
  ref: null,
  show: false,
  title: '',
  state: 'edit',
  model: { name: '', api_server: '', region: '', version: '', description: '', config: '' },
  items: [
    { label: '集群名称', prop: 'name', span: 3, required: true },
    { label: 'api_server', span: 3, prop: 'api_server', required: true },
    { label: 'Region', span: 3, prop: 'region' },
    { label: '集群版本', span: 3, prop: 'version' },
    {
      label: '介绍',
      prop: 'description',
      span: 3,
      attrs: { type: 'textarea', rows: 5 }
    },
    {
      label: '集群凭证',
      prop: 'config',
      span: 3,
      attrs: { type: 'textarea', rows: 20 },
    }
  ],
  toAdd: () => {
    form.show = true
    form.title = '新增集群'
    form.state = 'add'
    nextTick(() => {
      form.ref?.resetFields()
    })
  },
  toEdit: (row) => {
    form.show = true
    form.title = '编辑集群'
    form.state = 'edit'
    nextTick(() => {
      form.model = { ...row }
    })
  },
  toView: (row) => {
    form.show = true
    form.title = '查看集群'
    form.state = 'view'
    nextTick(() => {
      form.model = { ...row }
    })
  },
  submit: () => {
    form.ref?.validate().then(() => {
      form.show = false
      table.create(form.model)
    })
  },
  submitEdit: () => {
    form.ref?.validate().then(() => {
      form.show = false
      table.edit(form.model)
    })
  }
})

const table = reactive({
  loading: false,
  border: true,
  data: [],
  columns: [
    { label: '#', type: 'selection' },
    { label: '集群名称', prop: 'name', width: 200 },
    { label: '集群api_server', prop: 'api_server', width: 200 },
    { label: 'Region', prop: 'region', width: 100 },
    { label: '集群版本', prop: 'version', width: 100 },
    { label: '介绍', prop: 'description', showOverflowTooltip: true },
    {
      width: '120px',
      label: '操作',
      fixed: 'right',
      slot: (scope) => {
        return (
          <div>
            <el-button link type="primary" onClick={() => form.toView(scope.row)}>
              查看
            </el-button>
            <el-divider direction="vertical" />
            <el-button link type="primary" onClick={() => form.toEdit(scope.row)}>
              编辑
            </el-button>
          </div>
        )
      }
    }
  ],
  request: () => {
    table.loading = true
    http.get(import.meta.env.VITE_APP_BASE_URL + `/cicd/fetch/repo/cluster`).then((res) => {
      table.data = res.result
      table.loading = false
    })
  },
  create: (form) => {
    form.config = yaml.dump(form.config)
      .replace(/^ *\|-\n/, '')    // 去掉 |-\n 后的两个空格
      .replace(/\\n {2}/g, '\n')    // 去掉 \n 后的两个空格
      .toString()
    table.loading = true
    http.post(import.meta.env.VITE_APP_BASE_URL + `/cicd/create/repo/cluster`, form).then((res) => {
      table.loading = false
      table.request()
    })
  },  
  edit: (form) => {
    table.loading = true
    http.post(import.meta.env.VITE_APP_BASE_URL + `/cicd/update/repo/cluster`, form).then((res) => {
      table.loading = false
      table.request()
    })
  }
})

table.request()
</script>
