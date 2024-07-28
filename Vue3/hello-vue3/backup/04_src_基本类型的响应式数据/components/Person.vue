<!--<script lang="ts">-->
<!--export default {-->
<!--  name: 'Person',-->
<!--  setup(props, ctx) {-->
<!--    let a = 99-->
<!--    return {a}-->
<!--  },-->
<!--}-->
<!--</script>-->

// 使用一个单独的script标签来设置name
<!--<script lang="ts">-->
<!--export default {-->
<!--  name: 'Person234'-->
<!--}-->
<!--</script>-->

// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {onMounted, reactive, ref} from 'vue'
let name = ref('张三')
let age = ref(18)
let address = '上海宛平南路600号'
console.log(1, name)
function modifyName() {
  name.value = 'zhangsan'
}
let modifyAge = () => age.value += 1
let showTel = () => alert(address)

</script>

<template>
  <div class="person">
    <h2>姓名： {{ name }}</h2>
    <h2>年龄： {{ age }}</h2>
    <h2>地址： {{ address }}</h2>
    <button @click="modifyName">修改名字</button>
    <button @click="modifyAge">修改年龄</button>
    <button @click="showTel">显示联系方式</button>
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