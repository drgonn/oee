import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryWorktimeList(params?: TableListParams) {
  return request(`/api/worktime/list?token=${getToken()}`, {
    params,
  });
}

export async function addWorktime(params: TableListParams) {
  return request(`/api/worktime?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeWorktime(params: { key: number[] }) {
  return request(`/api/worktime?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateWorktime(params: TablePutItem) {
  return request(`/api/worktime/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
