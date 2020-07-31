const path = require('path');

module.exports = {
  outputDir: path.resolve(__dirname, './build/public'),
  indexPath: path.resolve(__dirname, './build/index.html'),
  configureWebpack: {
    devServer: {
      proxy: {
        '/api': {
          target: process.env.DEV_BACKEND_URL,
          changeOrigin: true
        },
        '/uploads': {
          target: process.env.DEV_BACKEND_URL
        },
      }
    }
  }
};
