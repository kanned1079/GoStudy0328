const express = require('express');
const app = express();

app.get('/home', (request, response) => {
    response.sendFile(__dirname + '/public/index.html');
});

app.listen(8080, () => {
    console.log('Server started...')
})