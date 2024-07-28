// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {reactive, toRefs, toRef} from "vue";
// 数据
let person = reactive({
  name: '张三',
  age: 18
})
// 方法
// let changeName = () => name += '~'
let changeName = () => name.value += '~'
// let changeAge = () => age += 1
let changeAge = () => age.value += 1
// 使用toRefs配合解构函数 将一个reactive解构为多个ref 随后对ref响应即可
let {name, age} = toRefs(person)
console.log(toRefs(person))
// toRef解构对象的单个属性
let nl = toRef(person, 'age')
console.log(nl)
</script>

<template>
  <div class="person">
    <h2>姓名 {{ name }}</h2>
    <h2>年龄 {{ age }}</h2>
    <button @click="changeName">changeName</button>
    <button @click="changeAge">changeAge</button>

  </div>
</template>

<style scoped lang="less">
.person {
  background-color: skyblue;
  border-radius: 10px;
  padding: 20px;

  h1 {
    color: white;
  }

  button {
    margin: 0 5px;
  }
}
</style>