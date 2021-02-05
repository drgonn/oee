import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryBugList(params?: TableListParams) {
  return request(`/api/bug/list?token=${getToken()}`, {
    params,
  });
}

export async function addBug(params: TableListParams) {
  return request(`/api/bug?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeBug(params: { key: number[] }) {
  return request(`/api/bug?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateBug(params: TablePutItem) {
  return request(`/api/bug/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
