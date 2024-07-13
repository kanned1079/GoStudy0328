<script>
export default {
  name: 'MyCategory',
  // props: ['listData', 'title'],
  props: ['title'],
}
</script>

<template>
<div class="category">
  <h3>{{ title }}</h3>
<!--  定义一个插槽 等待组件使用者填充内容-->
  <slot>我是一些默认值 如果没有传递默认值就会出现</slot>
<!--  <ul v-show="title !== '美食'">-->
<!--    <li v-for="(item,index) in listData" :key="index">{{item}}</li>-->
<!--  </ul>-->
<!--  <img :src="this.$imgAddr" alt="" v-show="title === '美食'">-->
</div>
</template>

<style scoped>
.category {
  background-color: skyblue;
  width: 200px;
  height: 300px;
  border-radius: 5px;
}
h3 {
  text-align: center;
  background-color: orange;
}
img, video {
  width: 100%;
}
</style>