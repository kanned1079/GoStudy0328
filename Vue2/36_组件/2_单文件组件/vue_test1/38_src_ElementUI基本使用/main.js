import Vue from 'vue'
import App from './App.vue'

// 引入element-ui组件库
import ElementUI from 'element-ui';
// 引入element-ui组件库的全部样式
import 'element-ui/lib/theme-chalk/index.css';

Vue.config.productionTip = false

// 应用element-ui
Vue.use(ElementUI)

new Vue({
  render: h => h(App),

  beforeMount() {
    Vue.prototype.$bus = this;
  }
}).$mount('#app')
