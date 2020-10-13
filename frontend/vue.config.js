const path = require('path');


module.exports = {
    outputDir: path.resolve(__dirname, './build/public'),
    indexPath: path.resolve(__dirname, './build/index.html'),
    configureWebpack: {
        devServer: {
            proxy: {
                '/api/users': {
                    target: process.env.DEV_SERVICE_USERS_URL,
                    changeOrigin: true
                },
                '/api/chat': {
                    target: process.env.DEV_SERVICE_CHAT_URL
                },
            }
        },

        resolve: {
            extensions: ['*', '.js', '.vue', '.json'],
            alias: {
                '@': path.resolve(__dirname, './src'),
                utils: path.resolve(__dirname, './src/utils'),
                api: path.resolve(__dirname, './src/api'),
                assets: path.resolve(__dirname, './src/assets'),
                http: path.resolve(__dirname, './src/http'),
                pages: path.resolve(__dirname, './src/pages'),
                plugins: path.resolve(__dirname, './src/plugins'),
                router: path.resolve(__dirname, './src/router'),
                store: path.resolve(__dirname, './src/store'),
                components: path.resolve(__dirname, './src/components'),
            }
        },
    }
};
