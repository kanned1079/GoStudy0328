const fs = require('fs');

// 使用回调地狱方式
// fs.readFile('./txt1.txt', (err, data1) => {
//     fs.readFile('./txt2.txt', (err, data2) => {
//         fs.readFile('./txt3.txt', (err, data3) => {
//             let result = data1 +'\r\n'+ data2 +'\r\n'+ data3;
//             console.log(result);
//         });
//     });
// });

// 使用Promise实现
const p = new Promise((resolve, reject) => {
    fs.readFile('./txt1.txt', (err, data) => {
        resolve(data);
    });
});

p.then(value => {
    return new Promise((resolve, reject) => {
        fs.readFile('./txt2.txt', (err, data) => {
            resolve([value, data]);
            // 这里将是前面两个文件的值 那么也就是首部then的返回值
        });
    });
}).then(value => {
    return new Promise((resolve, reject) => {
        fs.readFile('./txt3.txt', (err, data) => {
            value.push(data);
            resolve(value);
            // 这里将是前面两个文件的值 那么也就是首部then的返回值
        });
    });
}).then(value => {
    console.log(value.join('\r\n'));
})


