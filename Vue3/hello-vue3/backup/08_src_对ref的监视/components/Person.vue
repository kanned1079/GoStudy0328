// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {ref, watch} from 'vue'
// 数据
let sum = ref(0)
// 方法
let changeSum = () => sum.value += 1
//监视 这里sum不需要.value
const stopWatch = watch(sum, (newValue, oldValue) => {
  console.log('sum changed', newValue, oldValue)
  if (newValue >= 10)
    stopWatch()
})
// ------------------------------
// 数据
let person = ref({name: '张三', age: 18});
// 方法
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



</script>

<template>
  <div class="person">
    <h1>情况1 监视ref定义的值</h1>
    <h2>当前求和为{{ sum }}</h2>
    <button @click="changeSum">sum+1</button>
  </div>
  <div class="person">
    <h1>情况2 监视ref定义的对象数据</h1>
    <h2>name{{ person.name }}</h2>
    <h2>age{{ person.age }}</h2>
    <button @click="changeName">changeName</button>
    <button @click="changeAge">changeName</button>
    <button @click="changePerson">changePerson</button>
  </div>
</template>

<style scoped lang="less">
.person {
  background-color: skyblue;
  border-radius: 10px;
  padding: 20px;
  margin-bottom: 5px;

  h1 {
    color: white;
  }

  button {
    margin: 0 5px;
  }
}
</style>