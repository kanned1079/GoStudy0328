import Vue from 'vue';
import App from './App.vue';
import {mixin_showName, hunhe2} from "@/mixin";

Vue.config.productionTip = false;

// 定义全局的混入(mixin)
Vue.mixin(mixin_showName);
Vue.mixin(hunhe2);

new Vue({
    // el: '#root',
    render: createElement => createElement(App)
}).$mount('#app');