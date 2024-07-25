import http from '../utils/request'

// 请求首页数据
export const getData = () => {
    // 返回一个Promise对象
    return http.get('/home/getData', {})
}

// 返回用户列表
export const getUsers = ({params}) => {
    // console.log('getUsers.params ', params)
    console.log('页数：', params.page)
    console.log('限制条数：', params.limit)
    return http.get('/user/get', {
        params: {
            page: params.page,
            limit: params.limit,
        }
    })
}

// 新增用户
export const addUser = (data) => {
    console.log('addUser', data)
    return http.post('/user/add', data)
}

// 编辑用户
export const editUser = (data) => {
    return http.post('/user/edit', data)
}

// 删除用户
export const delUser = (data) => {
    return http.post('/user/del', data)
}