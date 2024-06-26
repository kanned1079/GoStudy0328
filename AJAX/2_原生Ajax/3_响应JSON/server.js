const express = require('express');
const {response} = require("express");
const app = new express();

app.post('/json-server', () => {
    response.setHeader('Access-Control-Allow-Origin', '*');
    // 例如这里有一个对象
    let user = {
        name: 'kinggyo',
    }
    // 将对象序列化
    let str = JSON.stringify(user);
    response.send(str);

});

app.listen('8000', () => {
    console.log("Server started, listening 8000 port...")
})