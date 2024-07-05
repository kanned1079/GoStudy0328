// 该文件是整个项目的入口文件

// 引入vue
import Vue from 'vue'
// 引入App组件 它是所有组件的父组件
import App from './App.vue'

// 关闭生产提示
Vue.config.productionTip = false

new Vue({
    el: '#app',

    // 将App组件放入容器中中
  // render: h => h(App),
  //   components: {
  //       App
  //   }
  //   render(createElement) {
  //       console.log('666');
  //       return createElement('h1', 'Hello');
  //   }
    // 精简
    // Vue: 核心+模版解析器
    // render: createElement => createElement('h1', 'Hello')
    render: createElement => createElement(App)
})

console.log('888');