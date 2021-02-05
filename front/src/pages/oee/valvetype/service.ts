import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryValvetypeList(params?: TableListParams) {
  return request(`/api/valvetype/list?token=${getToken()}`, {
    params,
  });
}

export async function addValvetype(params: TableListParams) {
  return request(`/api/valvetype?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeValvetype(params: { key: number[] }) {
  return request(`/api/valvetype?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateValvetype(params: TablePutItem) {
  return request(`/api/valvetype/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
