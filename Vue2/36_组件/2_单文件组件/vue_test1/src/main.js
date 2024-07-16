import Vue from 'vue'
import App from './App.vue'

// 完整引入
// 引入element-ui组件库
import ElementUI from 'element-ui';
// 引入element-ui组件库的全部样式
import 'element-ui/lib/theme-chalk/index.css';

// 按需引入
// import {Button, Row, DatePicker} from 'element-ui';
// Vue.component(Button.name, Button);
// Vue.component(Row.name, Row);
// Vue.component(DatePicker.name, DatePicker);

Vue.config.productionTip = false

// 应用element-ui
Vue.use(ElementUI)

new Vue({
  render: h => h(App),

  beforeMount() {
    Vue.prototype.$bus = this;
  }
}).$mount('#app')
