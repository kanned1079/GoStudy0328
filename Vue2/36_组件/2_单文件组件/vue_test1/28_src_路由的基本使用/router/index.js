// 该文件用于创建整个应用的路由器
import VueRouter from 'vue-router'
// 引入组件
import MyAbout from '../components/About.vue'
import MyHome from '../components/Home.vue'

// 创建一个路由器
// 创建并暴露
export default new VueRouter({
    routes: [
        {
            path: '/about',
            component: MyAbout,
        },
        {
            path: '/home',
            component: MyHome,
        },
    ],
})

// export default router