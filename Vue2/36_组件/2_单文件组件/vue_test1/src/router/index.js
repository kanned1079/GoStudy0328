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
const router = new VueRouter({
    mode: 'history',    // 选择路由的工作模式
    routes: [
        // 一级路由
        {
            name: 'guanyu', // 起名 可以简化跳转编码
            path: '/about',
            component: MyAbout,
            meta: {title: '关于', isAuth: true}
        },
        {
            name: 'zhuye',
            path: '/home',
            component: MyHome,
            meta: {title: '主页'},
            // 二级路由
            children: [
                {
                    name: 'xiaoxi',
                    path: 'message',
                    component: MyMessage,
                    meta: {title: '消息'},
                    children: [
                        {
                            name: 'xiangqing',  // 起名
                            path: 'detail/:id/:title',
                            component: MyDetail,
                            // props的第一种写法 值为对象 该对象中所有的key value都会以props的形式传递给Detail组件
                            // props: {a: 1, b: 'hellohello'}
                            // 第二种写法 如果是true 就会把该路由组件收到的所有params参数 以props的形式传给Detail
                            props: true,
                            // 第三种写法 值为函数
                            // props($route) {
                            //     return {id: $route.query.id, title: $route.query.title};
                            // }
                            meta: {isAuth: true, title: '详情'}

                        }
                    ]
                },
                {
                    name: 'xinwen',
                    path: 'news',
                    component: MyNews,
                    meta: {isAuth: true,title: '新闻'},
                    // 组件独有的 独享路由守卫
                    // 注意独享路由守卫只有一个enter
                    // beforeEnter(to, from, next) {
                    //     if (to.meta.isAuth) {
                    //         if (localStorage.getItem('school') === 'NNU') {
                    //             next()
                    //         } else {
                    //             alert('学校名不可用')
                    //         }
                    //     } else {
                    //         next();
                    //     }
                    // },


                }
            ],

        },

    ],
})

// // 全局前置路由守卫
// // 每一次路由切换之前被调用 和初始化前被调用
// router.beforeEach((to, from, next) => {
//     console.log('全局前置路由守卫', to, from);
//     if (to.meta.isAuth) {
//         if (localStorage.getItem('school') === 'NNU')
//             next();
//         else
//             alert('学校名错误')
//     } else {
//         next();
//     }
// })
//
// 全局后置路由守卫
// 每一次路由切换之后被调用 和初始化后被调用
router.afterEach((to, from) => {
    console.log('后局前置路由守卫', to, from);
    document.title = to.meta.title;
})



export default router
