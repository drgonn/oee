import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryValvetimeList(params?: TableListParams) {
  return request(`/api/valvetime/list?token=${getToken()}`, {
    params,
  });
}

export async function addValvetime(params: TableListParams) {
  return request(`/api/valvetime?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeValvetime(params: { key: number[] }) {
  return request(`/api/valvetime?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateValvetime(params: TablePutItem) {
  return request(`/api/valvetime/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
