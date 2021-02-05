import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryValveList(params?: TableListParams) {
  return request(`/api/valve/list?token=${getToken()}`, {
    params,
  });
}

export async function addValve(params: TableListParams) {
  return request(`/api/valve?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeValve(params: { key: number[] }) {
  return request(`/api/valve?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateValve(params: TablePutItem) {
  return request(`/api/valve/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
