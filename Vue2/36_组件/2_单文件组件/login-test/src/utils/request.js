import Axios from 'axios'

const http = Axios.create({
    // baseURL: 'http://localhost:5000/api',
    baseURL: 'http://localhost:5001/api',
    timeout: 10000,
    headers: {
        //'Access-Control-Allow-Headers': '*',
    }
})

// 添加请求拦截器
http.interceptors.request.use(function (config) {
    // 在发送请求前做什么
    return config;
}, function (error) {
    // 对请求错误做什么
    return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(response => {
    // 对响应数据做点什么
    return response;
}, error => {
    // 对响应错误做点什么
    return Promise.reject(error);
});

export default http