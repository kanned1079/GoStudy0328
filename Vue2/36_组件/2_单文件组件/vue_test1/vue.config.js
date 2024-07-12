const {defineConfig} = require('@vue/cli-service')
const {fastKey} = require("core-js/internals/internal-metadata");
module.exports = defineConfig({
    transpileDependencies: true,
    lintOnSave: false,


})

// 方式1
// 开启代理服务器
// module.exports = {
//   devServer: {
//     proxy: 'http://localhost:5000'
//   }
// }

// 方式2
// 可以指定多个代理
module.exports = {
    devServer: {
        proxy: {
            '/prefix1': {  // 这个叫请求前缀 加在axios端口号后边的
                target: 'http://localhost:5001',
                pathRewrite: {'^/prefix1': ''}, // 将匹配所有</prefix1>都去掉
                ws: true,   // 用于支持WebSocket    默认值为true 但是在React中默认为false
                changeOrigin: true, // 用于控制请求头中的host值
            },
            '/prefix2': {
                target: 'http://localhost:5002',
                pathRewrite: {'^/prefix2': ''},
                ws: true,
                changeOrigin: true,
            }
        },

    }
}

