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
    - **接收数据** A组件想接收数据 则在A组件中给$bus绑定自定义事件 事件的 **回调留在A组件自身**
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
        - 最好在`beforeDestory`中用`$off()`去解绑 **当前组件用到的事件**

## 消息订阅与发布(pubsub-js第三方库实现)

- 一种组件间通信的方式 适用于**任意组件间通信**
    - **使用步骤：**
        1. 安装pubsub `npm i pubsub-js@1.6`
        2. 引入 `import pubsub from 'pubsub-js'`
            3. 接收数据：A组件想要接收数据 则在A组件中订阅消息 订阅的 **回调留在A组件自身**
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
    1. 优点：配置简单，请求资源时直接发给前端8080
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

## 插槽(slot)

- **作用：** 让父组件可以向子组件制定位置插入html结构 也是一种组件间通信方式 适用于 **父组件 -> 子组件**
- **分类：** 默认插槽 具名插槽 作用域插槽
- 使用方式
1. 默认插槽
    ```html
    父组件中
    <MyCategory title="美食">
        <img :src="this.$imgAddr" alt="">
    </MyCategory>
    子组件中
    <template>
        <div>
        <!--定义插槽-->
        <slot>插槽默认内容</slot>
        </div>
    </template>
    ```
    -
        2. 具名插槽
    ```html
    父组件中
    <MyCategory slot="center">
        <div>html结构</div>
    </MyCategory>
    <MyCategory v-slot:center>
        <div>html结构</div>
    </MyCategory>
    子组件中
    <template>
        <div>
        <!--定义插槽-->
        <slot name="center">插槽默认内容1</slot>
        <slot name="footer">插槽默认内容2</slot>
        </div>
    </template>
    ```

    -
        3. 作用域插槽

        - **理解：** <p style="color: red">数据在组件的自身 但根据数据生成的结构需要组件的使用者来决定</p>
        - **具体使用方法：**
          ```html
          父组件中
          <MyCategory title="游戏">
             <!--作用域插槽必须要使用template标签和scoped-->
           <template slot-scope="aaa">
             <ul>
               <li v-for="(item,index) in aaa.games" :key="index">{{item}}</li>
             </ul>
           </template>
          </MyCategory>
    
          <MyCategory title="游戏">
          <template slot-scope="aaa">
            <ol>
              <li v-for="(item,index) in aaa.games" :key="index">{{item}}</li>
            </ol>
          </template>>
          </MyCategory>
          子组件中
          <template>
            <div class="category">
                <h3>{{ title }}</h3>
            <slot :games="games" :msg="msg">我是一些默认值 如果没有传递默认值就会出现1</slot>
            </div>
          </template>
          ```
        - **ps：** 作用域插槽中也可以写name

## Vuex

- **概念：** 专门在Vue中实现集中式状态(数据)管理的一个插件 对Vue应用中多个组件的共享状态进行集中式的管理(r/w) 也是一种组件间通信
  且适用于任意组件间通信
