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
                target: 'http://localhost:8083/',
                changeOrigin: true,
                ws: false,
                pathRewrite: {
                    // '^/api/': '/'
                }
            }
        }
    }
}