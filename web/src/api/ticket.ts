import axios from 'axios';
import { useUserInfoStore } from '@/store/modules/user-info';

// 获取环境变量中的API基础URL
const baseURL = import.meta.env.VITE_APP_BASE_URL;

/**
 * 从Cookie或user-info存储中获取当前用户ID
 * @returns 当前用户ID，如果未找到则返回0
 */
export const getCurrentUserId = async (): Promise<number> => {
  try {
    // 从用户存储中获取用户信息
    const userStore = useUserInfoStore();
    const username = userStore.account.username;
    
    if (!username) {
      console.warn('未获取到用户名，请先登录');
      return 0;
    }
    
    // 调用API获取用户ID
    const response = await axios.get(`${baseURL}/user/current`);
    
    if (response.data && response.data.status === 200) {
      const userId = response.data.result.id;
      return userId;
    } else {
      throw new Error(response.data?.message || '获取用户ID失败');
    }
  } catch (error) {
    console.error('获取用户ID失败:', error);
    return 0;
  }
};

/**
 * 工单管理API服务
 */
export default {
  /**
   * 获取当前用户ID
   */
  getCurrentUserId,

  /**
   * 获取我的工单列表
   * @param params 查询参数
   */
  getMyTickets(params) {
    return axios.get(`${baseURL}/ticket/list`, { params });
  },

  /**
   * 获取所有工单列表
   * @param params 查询参数
   */
  getAllTickets(params) {
    return axios.get(`${baseURL}/ticket/all`, { params });
  },

  /**
   * 获取待审批工单列表
   * @param params 查询参数
   */
  getPendingApprovalTickets(params) {
    return axios.get(`${baseURL}/ticket/approval/list`, { params });
  },

  /**
   * 获取工单详情
   * @param id 工单ID
   */
  getTicketDetail(id) {
    return axios.get(`${baseURL}/ticket/${id}/detail`);
  },

  /**
   * 创建工单
   * @param data 工单数据
   */
  createTicket(data) {
    return axios.post(`${baseURL}/ticket/create`, data);
  },

  /**
   * 更新工单
   * @param id 工单ID
   * @param data 工单数据
   */
  updateTicket(id, data) {
    return axios.post(`${baseURL}/ticket/${id}/update`, data);
  },

  /**
   * 提交工单
   * @param id 工单ID
   */
  submitTicket(id) {
    return axios.post(`${baseURL}/ticket/${id}/submit`);
  },

  /**
   * 取消工单
   * @param id 工单ID
   */
  cancelTicket(id) {
    return axios.post(`${baseURL}/ticket/${id}/cancel`);
  },

  /**
   * 审批通过工单
   * @param data 审批数据
   */
  approveTicket(data) {
    return axios.post(`${baseURL}/ticket/approval/approve`, data);
  },

  /**
   * 拒绝工单
   * @param data 拒绝数据
   */
  rejectTicket(data) {
    return axios.post(`${baseURL}/ticket/approval/reject`, data);
  },

  /**
   * 获取工单模板列表
   */
  getTemplates() {
    return axios.get(`${baseURL}/ticket/template/list`);
  },

  /**
   * 获取工单模板详情
   * @param id 模板ID
   */
  getTemplateDetail(id) {
    return axios.get(`${baseURL}/ticket/template/${id}`);
  },

  /**
   * 创建工单模板
   * @param data 模板数据
   */
  createTemplate(data) {
    return axios.post(`${baseURL}/ticket/template/create`, data);
  },

  /**
   * 更新工单模板
   * @param id 模板ID
   * @param data 模板数据
   */
  updateTemplate(id, data) {
    return axios.put(`${baseURL}/ticket/template/${id}`, data);
  },

  /**
   * 删除工单模板
   * @param id 模板ID
   */
  deleteTemplate(id) {
    return axios.delete(`${baseURL}/ticket/template/${id}`);
  },

  /**
   * 获取工作流列表
   */
  getWorkflows() {
    return axios.get(`${baseURL}/ticket/workflow/list`);
  },

  /**
   * 获取工作流详情
   * @param id 工作流ID
   */
  getWorkflowDetail(id) {
    return axios.get(`${baseURL}/ticket/workflow/${id}`);
  },

  /**
   * 创建工作流
   * @param data 工作流数据
   */
  createWorkflow(data) {
    return axios.post(`${baseURL}/ticket/workflow/create`, data);
  },

  /**
   * 更新工作流
   * @param id 工作流ID
   * @param data 工作流数据
   */
  updateWorkflow(id, data) {
    return axios.put(`${baseURL}/ticket/workflow/${id}`, data);
  },

  /**
   * 删除工作流
   * @param id 工作流ID
   */
  deleteWorkflow(id) {
    return axios.delete(`${baseURL}/ticket/workflow/${id}`);
  },

  /**
   * 获取下一步节点选项
   * @param ticketId 工单ID
   * @param currentNodeId 当前节点ID
   */
  getNextNodes(ticketId, currentNodeId) {
    return axios.get(`${baseURL}/ticket/workflow/${ticketId}/next-nodes`, {
      params: { currentNodeId }
    });
  },

  /**
   * 下载附件
   * @param id 附件ID
   */
  downloadAttachment(id) {
    return `${baseURL}/ticket/attachment/${id}/download`;
  }
}; 