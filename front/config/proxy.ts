export default {
  dev: {
    '/api/': {
      target: 'http://localhost:8002/api/v1/oee',
      changeOrigin: true,
      pathRewrite: { '^/api': '' },
    },
    '/user/': {
      target: 'http://localhost:20216/api/v3/user',
      changeOrigin: true,
      pathRewrite: { '^/user': '' },
    },
  },
  test: {
    '/api/': {
      target: 'https://preview.pro.ant.design',
      changeOrigin: true,
      pathRewrite: { '^': '' },
    },
  },
  pre: {
    '/api/': {
      target: 'your pre url',
      changeOrigin: true,
      pathRewrite: { '^': '' },
    },
  },
};
