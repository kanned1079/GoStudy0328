import Vue from 'vue'
import App from './App.vue'

// 引入vue-resource
import vueResource from 'vue-resource'

Vue.config.productionTip = false

// 使用插件
Vue.use(vueResource);

new Vue({
  render: h => h(App),
  beforeMount() {
    Vue.prototype.$bus = this;
  }
}).$mount('#app')
