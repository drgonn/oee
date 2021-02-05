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
              redirect: '/oee/machine',
            },
            {
              path: '/oee',
              name: 'OEE时间分析',
              routes: [
                {
                  path: '/oee/machine',
                  name: '监控台',
                  routes:[
                    {
                      path: '/oee/machine',
                      component: './oee/machine',
                    },
                    {
                      path: '/oee/machine/detail',
                      component: './oee/machine/device_detail',
                    },
                  ],
                },
                {
                  path: '/oee/device',
                  name: '机器',
                    routes:[
                      {
                        path: '/oee/device',
                        component: './oee/device',
                      },
                      {
                        path: '/oee/device/detail',
                        component: './oee/device_detail',
                      },
                      ]
                },
                {
                  path: '/oee/worktime',
                  name: '工作时间',
                  component: './oee/worktime',
                },
                // {
                //   path: '/oee/device/detail',
                //   component: './oee/device_detail',
                //   name:'机器详情',
                //   hideInMenu: true,
                // },
                {
                  path: '/oee/valve',
                  name: '点胶阀',
                  component: './oee/valve',
                },
                {
                  path: '/oee/bug',
                  name: '故障',
                  component: './oee/bug',
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
                  component: './plan/plan',
                },
                {
                  path: '/plan/project',
                  name: '项目分类',
                  component: './plan/project',
                },
              ]
            },
            {
              path: '/user_system',
              name: '用户管理',
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
