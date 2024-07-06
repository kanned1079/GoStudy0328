# 笔记

## ref属性
1. 被用来给元素或组件注册引用信息（id的替代者）
2. 应用在htm1标签上获取的是真实DOM元素，应用在组件标签上是组件实例对象(vc)
3. 3.使用方式:
   - 打标识: <hl ref"xxx">.....</h1> 或 <School ref "xxx"></School>
   - 获取: this.$refs.xxx

## props配置
1. 传递数据  `<Demo name></Demo>`
2. 接收数据 
   1. 只接收 `props['name']`
   2. 限制类型 `props: {name: Number}`
   3. 限制类型、必要性、指定默认值 `props: { name: {type: String, required: true, default: 'kanna''}}`
3. mark: props是只读的 vue底层会监测你对props的修改 如果进行了修改会触发警告 如果需要修改 需要拷贝一份来修改

## mixin(混入)
- 功能：可以把多个组件共用的配置提取成一个混入对象
- 使用方法：
  1. 先定义mixin (如mixin_1.js) `{data() {...}, methods: {sayHello() {...}, ...}}`
  2. 使用mixin
     1. 全局混入 在 `main.js` 中 `Vue.mixin(xxx)`
     2. 局部混入 `mixin: []` 

## 插件
- 功能：用于哦增强Vue
- 本质：包含install方法的一个对象 install的第一个参数是Vue 第二个以后的参数是插件使用者传递的数据
- 定义插件：`对象.install = function(Vue, options) {Vue.filter(...), Vue.directive('...', {...})}, Vue.mixin(...), Vue.prototype.$myMethod = function() {...}`
- 使用插件：`Vue.use(<插件名>)`

## scoped样式
- 作用：让样式局部生效 防止冲突
- 写法 `<style scoped> ... </style>` 