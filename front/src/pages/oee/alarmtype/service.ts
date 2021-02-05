import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryAlarmtypeList(params?: TableListParams) {
  return request(`/api/alarmtype/list?token=${getToken()}`, {
    params,
  });
}

export async function addAlarmtype(params: TableListParams) {
  return request(`/api/alarmtype?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeAlarmtype(params: { key: number[] }) {
  return request(`/api/alarmtype?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateAlarmtype(params: TablePutItem) {
  return request(`/api/alarmtype/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
