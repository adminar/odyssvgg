<template>
  <div class="odyssey-page">
    <div class="odyssey-inner-page">
      <a-form :model="formData.form" :label-col-props="{ span: 6 }" :wrapper-col-props="{ span: 18 }">
        <a-row :gutter="16">
          <a-col :span="6">
            <a-form-item field="name" label="姓名">
              <a-input v-model="formData.form.username" placeholder="请输入姓名" allow-clear />
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item field="phone" label="手机号">
              <a-input v-model="formData.form.phone" placeholder="请输入手机号" allow-clear />
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item field="email" label="邮箱">
              <a-input v-model="formData.form.email" placeholder="请输入邮箱" allow-clear />
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-space>
              <a-button type="primary">
                <template #icon>
                  <icon-search />
                </template>
                <template #default>查询</template>
              </a-button>
              <a-button>
                <template #icon>
                  <icon-refresh />
                </template>
                <template #default>重置</template>
              </a-button>
              <a-button type="text" @click="formData.search = !formData.search">
                <template #icon>
                  <icon-up v-if="formData.search" />
                  <icon-down v-else />
                </template>
                <template #default>{{ formData.search ? "收起" : "展开" }}</template>
              </a-button>
            </a-space>
          </a-col>
        </a-row>
        <a-row :gutter="16" v-if="formData.search">
          <a-col :span="6">
            <a-form-item field="roles" label="角色">
              <a-input v-model="formData.form.roles" placeholder="请输入角色" allow-clear />
            </a-form-item>
          </a-col>
          <a-col :span="6">
            <a-form-item field="status" label="用户状态">
              <a-select v-model="formData.form.status" placeholder="请选择用户状态" allow-clear>
                <a-option value="1">停用</a-option>
                <a-option value="2">启用</a-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
      <a-table
        row-key="username"
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
            <a-button size="mini" @click="openReviewDrawer(record)">修改</a-button>
            <a-popconfirm content="确定删除这条数据吗?" type="warning" @ok="() => handleDelete(record.username)">
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
import { Message } from "@arco-design/web-vue";
import axios from 'axios';
  // @ts-ignore
const API_BASE_URL = import.meta.env.VITE_APP_BASE_URL;
interface User {
  username: string;
  phone: string;
  email: string;
  roles: string;
  status: number;
  createTime: string;
}

// 定义类型
interface RolesPro {
  project: string[];
  env: string[];
  roles: string[];
}

const formData = reactive({
  form: {
    username: "",
    phone: "",
    email: "",
    roles: "",
    createTime :"",
    status: null,
  },
  search: false
});

const formRolesPro = reactive<RolesPro>({
  project: [],
  env: [],
  roles: []
});

const currentRecord = ref<User | null>(null); // 当前选中的记录
const drawerEditVisible = ref(false); // 控制抽屉的可见性
const selectedRoles = ref<string[]>([]); 
const selectedProject = ref<string[]>([]); 
const selectedEnv = ref<string[]>([]); 
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
  { title: "姓名", dataIndex: "username" },
  { title: "用户类型", dataIndex: "type" },
  { title: "手机号", dataIndex: "phone" },
  { title: "Email", dataIndex: "email" },
  { title: "用户组", dataIndex: "user_group" },
  { title: "最后登录时间", dataIndex: "last_login" },
  { title: "操作", slotName: "optional", align: "center" }
];

const data = reactive<User[]>([]);

const fetchData = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/api/user/list`);
    //console.log(response.data);
    data.splice(0, data.length, ...response.data.result.users);
    pagination.value.total = response.data.result.total;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

const handleDelete = async (username: string) => {
  try {
    // 请求后端以POST方法删除数据，携带ID参数
    const response = await axios.delete(`${API_BASE_URL}/api/user/delete/${username}`);
    // 检查删除是否成功
    if (response.data.status === 200) {
      Message.success("删除成功");
      // 重新获取数据以刷新表格
      fetchData();
    } else {
      Message.error( "删除失败");
    }
  } catch (error) {
    Message.error("删除失败，请重试");
  }
};

const changeUserInfo = async () => {
  if (currentRecord.value?.roles !== selectedRoles.value.join(",")){
    const rolesData: RolesPro = {
      project: selectedProject.value,
      env: selectedEnv.value,
      roles: selectedRoles.value,
    };
    const userData = {
      username: currentRecord.value?.username,
      roles: JSON.stringify(rolesData),
    };
    console.log(rolesData)
    try {
      // 请求后端以POST方法删除数据，携带ID参数
      const response = await axios.post(`${API_BASE_URL}/api/user/change`,userData);
      console.log('提交响应:', response.data);
      Message.success("修改成功");
      fetchData()
    } catch (error) {
      Message.error("修改失败");
    }
  }
}

const openReviewDrawer = (record: User) => {
  let parsedRolesData: RolesPro | null = null;
  parsedRolesData = JSON.parse(record.roles) as RolesPro;
  console.log(parsedRolesData)
  selectedRoles.value = parsedRolesData?.roles || []; // 选中当前记录的角色
  selectedProject.value = parsedRolesData?.project || []; // 选中当前记录的项目
  selectedEnv.value = parsedRolesData?.env || []; // 选中当前记录的环境
  currentRecord.value = record;
  drawerEditVisible.value = true; // 打开抽屉
};

const closeDrawer = () => {
  drawerEditVisible.value = false; // 关闭抽屉
};

onMounted(() => {
  fetchData();
});
</script>

<style lang="scss" scoped></style>
