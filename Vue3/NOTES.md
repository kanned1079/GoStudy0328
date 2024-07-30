# Vue3笔记

## 创建项目

### 基于Vite创建

- 终端运行 `npm create vue@latest`
- 加入TypeScript支持

#### 注意

- `Vite`项目中 `index.html`是入口文件 在项目的最外层
- 加载`index.html`后 Vite解析`<script type=""module" src="xxx">`指向的JavaScripts
- Vue3通过`createApp()`函数创建一个简单的实例

## Vue3核心语法

### 选项式API和组合式API
- `Vue2`的API设计是`选项式(Options)`风格
- `Vue3`的API设计是`组合式(Composition)`风格

### 选项式API的弊端
`Options`类型的API 数据、方法、计算属性等 是分散在`data` `methods` `computed`中的 如果想新增或者修改一个需求 就需要分别修改
不利于维护
```html
<script lang="ts">
    export default {
        name: 'Person',
        data() { return { ... }},
        methods: { ... },
        computed: { ... },
        watch: { ... }
    }
</script>

<template>
    <div class="person">
        <h2>姓名： {{ name }}</h2>
        <h2>年龄： {{ age }}</h2>
        <button @click="modifyName">修改名字</button>
        <button @click="modifyAge">修改年龄</button>
        <button @click="showTel">显示联系方式</button>
    </div>
</template>
```

### 拉开序幕的 `setup`
`setup`是Vue3中一个新的配置项 值是一个函数 它是`Composition API`"表演的舞台" 组件中所用到的 数据 方法 计算属性 监视
均配置在`setup`中

#### 特点如下
- `setup`函数返回对西那个中的内容 可直接在模版中使用
- `setup`中访问`this`是`undefined`
- `setup`函数会在`beforeCreate`中调用 它是领先所有hooks执行的

#### `setup`语法糖
- `setup`有一个语法糖 可以让我们把`setup`独立出去
    ```html
    <script lang="ts" setup>
        console.log(this);  // undefined
        let x = 99;
    </script>
    ```
- 如果使用上述写法 还需要写一个不带`setup`的标签来设置name 因此可以借助一个插件来实现简化

##### 使用`vite-plugin-vue-setup-extend`插件
1. 安装插件 `yarn add vite-plugin-vue-setup-extend -D`
2. 编辑 `vite.config.ts` 加入两行内容
    ```typescript
    import {fileURLToPath} from 'node:url'
    
    import {defineConfig} from 'vite'
    import vue from '@vitejs/plugin-vue'
    // 第一步导入插件
    import VueSetupExtend from 'vite-plugin-vue-setup-extend'
    
    // https://vitejs.dev/config/
    export default defineConfig({
        plugins: [
            vue(),
            VueSetupExtend(), // 第二步使用插件
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        }
    })
    ```
3. 在标签中加入`name`属性即可
    ```html
    <script lang="ts" setup name="Person234">
     ... 
    </script>
     ```

### ref创建: 基本类型的响应式数据
- **作用：** 定义响应式变量(基本/对象都可)
- **语法：** `let <变量名> = ref(<初始值>);`
- **返回值：** 一个`RefImpl`实例对象 对象的`value`属性是响应式的
- **注意点：**
    - `JS`中操作数据需要`<变量>.value` 但是模版中不需要`.value`
    - 对于`let name = ref('张三');`来说 `name`不是响应式的 `name.value`是响应式的

### reactive创建: 对象类型的响应式数据
- **作用：** 定义响应式变量(只能用在对象中)
- **语法：** `let <变量名> = reactive({<源对象>});`
- **返回值：** 一个`Proxy`实例对象 简称`响应式对象`
- **注意点：**
    - `reactive`中定义的响应式数据是"深层次"的

### `ref`对比`reactive`
#### 宏观角度
> 1. `ref`用来定义 **基本数据类型**、**对象类型数据**
> 2. `reactive`用来定义 **对象类型数据**
#### 区别
> 1. `ref`创建的变量必须使用`.value`(可以使用volar插件自动添加.value)
> 2. `reactive`重新分配一个对象时 会**失去**响应式 (可以使用`Object.assign`去整体替换)
```typescript
// reactive的情况
car = {brand: '奥拓', price: 12}  // 这就叫重新分配一个对象 会因此失去响应式
Object.assign(car, {brand: '奥拓', price: 12})  // 可以使用Object.assign
// ref的情况 可以直接使用
car.value = {brand: '奥拓', price: 12}
```
#### 使用原则
> 1. 若需要一个基本类型的响应式数据 必须使用`ref`
> 2. 若需要一个响应式对象 层级不深 `ref`和`reactive`都可以
> 3. 若需要一个响应式对象 且层级较深 推荐使用`reactive`

