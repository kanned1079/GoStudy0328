// 该文件用于创建整个应用的路由器
import VueRouter from 'vue-router'
// 引入组件
import MyAbout from '../pages/About.vue'
import MyHome from '../pages/Home.vue'
import MyMessage from '../pages/Message.vue'
import MyNews from '../pages/News.vue'
import MyDetail from '../pages/Detail.vue'

// 创建一个路由器
// 创建并暴露
export default new VueRouter({
    routes: [
        // 一级路由
        {
            name: 'guanyu', // 起名 可以简化跳转编码
            path: '/about',
            component: MyAbout,
        },
        {
            path: '/home',
            component: MyHome,
            // 二级路由
            children: [
                {
                    path: 'message',
                    component: MyMessage,
                    children: [
                        {
                            name: 'xiangqing',  // 起名
                            path: 'detail/:id/:title',
                            component: MyDetail,
                            // props的第一种写法 值为对象 该对象中所有的key value都会以props的形式传递给Detail组件
                            // props: {a: 1, b: 'hellohello'}
                            // 第二种写法 如果是true 就会把该路由组件收到的所有params参数 以props的形式传给Detail
                            // props: true,
                            // 第三种写法 值为函数
                            props($route) {
                                return {id: $route.query.id, title: $route.query.title};
                            }
                        }
                    ]
                },
                {
                    path: 'news',
                    component: MyNews,
                }
            ],

        },

    ],
})

// export default router