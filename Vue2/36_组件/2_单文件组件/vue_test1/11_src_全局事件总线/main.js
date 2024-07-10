import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;

// Vue.prototype.x = {a: 1, b:2}

// const Demo = Vue.extend({})
// const d = new Demo()
//
// Vue.prototype.x = d

new Vue({
    // el: '#root',
    render: createElement => createElement(App),
    beforeCreate() {
        // 当前的this就是Vue 还没有开始解析模版
        Vue.prototype.$bus = this; // 安装全局事件总线
    }
}).$mount('#app');