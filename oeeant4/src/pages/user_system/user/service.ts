import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data';

export async function queryUserList(params?: TableListParams) {
  return request(`/user/list?token=${getToken()}`, {
    params,
  });
}

export async function addUser(params: TableListParams) {
  return request(`/user/user?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeUser(params: { key: number[] }) {
  return request(`/user/user?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateUser(params: TablePutItem) {
  return request(`/user/user/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
