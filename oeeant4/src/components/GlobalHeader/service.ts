import request from '@/utils/request';
import { getToken } from '@/utils/authority';
import { TableListParams , TablePutItem } from './data.d';

export async function queryMessageList(params?: TableListParams) {
  return request(`/api/message/list?token=${getToken()}`, {
    params,
  });
}

export async function addMessage(params: TableListParams) {
  return request(`/api/message?token=${getToken()}`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeMessage(params: { key: number[] }) {
  return request(`/api/message?token=${getToken()}`, {
    method: 'DELETE',
    data: {
      ...params,
      method: 'delete',
    },
  });
}

export async function updateMessage(params: TablePutItem) {
  return request(`/api/message/${params.id}?token=${getToken()}`, {
    method: 'PUT',
    data: {
      ...params,
    },
  });
}
