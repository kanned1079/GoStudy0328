// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
// ts语法
import {type PersonInter, type Persons} from '@/types'
// 必须要符合接口规范
let person: PersonInter = {id: '0001', name: '张三', age: 60}

let personList:Array<PersonInter> = [
  {id: '0001', name: '张三', age: 60},
  {id: '0002', name: '李四', age: 61},
  {id: '0003', name: '王五', age: 62},
]

let personsList:Persons = [
  {id: '0001', name: '张三', age: 60},
]

</script>


<template>
  <div class="person">
    ???
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