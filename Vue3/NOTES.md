# Vue3

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
- Vue2的API设计是`选项式(Options)`风格
- Vue3的API设计是`组合式(Composition)`风格

### OptionsAPI的弊端
`Options`类型的API 数据、方法、计算属性等 是分散在`data` `methods` `computed`中的 如果想新增或者修改一个需求 就需要分别修改 不利于维护
```html
<script lang="ts">
export default {
  name: 'Person',
  data() {
    return {
      name: '张三',
      age: 19,
      tel: '1243211131'
    }
  },
  methods: {
    showTel() {
      alert(this.tel)
    },
    modifyName() {
      this.name = 'zhangsan'
    },
    modifyAge() {
      this.age += 1
    }
  }
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
`setup`是Vue3中一个新的配置项 值是一个函数 它是`Composition API`"表演的舞台" 组件中所用到的 数据 方法 计算属性 监视 均配置在`setup`中
#### 特点如下
- `setup`函数返回对西那个中的内容 可直接在模版中使用
- `setup`中访问`this`是`undefined`
- `setup`函数会在`beforeCreate`中调用 它是领先所有hooks执行的

### ref创建: 基本类型的响应式数据

### reactive创建: 对象类型的响应式数据

### ref对比reactive

### toRefs和toRef

### computed

### watch