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
    Vue.prototype.$imgAddr = 'https://ikanned.com:24444/d/R730xd_SSD/13.jpeg';
    Vue.prototype.$videoAddr = 'https://ikanned.com:24444/d/R730xd_SSD/%E8%BD%AC%E6%9C%AC%E7%94%A8/%E8%AE%A1%E7%AE%97%E6%9C%BA%E5%9F%BA%E7%A1%80/%E8%AE%A1%E7%AE%97%E6%9C%BA%E5%9F%BA%E7%A1%80/%E9%A2%86%E6%AD%A3%E4%B8%93%E8%BD%AC%E6%9C%AC%EF%BC%88%E5%9F%B9%E8%AE%AD%E6%9C%BA%E6%9E%84%E8%A7%86%E9%A2%91office%E5%8F%AF%E5%BF%BD%E7%95%A5%EF%BC%89/%E8%AE%A1%E7%AE%97%E6%9C%BA%E8%A7%86%E9%A2%91/1.0%E5%AF%BC%E5%AD%A6.mp4';
  }
}).$mount('#app')