### `toRefs`和`toRef`
#### toRefs函数
- 可以将一个响应式对象转换为一个普通的对象 该对象的每个属性都是独立的`ref`对象
- 返回的对象可以进行解构 每个属性都可以像普通的`ref`对象一样访问和修改，而且会保持响应式的关联
- 使用场景主要是在将响应式对象作为属性传递给子组件时 确保子组件可以正确地访问和更新这些属性
    ```typescript
    let {name, age} = toRefs(person)    // 此时name age都是响应式的ref
    ```
#### toRef函数
- 可以将一个响应式对象的属性转换为一个独立的`ref`对象
- 返回的是一个指向源对象属性的`ref`引用 任何对该引用的修改都会同步到源对象属性上。
- 使用`toRef`时需要传入源对象和属性名作为参数 `toRef(<对象名>, '<对象中的属性名>')`
    ```typescript
    let nl = toRef(person, 'age') 
    ```

### computed计算属性
- 基本语法 `let <计算属性名> = computed(() => { ... return ... })`
- 计算属性有缓存 多次相同值调用只计算一次
#### 定义一个计算属性 只读
```typescript
let fullName = computed(() => {
    console.log('computed');
    return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
})
```
#### 定义一个计算属性 可读可写
```typescript
// 计算属性
let fullName = computed({
    get() {
        return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
    },
    set(val) {
        let [str1, str2] = val.split('-');
        firstName.value = str1;
        lastName.value = str2;
    }
})

// 由点击事件触发
let changeFullName = () => {
    fullName.value = 'li-si'
}
```

### watch
- **作用：** 监视数据的变化 (和`Vue2`中的*作用*是一致的)
- **特点：** `Vue3`中的`watch`只能监视以下**四种数据**：
    > 1. `ref`定义的数据
    > 2. `reactive`定义的数据
    > 3. 函数返回一个值(一个`getter()`函数)
    > 4. 一个包含上述内容的数组
#### 语法
- **有三个参数：** `watch(<被监视的数据>, (newVal, oldVal) => {<回调函数>}), {<配置对象>}`
#### 一般遇到的几个情况
##### 情况1
- 监视`ref`定义的**基本类型**数据 直接写数据名即可 不需要`.value`
  ```typescript
  // 数据
  let sum = ref(0)
  // 监视
  const stopWatch = watch(sum, (newValue, oldValue) => {
    console.log('sum changed', newValue, oldValue)
    if (newValue >= 10)
      stopWatch()   // 解除监视
  })
  ```
##### 情况2
- 监视`ref`定义的**对象类型**数据 监视的是对象的**地址值** 
- 若想监视其内部数据 要手动开启深度监视
> 注意
> - 若修改的`ref`定义的对象的属性 `newValue`和`oldValue`都是新的值
> - 若修改整个`ref`对象 `newValue`是新值 `oldValue`是旧值 因为不是同一个对象了
```typescript
// 数据
let person = ref({name: '张三', age: 18});
let changeName = () => person.value.name += '~';
let changeAge = () => person.value.age += 1;
let changePerson = () => person.value = {name: '李四', age: 29};
// 监视
watch(person, (newValue, oldValue) => {
  console.log('person变化了', newValue, oldValue)  // 默认监视的是对象的地址值 需要开启深度监视
}, {
  deep: true, // 需要开启深度监视
  // immediate: true, // 立即执行
})
```
##### 情况3
- 监视`reactive`定义的**对象类型**数据
- 默认开启了深度监视 而且不可关闭
```typescript
// 数据
let obj = reactive({a: {b: {c: 666,}}})
// 方法
let changeVal = () => obj.a.b.c = 999
// 监视
watch(obj, (newVal, oldVal) => {
  console.log('obj变化了', newVal, oldVal)
})
```