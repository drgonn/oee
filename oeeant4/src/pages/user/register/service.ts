import request from 'umi-request';
import { UserRegisterParams } from './index';

export async function fakeRegister(params: UserRegisterParams) {
  console.log(params)
  return request('/user/register', {
    method: 'POST',
    data: params,
  });
}
// export async function fakeRegister(params) {
//   return request(`/user/register`, {
//     method: 'POST',
//     data: {
//       ...params,
//     },
//   });
// }

export async function queryMsg(params) {
  console.log(params)
  return request(`/user/register/msg`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
export async function ResetPassword(params) {
  return request(`/user/register/password`, {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
