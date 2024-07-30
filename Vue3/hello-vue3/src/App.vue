<script lang="ts" setup name="App">
import Person from '@/components/Person.vue'
import {type Persons} from "@/types";
import {reactive} from "vue";
// 推荐写法
let personList = reactive<Persons>([
  {id: '0001', name: '张三', age: 18},
  {id: '0002', name: '李四', age: 19},
  {id: '0003', name: '王五', age: 20, x: 999},
])
</script>

<template>
<!--  :代表表达式-->
  <Person v-bind:list="personList"></Person>
</template>

<style lang="less">

</style>