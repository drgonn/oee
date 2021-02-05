import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryPlanList(params?: TableListParams) {
  return request(`/api/plan/list?token=${getToken()}`, {
    params,
  });
}

export async function addPlan(params: TableListParams) {
  return request(`/api/plan?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removePlan(params: { key: number[] }) {
  return request(`/api/plan?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updatePlan(params: TablePutItem) {
  return request(`/api/plan/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
