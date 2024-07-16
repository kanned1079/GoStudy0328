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
    console.log('独享路由守卫-进入-About---beforeRouteEnter');
    console.log(to, from, next);
    if (to.meta.isAuth) {
      if (localStorage.getItem('school') === 'NNU') {
        next();
      }
    } else {
      next();
    }
  },
  // beforeRouteLeave 通过路由规则离开组件时被调用
  beforeRouteLeave(to, from, next) {
    console.log('独享路由守卫-离开-About---beforeRouteLeave');
    console.log(to, from, next);
    next();
  }
}
</script>

<template>
  <h2>我是About的内容</h2>
</template>

<style scoped>

</style>