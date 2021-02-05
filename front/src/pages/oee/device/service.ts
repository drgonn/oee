import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryDeviceList(params?: TableListParams) {
  return request(`/api/device/list?token=${getToken()}`, {
    params,
  });
}

export async function addDevice(params: TableListParams) {
  return request(`/api/device?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeDevice(params: { key: number[] }) {
  return request(`/api/device?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateDevice(params: TablePutItem) {
  return request(`/api/device/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
