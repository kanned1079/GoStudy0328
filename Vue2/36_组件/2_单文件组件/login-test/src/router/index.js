import Vue from 'vue'
import VueRouter from 'vue-router'
// 引入页面
import Home from "@/views/Home.vue";
import User from "@/views/User.vue"
import Main from "@/views/Main.vue"
import Mall from "@/views/Mall.vue"
import PageOne from "@/views/PageOne.vue"
import PageTwo from "@/views/PageTwo.vue"

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        component: Main,
        redirect: '/home',  // 重定向到首页
        children: [
            {
                path: 'home',   // 主页
                name: 'home',
                component: Home,
            },
            {
                path: 'user',   // 用户管理
                name: 'user',
                component: User,
            },
            {
                path: 'mall',   // 商品管理
                name: 'mall',
                component: Mall,
            },
            {
                path: 'page1',   // 页面1
                name: 'page1',
                component: PageOne,
            },
            {
                path: 'page2',   // 页面2
                name: 'page2',
                component: PageTwo,
            },

        ]
    },

]

const router = new VueRouter({
    mode: 'hash',
    routes,
})

export default router