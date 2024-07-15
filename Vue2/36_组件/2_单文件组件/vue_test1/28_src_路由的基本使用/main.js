import Vue from 'vue'
import App from './App.vue'
// 引入VueRouter
import VueRouter from "vue-router";
// 引入自定义的路由器router
import router from './router/index'

Vue.config.productionTip = false

Vue.use(VueRouter)

new Vue({
  render: h => h(App),

  router,

  beforeMount() {
    Vue.prototype.$bus = this;
  }
}).$mount('#app')
