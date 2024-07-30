// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {reactive, watch} from 'vue'

// 数据
let person = reactive({name: '张三', age: 18});
let obj = reactive({  //隐式创建了深度监视 且该深度监视无法关闭
  a: {
    b: {
      c: 666,
    }
  }
})
// 方法
let changeName = () => person.name += '~';
let changeAge = () => person.age += 1;
// let changePerson = () => person = {name: '李四', age: 29};
let changePerson = () => Object.assign(person, {name: '李四', age: 29})
let changeVal = () => {obj.a.b.c = 999}
// 监视
watch(person, (newValue, oldValue) => {
  console.log('person变化了', newValue, oldValue)  // 默认监视的是对象的地址值 需要开启深度监视
}, {
  deep: true, // 需要开启深度监视
  // immediate: true, // 立即执行
})
watch(obj, (newVal, oldVal) => {
  console.log('obj变化了', newVal, oldVal)
})

</script>

<template>
<!--  <div class="person">-->
<!--    <h1>情况1 监视ref定义的值</h1>-->
<!--    <h2>当前求和为{{ sum }}</h2>-->
<!--    <button @click="changeSum">sum+1</button>-->
<!--  </div>-->
  <div class="person">
    <h1>情况2 监视reactive定义的对象数据</h1>
    <h2>name{{ person.name }}</h2>
    <h2>age{{ person.age }}</h2>
    <button @click="changeName">changeName</button>
    <button @click="changeAge">changeName</button>
    <button @click="changePerson">changePerson</button>
  </div>
  <div>
    <h1>测试</h1>
    <h2>{{ obj.a.b.c }}</h2>
    <button @click="changeVal">修改</button>
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