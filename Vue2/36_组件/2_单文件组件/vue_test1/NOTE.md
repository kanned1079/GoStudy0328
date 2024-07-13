# 笔记

## ref属性

1. 被用来给元素或组件注册引用信息（id的替代者）
2. 应用在htm1标签上获取的是真实DOM元素，应用在组件标签上是组件实例对象(vc)
3. **使用方式:**
    - 打标识: <hl ref"xxx">.....</h1> 或 <School ref "xxx"></School>
    - 获取: this.$refs.xxx

## props配置

1. 传递数据  `<Demo name></Demo>`
   2. 接收数据
       1. 只接收 `props['name']`
       2. 限制类型 `props: {name: Number}`
       3. 限制类型、必要性、c 
       ```javascript
       props: { 
            name: {
                type: String,       //限制类型
                required: true,     //必要性
                default: 'kanna'    //必要性
            }
        }
        ```
3. mark: props是只读的 vue底层会监测你对props的修改 如果进行了修改会触发警告 如果需要修改 需要拷贝一份来修改

## mixin(混入)

- **功能：** 可以把多个组件共用的配置提取成一个混入对象
- **使用方法：**
    1. 先定义mixin (如mixin_1.js) `{data() {...}, methods: {sayHello() {...}, ...}}`
    2. 使用mixin
        1. 全局混入 在 `main.js` 中 `Vue.mixin(xxx)`
        2. 局部混入 `mixin: []`

## 插件

- **功能：** 用于增强Vue
- **本质：** 包含install方法的一个对象 install的第一个参数是Vue 第二个以后的参数是插件使用者传递的数据
- **定义插件：** ``
    ```javascript
    对象.install = function(Vue, options) {
        Vue.filter(...), 
        Vue.directive('...', {...})}, 
        Vue.mixin(...), 
        Vue.prototype.$myMethod = function() {...}
    }
    ```
- **使用插件：** `Vue.use(<插件名>)`

## scoped样式

- **作用：** 让样式局部生效 防止冲突
- **写法：**  
    ```html
    <style scoped>
        /*...*/
    </style>
    ```

## 全局事件总线(GlobalEventBus)

- 一种组件间通信的方式 适用于任何组件间通信
  - 安装全局事件总线 
      ```javascript
      new Vuw({
        // ... 
            beforeCreate() {
            Vue.prototype.$bus = this;
        }, 
         // ...
    })
    ```
- 使用事件总线
    - **接收数据** A组件想接收数据 则在A组件中给$bus绑定自定义事件 事件的**回调留在A组件自身**
        ```javascript
        methods: { 
            demo(data) {
                // 一些数据
            }
        },
      // ......
        mounted() {
            this.$bus.$on('xxxx', this.demo)
        }
        
      ```
    - **提供数据** `this.$bus.$emit('xxxx', this.demo)`
        - 最好在`beforeDestory`中用`$off()`去解绑**当前组件用到的事件**

## 消息订阅与发布(pubsub-js第三方库实现)

- 一种组件间通信的方式 适用于**任意组件间通信**
  - **使用步骤：**
      1. 安装pubsub `npm i pubsub-js@1.6`
      2. 引入 `import pubsub from 'pubsub-js'`
         3. 接收数据：A组件想要接收数据 则在A组件中订阅消息 订阅的**回调留在A组件自身** 
           ```javascript
            methods: { 
                <回调函数>(msgName, data) {
            }}, 
            mounted() { 
                this.pid = pubsub.subscribe('<消息名/主题>, <回调函数名>)
            }
           ```
      4. 提供数据  `pubsub.publish('<消息名/主题>', <消息内容>)`

## nextTick方法

- **语法：** `this.$nextTick(<回调函数(用function)>)`
- **作用：** 在下一次DOM更新结束后执行其指定的回调
- **什么时候用：** 当数据改变后 要基于更新后的新的DOM进行某些操作时 要在nextTick所指定的回调函数中执行

## Vue脚手架配置代理

### 方法1

- 在vue.conf.js中加入以下配置
    ```javascript
    module.exports = {
        devServer: {
            proxy: 'http://localhost:5000'
        }
    }
    ```
    1. 优点：配置简单，请求资源时直接发给前端8080jike
    2. 缺点：不能配置多个代理 不能灵活的控制请求是否走代理
    3. 工作方式：若按照上述配置代理 当请求了不再前端的资源时 那么请求会转发给服务器（优先匹配前端资源）

### 方法2
- 编写vue.conf.js配置具体代理规则
    ```javascript
    module.exports = {
        devServer: {
            proxy: {
                '/prefix1': {  // 这个叫请求前缀 加在axios端口号后边的
                    target: 'http://localhost:5001',
                    pathRewrite: {'^/prefix1': ''}, // 将匹配所有</prefix1>都去掉
                    ws: true,   // 用于支持WebSocket    默认值为true 但是在React中默认为false
                    changeOrigin: true, // 用于控制请求头中的host值
                },
                '/prefix2': {
                    target: 'http://localhost:5002',
                    pathRewrite: {'^/prefix2': ''},
                    ws: true,
                    changeOrigin: true,
                }
            },
        }
    }
    ```
    1. 优点：可以配置多个代理 且可以灵活的控制请求是否走代理
    2. 缺点：配置略微繁琐 请求资源必须加前缀 