- **Github地址：** [Vuex地址](https://github.com/vuejs/vuex)
- **什么时候使用：** (共享)
    - 多个组件依赖于同一状态
    - 来自不同组件的行为需要变同一状态
- **原理图：**
- <img src="https://vuex.vuejs.org/vuex.png" style="width: 60%; border-radius: 10px">

#### <p style="color: #0c72d2; font-weight: bold;">State</p>

- 这是一个对象 存放着数据 如 `{ todos: [], sum: 0, }`

#### <p style="color: #0c72d2; font-weight: bold;">Vue Components</p>

- 这是一个对象 使用`dispatch('jia', 2)`来操作数据

#### <p style="color: #0c72d2; font-weight: bold;">Action</p>

- 这也是一个对象 内容为键值对 恰好又一个方法与上边对应`{... jia: function(){}, ...}`
- 需要再自定义的函数中自己调用`commit('jia', 2)`进行提交

#### <p style="color: #0c72d2; font-weight: bold;">Mutations</p>

- 这也是一个Object对象 具有和上面commit中对应的函数 `{... jia:function() {}, ...}`
- `jia:function`中含有两个对象
    1. state
    2. 2 (具体加的值)
- `jia:function`中写入`state.sum += 2`会自动触发底层的Mutate
- 重新解析模版 再进行渲染

#### 为何Actions显得多余

- 假设执行`dispatch('chu', <值>)` 这里的值不确定 需要请求后端服务起来获得
- 发送Ajax请求来获取值 (图中 Backend API)
- 获得值来进行继续的计算
- **ps:** 如果知道具体的值 则允许跳过Actions直接调用commit

#### Devtools

- Vuex官方出的一个开发者调试工具
- 注意开发者只能直接与Mutations对话

#### <p style="font-weight: bold; color: #985f0d">store</p>

- 图中的State Actions Mutations 都需要经过 **store**进行管理
- 调用的`dispatch()`是由store提供的而不是window

### 搭建vuex环境

- 创建 `store/index.js`文件
    ```javascript
    import Vue from 'vue'
    // 该文件用于创建vuex中的核心store
    // 引入vuex
    import Vuex from 'vuex'
    // 准备Actions 用于响应组件中的作用
    const actions = {}
    // 准备Mutations 用于操作数据
    const mutations = {}
    // 准备State 用于存储数据
    const state = {}
    Vue.use(Vuex);
    // 创建Store
    // 创建并导出Store
    export default new Vuex.Store({
    actions,
    mutations,
    state,
    })
    ```
- 在`main.js`中引入
    ```javascript
    // 引入store 默认index.js那就不用写
    import store from './store'
    new Vue({
        // ...
        store,
        // ...
    }).$mount('#app')
    ```

### 基本使用

1. 初始化数据`state` 配置`actions` `mutations` 操作文件`index.js`
    ```javascript
    import Vue from 'vue'
    // 该文件用于创建vuex中的核心store
    // 引入vuex
    import Vuex from 'vuex'
    // 准备Actions 用于响应组件中的作用
    const actions = {
        // 定义在组件中加的动作 完整写法 jia: function() {}
        jia(context, value) {   // 这里使用简写
            // 在这里可以执行复杂的业务逻辑 如发送Ajax请求等
            // 如果要判读值的 使用 context.state.sum
            context.commit('JIA', value);   // JIA是mutations中的方法 改为大写方便区分方法属于哪个对象
        } 
    }
    // 准备Mutations 用于操作数据
    const mutations = {
        // 在这里才是对数据进行操作 注意开发者工具与此处对话 因此不要在其他地方修改数据 否则检测不到改变
        // 真正执行加
        JIA(state, value) { // state对象中有存入数据 并且匹配了get和set方法
            state.sum += value;
        }
    }
    // 准备State 用于存储数据
    const state = {
        sum: 0, // 定义初始化的数据
    }
    Vue.use(Vuex);
    // 创建Store
    // 创建并导出Store
    export default new Vuex.Store({
    actions,
    mutations,
    state,
    })
    ```
2. 组件中读取vuex中的数据：`$store.state.sum`
3. 组件中修改vuex中的数据：
    1. 无明确方法值需要中间操作的：`$store.dispatch('<actions中的方法名>', <数据>)`
    2. 有明确的值可跳过dispatch直接调用commit：`$store.commit('<mutations中的方法名>', <数据>)`

### 四个map方法的使用

1. **mapState方法：** 用于映射`state`中的数据为 **计算属性**
    ```javascript
    computed: {
        // 对象写法
        ...mapState({he: 'sum', xuexiao: 'school', xueke: 'subject',}),
        // 数组写法
        ...mapState(['sum', 'school', 'subject']),
   }
    ```
2. **mapGetter方法：** 用于映射`getters`中的数据为 **计算属性**
    ```javascript
    computed: {
        // 对象写法
        ...mapGetters({bigSum: 'bigSum'}),
        // 数组写法
        ...mapGetters(['bigSum']),
   }
    ```
3. **mapActions方法：** 用于生成与`actions`对话的方法 (包含`$store.dispatch(xxx)`的**函数**)
    ```javascript
    computed: {
        // 对象写法
        ...mapActions({incrementOdd: 'jiaOdd', incrementWait: 'jiaWait'}),
        // 数组写法
        ...mapActions(['jiaOdd', 'jiaWait']),
   }
    ```
4. **mapMutations方法：** 用于生成与`mutations`对话的方法 (包含`$store.commit(xxx)`的**函数**)
    ```javascript
    computed: {
        // 对象写法
        ...mapMutations({increment: 'JIA', decrement: 'JIAN'}),
        // 数组写法
        ...mapMutations(['JIA', 'JIAN']),
    ```

### 模块化+命名空间

- **目的：** 让代码更好维护 让多种数据分类更加明确
- 修改 `index.js`
    ```javascript
    const countAbout = {
        namespaced: true,   // 开启命名空间
        state: { x: 1, },
        mutations: { ... },
        actions: { ... },
        getters: {
            getSum(state) {
                return state.x * 10;
            }
        } 
    }
    const personAbout = {
        namespaced: true,   // 开启命名空间
        // ......
    }
    const store = new Vuex.Store({
        modules: {
            countAbout,
            personAbout
        }
   })
    ```

- 开启命名空间(namespaced)后 组件中读取state数据
    ```javascript
    // 方式1 直接自己读取
    this.$store.state.personAbout.personList
    // 方式2 借助mapState读取
    ...mapState('countAbout', ['sum', 'school', 'subject'])
    ```
- 开启命名空间(namespaced)后 组件中读取getters数据
    ```javascript
    // 方式1 直接自己读取
    this.$store.getters['personAbout/firstPersonName']
    // 方式2 借助mapGetters读取
    ...mapGetters('countAbout', ['bigSum'])
    ```
- 开启命名空间(namespaced)后 组件中调用dispatch
    ```javascript
    // 方式1 直接的dispatch
    this.$store.dispatch('personAbout/addPersonWang', person)
    // 方式2 借助mapActions
    ...mapActions('countAbout', {incrementOdd: 'jiaOdd', incrementWait: 'jiaWait'})
    ```
- 开启命名空间(namespaced)后 组件中调用dispatch
    ```javascript
    // 方式1 直接的commit
    this.$store.commit('personAbout/ADD_PERSON', person)
    // 方式2 借助mapMutations
    ...mapMutations('countAbout', {increment: 'jia', decrement: 'jian'})
    ```

## 路由(Vue router)

- **理解：** vue的一个插件库 专门用来实现SPA应用
- **什么是SPA应用：**
    1. 单页Web应用 (Single page application, SPA)
    2. 整个页面只有 **一个完整的页面**
    3. 点击页面中的导航链接 **不会刷新** 页面 只会做页面的 **局部更新**
    4. 数据需要通过ajax请求获取
- **路由的理解：**
    1. *什么是路由*
        - 一个路由就是一组映射关系 (key-value)
        - key为路径 value可能是function或component
    2. *路由的分类*
        1. 后端路由：
            - *理解：* value是function 用于处理客户端提交的请求
            - *工作过程：* 服务器收到一个请求时 根据**请求路径**找到匹配的**函数**来处理请求 返回响应数据
        2. 前端路由：
            - *理解：* value是component 用于展示页面内容
            - *工作过程：** 当浏览器的路径改变时 对应组件就会显示

### 路由的基本使用
1. 创建一个路由 (位于 `src/router/index.js` )
    ```javascript
    // 该文件用于创建整个应用的路由器
    import VueRouter from 'vue-router'
    // 引入组件
    import MyAbout from '../pages/About.vue'
    import MyHome from '../pages/Home.vue'
    // 创建并暴露
    export default new VueRouter({
        routes: [{
                path: '/about',
                component: MyAbout,
             },
             {
                path: '/home',
                component: MyHome,
             },
        ],
    })
    ```
2. 实现切换
    ```html
    <router-link class="list-group-item" active-class="active" to="/about">About</router-link>
    ```
3. 指定展示位置
    ```html
    <router-view></router-view>
    ```
### 几个注意点
1. 路由组件通常放在`pages`文件夹 一般组件一般放在`components`中
2. 通过切换 “隐藏”了路由组件 默认是被销毁掉的 需要的时候再进行挂载
3. 每个组件都有自己`$route`属性 里面存储着自己的路由信息
4. 真个应用只有一个router 可以通过组件的`$router`属性获取到 

### 嵌套(多级)路由
1. 配置路由规则 使用children配置项
    ```javascript
    export default new VueRouter({
        routes: [
            // 一级路由
            {
                path: '/home',
                component: MyHome,
                // 二级路由
                children: [
                    {
                        path: 'message',    // 二级路径不要写/
                        component: MyMessage,
                    },
                    {
                        path: 'news',
                        component: MyNews,
                    }
                ],
    
            },
    
        ],
    })
    ```
2. 跳转 (要写完整的路径)
    ```html
    <router-link to="/home/message">Message</router-link>
    ```

### 路由的query参数
1. 传递参数
    ```html
    <li v-for="m in messageList" :key="m.id">
        <!--        跳转路由并携带query参数 to的字符串写法-->
        <!--        <router-link :to=" `/home/message/detail?id=${m.id}&title=${m.title}` ">{{ m.title }}</router-link>-->
        <!--        to的对象写法-->
        <router-link :to="{
            path: '/home/message/detail',
            query: {
                id: m.id,
                title: m.title,
            },}">
            {{ m.title }}
        </router-link>
    </li>
    ```
2. 接收参数
    ```javascript
    $route.query.id
    $route.query.title
    ```
   
### 命名路由
- **作用：** 可以简化路由跳转的参数
- **使用：**
    1. 给路由命名
    ```javascript
    name: 'xiangqing',  // 起名
    path: 'detail',
    component: MyDetail,
    ```
    2. 简化跳转
    ```html
    <!--简化前-->
    <router-link to="/demo/test/welcome">跳转</router-link>
    <!--简化后-->
    <router-link :to="{name: 'hello'}"">跳转</router-link>
    <!--简化写法配合传递参数-->
    <router-link :to="{
        name: 'hello',
        query: {
            id: 666,
            title: '你好',
        }}">
    </router-link>
    ```
  
### 路由的params参数
1. **配置路由：** 声明接受params参数
    ```javascript
        name: 'xiangqing',  // 起名
        path: 'detail/:id/:title',  // 使用占位符声明接收params参数
        component: MyDetail,
    ```
2. 传递参数
    ```html
    <!--跳转并携带params参数 to的字符串写法-->
    <router-link to="/home/message/detail/666/你好">跳转</router-link>
    <!--to的对象写法-->
    <router-link :to="{
              // path: '/home/message/detail',
              // 注意如果传递params 那么这里不能使用path 只能用name
              name: 'xiangqing',
              params: {
                id: m.id,
                title: m.title,
              }
            }">
        {{ m.title }}
    </router-link>
    ```
3. **ps：** 路由携带params参数时 若使用to的对象写法 则不能使用path配置项 必须使用name配置项 
4. 接收参数
     ```html
     $route.params.id
     $route.params.title
     ```
   
### 路由的props配置
- **作用：** 让路由组件更方便的收到参数
    ```javascript
    name: 'xiangqing',  // 起名
    path: 'detail/:id/:title',
    component: MyDetail,
    // props的第一种写法 值为对象 该对象中所有的key value都会以props的形式传递给Detail组件
    // props: {a: 1, b: 'hellohello'}
    // 第二种写法 如果是true 就会把该路由组件收到的所有params参数 以props的形式传给Detail
    // props: true,
    // 第三种写法 值为函数
    props($route) {
    return {id: $route.query.id, title: $route.query.title};
    }
    ```
  
### `<router-link>`的replace属性
- 模拟栈的操作
- 默认为push 记录每一次的记录
- 加入 `replace` 属性则会替换当前操作

### 编程式路由导航
- **作用：** 不借助 `<router-link>` 实现路由跳转 让路由跳转更加灵活
- **使用：**
    ```javascript
    // 路由跳转的两个Api
    this.$router.push({
        name: 'xiangqing',
        params: {id: 'xxx', title: 'yyy'},
    })
    
    this.$router.replace({
        name: 'xiangqing',
        params: {id: 'xxx', title: 'yyy'},
    })
    
    this.$router.forward(); // 前进
    this.$router.back();    // 后退
    this.$router.go(3);     // 可前可后 根据传入的值 负数后退 正数前进
    ```
  
### 缓存路由组件
- **作用：** 让不展示的路由组件保持挂载 不被销毁
- **具体编码：**
    ```html
    <!--    <keep-alive include="MyNews">-->
    <!--    以上为缓存一个 下面的为缓存多个-->
    <keep-alive :include="['MyNews', 'MyMessage']">
    <!--    注意上边include的名字是组件中的name属性-->
        <router-view></router-view>
    </keep-alive>
    ```
  
### 两个新的声明周期钩子
- **作用：** 路由组件独有的两个钩子 用于捕获路由组件的激活状态
- **具体名字：**
  1. `activated() {}` 路由组件被激活时触发
  2. `deactivated() {}` 路由组件失活时触发

### 路由守卫
- **作用：** 对路由进行权限控制
- **分类：** 全局守卫、独享守卫、组件内守卫
- 全局守卫：
    ```javascript
    // 全局前置路由守卫
    // 每一次路由切换之前被调用 和初始化前被调用
    router.beforeEach((to, from, next) => {
        console.log('全局前置路由守卫', to, from);
        if (to.meta.isAuth) {
            if (localStorage.getItem('school') === 'NNU')
                next();
            else
                alert('学校名错误')
        } else {
            next();
        }
    })
    
    // 全局后置路由守卫
    // 每一次路由切换之后被调用 和初始化后被调用
    router.afterEach((to, from) => {
        console.log('后局前置路由守卫', to, from);
        document.title = to.meta.title;
    })
    ```
- 独享守卫 (**ps:** 独享守卫没有退出)
    ```javascript
    beforeEnter(to, from, next) {
        if (to.meta.isAuth) {
            if (localStorage.getItem('school') === 'NNU') {
                next()
            } else {
                alert('学校名不可用')
            }
        } else {
            next();
        }
    },
    ``` 
- 组件内守卫 (**ps:** 必须通过路由规则进入组件)
    ```javascript
    // beforeRouteEnter 通过路由规则进入组件时被调用
    beforeRouteEnter(to, from, next) {
        console.log('独享路由守卫-进入-About---beforeRouteEnter');
        console.log(to, from, next);
        if (to.meta.isAuth) {
            if (localStorage.getItem('school') === 'NNU') {
                next();
            }
        } else {
            next();
        }
    },
    // beforeRouteLeave 通过路由规则离开组件时被调用
    beforeRouteLeave(to, from, next) {
        console.log('独享路由守卫-离开-About---beforeRouteLeave');
        console.log(to, from, next);
        next();
    }
    ```
  
### 路由的两种工作模式
1. 对于一个url来说 #后面的就是哈希值
2. hash值不会包含在http请求中 (哈希值不会带给服务器)
3. Hash模式
   1. 地址中永远带着#号 不美观
   2. 若以后将地址通过第三方手机app分享 若app检验严格 则地址会被标记为不合法
   3. 兼容性较好
4. History模式
   1. 地址干净美观
   2. 兼容性和Hash模式相比略差
   3. 应用部署上线时需要后端人员支持 解决刷新页面服务端404的问题

## UI组件库
### 移动端常用
1. Vant [链接](http://youzan.github.io/vant)
2. Cube UI [链接](http://didi.github.io/cube-ui)
3. Mint UI [链接](http://mint-ui.github.io)

### PC端常用
1. Element UI [链接](http://element.eleme.cn)
2. IView UI [链接](http://www.iviewui.com)
