// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {ref, defineExpose} from 'vue'

let title2 = ref() // 存储ref标记的内容
let a = ref(1)
let b = ref(2)
let c = ref(3)

let showLog = () => {
  console.log(title2.value)
}
// 在子组件导出
defineExpose({a, b, c})
</script>


<template>
  <div class="person">
    <h1>北京</h1>
    <h2 ref="title2">上海</h2>
    <h3>南京</h3>
    <button @click="showLog">输出h2元素</button>
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