// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';

const { REACT_APP_ENV } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  locale: {
    // default zh-CN
    default: 'zh-CN',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      component: '../layouts/UserLayout',
      routes: [
        {
          name: 'login',
          path: '/user/login',
          component: './user/login',
        },
        {
          name: '注册',
          path: '/user/register',
          component: './user/register',
        },
        {
          name: '重设密码',
          path: '/user/reset',
        //  component: './user/reset',
        },
      ],
    },
    {
      path: '/',
      component: '../layouts/SecurityLayout',
      routes: [
        {
          path: '/',
          component: '../layouts/BasicLayout',
          authority: ['admin', 'user'],
          routes: [
            {
              path: '/',
              redirect: '/oee/device',
            },
            {
              path: '/oee',
              name: 'OEE时间分析',
              routes: [
                {
                  path: '/oee/device',
                  name: '机器',
                  component: './oee/device',
                },
                {
                  path: '/oee/worktime',
                  name: '机台工作时间',
                  component: './oee/worktime',
                },
                {
                  path: '/oee/valve',
                  name: '点胶阀',
                  component: './oee/valve',
                },
                {
                  path: '/oee/valvetime',
                  name: '阀工作时间',
                  component: './oee/valvetime',
                },
              ]
            },
            {
              path: '/warn',
              name: '故障报警',
              routes: [
                {
                  path: '/warn/bug',
                  name: '故障',
                  component: './oee/bug',
                },
                {
                  path: '/warn/alarm',
                  name: '报警',
                  component: './oee/alarm',
                },
                {
                  path: '/warn/message',
                  name: '通知',
                  component: './oee/message',
                },
              ]
            },
            {
              path: '/plan',
              name: '项目管理',
              routes: [
                {
                  path: '/plan/plan',
                  name: '项目进度',
                  component: './oee/plan',
                },
                {
                  path: '/plan/project',
                  name: '项目分类',
                  component: './oee/project',
                },
              ]
            },
            {
              path: '/arg',
              name: '参数设置',
              routes: [
                {
                  path: '/arg/valvetype',
                  name: '阀类型',
                  component: './oee/valvetype',
                },
                {
                  path: '/arg/bugtype',
                  name: '故障类型',
                  component: './oee/bugtype',
                },
                {
                  path: '/arg/alarmtype',
                  name: '报警',
                  component: './oee/alarmtype',
                },
                {
                  path: '/arg/devicestatu',
                  name: '机台开关状态',
                  component: './oee/devicestatu',
                },
              ]
            },
            {
              path: '/user_system',
              name: '用户管a理',
              routes: [
                {
                  path: '/user_system/user',
                  name: '用户',
                  component: './user_system/user',
                },
                {
                  path: '/user_system/role',
                  name: '角色',
                  component: './user_system/role',
                },
                {
                  path: '/user_system/userlog',
                  name: '用户日志',
                  component: './user_system/userlog',
                },
              ]
            },
            {
              path: '/admin',
              name: 'admin',
              icon: 'crown',
              component: './Admin',
              authority: ['admin'],
              routes: [
                {
                  path: '/admin/sub-page',
                  name: 'sub-page',
                  icon: 'smile',
                  component: './Welcome',
                  authority: ['admin'],
                },
              ],
            },
            {
              component: './404',
            },
          ],
        },
        {
          component: './404',
        },
      ],
    },
    {
      component: './404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
});
