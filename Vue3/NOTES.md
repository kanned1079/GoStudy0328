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

#### 选项式API的弊端

`Options`类型的API 数据、方法、计算属性等 是分散在`data` `methods` `computed`中的 如果想新增或者修改一个需求 就需要分别修改
不利于维护

```html

<script lang="ts">
    export default {
        name: 'Person',
        data() {
            return {...}
        },
        methods: {...},
        computed: {...},
        watch: {...}
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

##### 情况4

- 监视`ref`或`reactive`定义的**对象类型**数据中的**某个属性**有以下注意点
    1. 若该属性**不是** *对象类型* 需要写成函数形式
    2. 若该属性值**依然**是 *对象类型* 可直接编写 也可以写成函数 建议使用函数
- *ps:* 监视的要是对象中的属性 那么最好写函数式
- *注意点：* 若是对象监视的式地址值 需要关注对象内部 需要手动开启深度监视

  ```typescript
  // 只监视对象中的一个属性 该属性是基本类型 需要写成函数式
  watch(() => person.name, (newValue, oldValue) => {
      console.log('person.name变化了', newValue, oldValue)
  })
  // 使用函数式 建议的写法 此时监视的就是对象的地址值了
  watch(() => person.cars, (newValue, oldValue) => {
      console.log('person.cars变化了 函数式', newValue, oldValue)
  }, { deep: true, })
  ```  

##### 情况5

- 监视指定的多个属性
  ```typescript
  watch([() => person.name, () => person.cars.c1], (newValue, oldValue) => {
    console.log('person.name或person.cars.c1变化了', newValue, oldValue)
  }, {deep: true,})
  ```

### watchEffect

- 立即运行一个函数 同时响应式地之宗其以来 并在依赖更改时重新执行该函数
- `watch`对比`watchEffect`:
    - 都能监听响应式数据的变化 不同的是监听数据变化的方式不同
    - `watch` 要明确指出监视的数据
    - `watchEffect` 不用明确指出监视的数据 (函数中用到那些属性 那就监视哪些属性)

  ```typescript
  // 使用watch实现
  watch([temp, height], (newValue) => {
    console.log(newValue)
    let [newTemp, newHeight] = newValue;
    console.log(newTemp, newHeight)
    if (newTemp >= 60 || newHeight >= 80)
      console.log('sent req')
  })
  // 使用watchEffect实现
  watchEffect(() => {
    console.log('@'); // 会立即运行
    if (temp.value >= 60 || height.value >= 80)
      console.log('sent req')
  })
  ```

### 标签的`ref`标识

- **作用：** 用于注册模版引用
    - 用在普通`DOM`标签上 获取的是`DOM`节点
    - 用在组件标签上 获取的是组件的实例对象

  ```typescript
  let title2 = ref() // 存储ref标记的内容
  // 每个组件间的title2相互独立
  ```

  ```html
  <!--这是用在DOM标签上-->
  <h2 ref="title2">上海</h2>
  <!--这是用在组件实例上-->
  <Person ref="ren"></Person>
  ```

    - 父组件可以获取子组件导出的数据 使用`defineExpose`

  ```typescript
  // 引用defineExpose
  import {ref, defineExpose} from 'vue'
  // ...
  // 在子组件最后进行导出
  defineExpose({a, b, c})
  ```

### TS类型限制`interface`

#### 定义`@/type/index.ts`

```typescript
// 定义一个接口 用于限制Person对象的具体属性
export interface PersonInter {
    id: string
    name: string
    age: number
}

// 一个自定义类型
export type Persons = Array<PersonInter>
```

#### 使用`@/components/Person.vue`

```typescript
// 导入
import {type PersonInter, type Persons} from '@/types'
// 使用
// 单个属性
let person: PersonInter = {id: '0001', name: '张三', age: 60}
// 数组类型
let personList: Array<PersonInter> = [
    // 必须要符合接口规范
    {id: '0001', name: '张三', age: 60},
]
// 另一种写法
let personsList: Persons = [
    {id: '0001', name: '张三', age: 60},
]
```

### 父传递数据 `defineProps`和`withDefaults`的使用
- **作用：** 使用`defineProps`来定义父组件传递的`props`
#### 父组件中设置数据
```typescript
import {type Persons} from "@/types"; // 数据类型限制
// 推荐写法
let personList = reactive<Persons>([
  {id: '0001', name: '张三', age: 18},
  {id: '0002', name: '李四', age: 19},
  {id: '0003', name: '王五', age: 20, x: 999},
])
```

```html
<Person v-bind:list="personList"></Person>
```
  
#### 在子组件中接收
```typescript
// 导入限制
import { type Persons } from "@/types";
import { withDefaults } from "vue";
// 只接收a
// 但是这样的接收可能会引起类型不匹配问题
// defineProps(['a'])

// 接收a 同时将a保存
// let x = defineProps(['a', 'list'])

// 接收a + 限制类型
// defineProps<{list:Persons}>()

// 接收list 限制类型 限制必要性 指定默认值
withDefaults(defineProps<{ list?: Persons }>(), {
  list: () => { // 默认值需要以函数形式返回
    return [
      { id: '10002', name: 'kanna', age: 21 },
      { id: '10003', name: 'kinggyo', age: 16 },
      // ...
    ]
  }
})
```
