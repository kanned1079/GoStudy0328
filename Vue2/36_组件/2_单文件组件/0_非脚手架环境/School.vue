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
    name: 'School',
    data() {
      return {
        name: 'NNU',
        addr: 'Nanjing',
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
    <h2>学校名称 {{ name }}</h2>
    <h2>学校地址 {{ addr }}</h2>
    <button @click="showName" class="btn">显示学校名称</button>
  </div>
</template>

<style>
/* 组件的样式 */
.demo {
  background-color: #ffe57d;
}
.btn {
  width: 200px;
  height: 40px;
}
</style>