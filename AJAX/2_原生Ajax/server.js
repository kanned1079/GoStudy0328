const express = require('express');
const app = express();

app.get('/server', function (request, response) {
    // 设置相应头
    // 设置允许跨域
    response.setHeader('Access-Control-Allow-Origin', '*');
    // 设置相应体
    response.send('GET: Hello from Express');
});

app.post('/server', (request, response) => {
    response.setHeader('Access-Control-Allow-Origin', '*');
    response.setHeader('Access-Control-Allow-Headers', '*');
    response.send('POST: Hello from Express');

});

app.listen('8000', function () {
    console.log('Server started, listening 8000 port...');
})