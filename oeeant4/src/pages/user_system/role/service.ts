import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryRoleList(params?: TableListParams) {
  return request(`/user/role/list?token=${getToken()}`, {
    params,
  });
}

export async function addRole(params: TableListParams) {
  return request(`/user/role?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeRole(params: { key: number[] }) {
  return request(`/user/role?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateRole(params: TablePutItem) {
  return request(`/user/role/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
