<template>
  <div class="odyssey-page">
    <div class="odyssey-inner-page">
      <a-table
        row-key="user_group"
        size="small"
        :bordered="{
          cell: true
        }"
        :columns="columns"
        :data="data"
        :row-selection="rowSelection"
        v-model:selectedKeys="selectedKeys"
        :pagination="pagination"
        @page-change="pageChange"
        @page-size-change="pageSizeChange"
      >
        <template #avatar="{ record }">
          <a-avatar
            auto-fix-font-size
            :size="30"
            :style="{
              backgroundColor: '#14a9f8'
            }"
          >
            {{ record.avatar }}
          </a-avatar>
        </template>
        <template #status="{ record }">
          <a-space>
            <a-tag size="small" color="green" v-if="record.status == 1">启用</a-tag>
            <a-tag size="small" color="red" v-else>停用</a-tag>
          </a-space>
        </template>
        <template #optional="{ record }">
          <a-space>
            <a-button size="mini" type="primary">详情</a-button>
            <a-button size="mini" >修改</a-button>
            <a-popconfirm content="确定删除这条数据吗?" type="warning">
              <a-button size="mini" type="primary" status="danger" >删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue';
import axios from 'axios';
const API_BASE_URL = import.meta.env.VITE_APP_BASE_URL;
interface UserGroup {
  [key: string]: any; // 如果有额外属性，可以添加可选字段
}
const data = reactive<UserGroup[]>([]);
const selectedKeys = ref([]);
const rowSelection = reactive({
  type: "checkbox",
  showCheckedAll: true,
  onlyCurrent: false
});
const pagination = ref({ showPageSize: true, showTotal: true, current: 1, pageSize: 10, total: 10 });
const pageChange = (page: number) => {
  pagination.value.current = page;
};
const pageSizeChange = (pageSize: number) => {
  pagination.value.pageSize = pageSize;
};
const columns = [
  { title: "ID", dataIndex: "id" },
  { title: "组名称", dataIndex: "user_group" },
  { title: "描述", dataIndex: "description" },
  { title: "操作", slotName: "optional", align: "center" }
];

const fetchData = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/api/user/usergroup/list`);
    //console.log(response.data);
    data.splice(0, data.length, ...response.data.result.user_groups);
    pagination.value.total = response.data.result.total;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

onMounted(() => {
  fetchData();
});
</script>

<style lang="scss" scoped></style>
