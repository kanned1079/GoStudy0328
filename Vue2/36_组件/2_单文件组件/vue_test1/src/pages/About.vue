<script>

export default {
  name: 'MyAbout',
  beforeDestroy() {
    console.log('About beforeDestroy')
  },
  mounted() {
    console.log('About mounted')
  },
  // beforeRouteEnter 通过路由规则进入组件时被调用
  beforeRouteEnter(to, from, next) {
    console.log(to, from);
    next();
  },
  // beforeRouteLeave 通过路由规则离开组件时被调用
  beforeRouteLeave() {

  }
}
</script>

<template>
  <h2>我是About的内容</h2>
</template>

<style scoped>

</style>