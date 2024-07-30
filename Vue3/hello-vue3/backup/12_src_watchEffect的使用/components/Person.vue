// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">

import {ref, watch, watchEffect} from "vue";

let temp = ref(10)
let height = ref(0)

let changeTemp = () => temp.value += 1
let changeHeight = () => height.value += 10

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

</script>


<template>
  <div class="person">
    <h3>水温大于等于60 或水位达到80cm时 给服务器发送一个请求</h3>
    <h2>当前水温: {{ temp }}</h2>
    <h2>当前水位: {{ height }}</h2>
    <button @click="changeTemp"> 水温: {{ temp }}oc</button>
    <button @click="changeHeight"> 水位: {{ height }}cm</button>
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