<script lang="ts" setup name="App">
import Person from '@/components/Person.vue'
import {ref} from "vue";

let showLog = () => {
  console.log(ren.value)  // 这里获得的是组件实例
}
// console.log(title2.value)

let ren = ref()

</script>

<template>
  <h2 ref="title2">Hello</h2>
  <button @click="showLog">显示title2</button>
  <Person ref="ren"></Person>
</template>

<style lang="less">

</style>