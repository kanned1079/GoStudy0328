import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;

new Vue({
    // el: '#root',
    render: createElement => createElement(App),
    mounted() {
        // setTimeout(() => {
        //     // 销毁
        //     this.$destroy();
        // }, 3000)
    }
}).$mount('#app');