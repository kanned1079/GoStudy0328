import Vue from 'vue'
import App from './App.vue'

// 引入路由
import router from '@/router'
import VueRouter from "vue-router";

Vue.config.productionTip = false

// 先保存一份VueRouter原型对象
let originPush = VueRouter.prototype.push;
let originReplace = VueRouter.prototype.replace;

//重写 push / replace
VueRouter.prototype.push = function (location, resolve, reject) {
    // console.log(this);
    if (resolve && reject) {
        // call和apply的区别 都可以调用函数一次 都可以篡改函数的上下文一次
        // 不同点 call传递参数用<,>隔开 apply需要传递数组
        originPush.call(this, location, resolve, reject)
    } else {
        originPush.call(this, location, () => {
        }, () => {
        })
    }
}
VueRouter.prototype.replace = function (localtion, resolve, reject) {
    if (resolve && reject) {
        originReplace().call(this, location, resolve, reject)
    } else {
        originReplace.call(this, location, () => {
        }, () => {
        })
    }
}


new Vue({
    render: h => h(App),
    router,
}).$mount('#app')
