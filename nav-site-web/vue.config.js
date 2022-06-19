module.exports = {
    configureWebpack: {
        module: {
            rules: [
                {
                    include: /node_modules/,
                    test: /\.mjs$/,
                    type: 'javascript/auto'
                }
            ]
        }
    },

    devServer: {
        proxy: {
            '/api/': {
                target: 'http://127.0.0.1:8083/',
                changeOrigin: true,
                ws: false,
                pathRewrite: {
                    '^/api/': '/'
                }
            }
        }
    }
}