// 配置路由
import Vue from 'vue'
import VueRouter from 'vue-router'

// 使用插件
Vue.use(VueRouter)

// 引入路由组件
import Home from '@/pages/Home'
import Search from '@/pages/Search'
import Register from '@/pages/Register'
import Login from '@/pages/Login'

export default new VueRouter({
    // 配置路由
    routes: [
        {
            path: '/home',
            component: Home,
            meta: {show: true}
        },
        {
            name: 'search',
            path: '/search/:keyWord',
            component: Search,
            meta: {show: true},
            props: true,

        },
        {
            path: '/register',
            component: Register,
            meta: {show: false}

        },
        {
            path: '/login',
            component: Login,
            meta: {show: false}

        },
        // 重定向 在运行的时候 定向到首页
        {
            path: '*',
            redirect: '/home',
        }
    ]
})