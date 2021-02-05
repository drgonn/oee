import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryAlarmList(params?: TableListParams) {
  return request(`/api/alarm/list?token=${getToken()}`, {
    params,
  });
}

export async function addAlarm(params: TableListParams) {
  return request(`/api/alarm?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeAlarm(params: { key: number[] }) {
  return request(`/api/alarm?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateAlarm(params: TablePutItem) {
  return request(`/api/alarm/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
