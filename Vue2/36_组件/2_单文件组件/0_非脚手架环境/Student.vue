<script>
// 组件交互相关的代码
//   const school = Vue.extend({
//     data() {
//       return {
//         name: 'NNU',
//         addr: 'Nanjing',
//       }
//     }
//   });
//   export default school

  export default ({
    name: 'Student',
    data() {
      return {
        name: 'kanna',
        age: 18,
      }
    },
    methods: {
      showName() {
        alert(this.name)
      }
    }
  });


</script>

<template>
  <!--  组件的结构-->
  <div class="demo">
    <h2>学生名称 {{ name }}</h2>
    <h2>学生年龄 {{ age }}</h2>
  </div>
</template>

<style>
/* 组件的样式 */
.demo {
  background-color: bisque;
}
</style>