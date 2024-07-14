import Vue from 'vue'
import App from './App.vue'

// 引入vue-resource
import vueResource from 'vue-resource'
// 引入vuex
// import Vuex from 'vuex'
// 引入store 默认index.js那就不用写
import store from './store'

Vue.config.productionTip = false

// 使用插件
Vue.use(vueResource);
// Vue.use(Vuex);
// 创建vm的时候就可以传入store配置项

new Vue({
  render: h => h(App),
  store,
  beforeMount() {
    Vue.prototype.$bus = this;
  }
}).$mount('#app')
