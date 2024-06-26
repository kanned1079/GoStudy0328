// 先引入express
const express = require('express');

// 创建应用对象
const app = express();

// 创建路由规则
// request是对请求报文的封装
// response是对相应报文的封装
app.get('/', (request, response) => {
    response.send("Hello from express")
});

app.listen(8000, ()=>{
    console.log("服务器已经启动，监听8000端口...");
})