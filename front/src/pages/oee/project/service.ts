import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryProjectList(params?: TableListParams) {
  return request(`/api/project/list?token=${getToken()}`, {
    params,
  });
}

export async function addProject(params: TableListParams) {
  return request(`/api/project?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeProject(params: { key: number[] }) {
  return request(`/api/project?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateProject(params: TablePutItem) {
  return request(`/api/project/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
