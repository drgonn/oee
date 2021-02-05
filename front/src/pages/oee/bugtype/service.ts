import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryBugtypeList(params?: TableListParams) {
  return request(`/api/bugtype/list?token=${getToken()}`, {
    params,
  });
}

export async function addBugtype(params: TableListParams) {
  return request(`/api/bugtype?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeBugtype(params: { key: number[] }) {
  return request(`/api/bugtype?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateBugtype(params: TablePutItem) {
  return request(`/api/bugtype/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
