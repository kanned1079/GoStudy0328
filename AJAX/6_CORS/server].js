const express = require('express');
const app = new express();

app.all('/cors', (request, response) => {
    response.setHeader('Access-Control-Allow-Origin', '*');
    response.send('Hello CORS.');
});

app.listen(8080, () => {
    console.log("Server started ...");
})