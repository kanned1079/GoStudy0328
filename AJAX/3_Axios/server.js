const express = require('express');
const app = new express();

// 中间件处理预检信息
app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
    res.header('Access-Control-Allow-Headers', 'Content-Type, Authorization');

    if (req.method === 'OPTIONS') {
        return res.sendStatus(200);
    }
    next();
});

app.get('/axios1', function (request, response) {
    response.setHeader('Access-Control-Allow-Origin', '*');
    response.send('GET: Hello from Express');
});

app.post('/axios2', () => {
    response.setHeader('Access-Control-Allow-Origin', '*');
    response.send('POST: Hello from Express');
})

app.listen('9000', () => {
    console.log('Server started...');
})