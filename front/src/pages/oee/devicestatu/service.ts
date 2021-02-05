import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryDevicestatuList(params?: TableListParams) {
  return request(`/api/devicestatu/list?token=${getToken()}`, {
    params,
  });
}

export async function addDevicestatu(params: TableListParams) {
  return request(`/api/devicestatu?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeDevicestatu(params: { key: number[] }) {
  return request(`/api/devicestatu?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateDevicestatu(params: TablePutItem) {
  return request(`/api/devicestatu/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
