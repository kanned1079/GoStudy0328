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
}).$mount('#app');