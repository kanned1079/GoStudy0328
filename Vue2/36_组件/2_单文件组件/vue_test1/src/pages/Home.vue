<script>
export default {
  name: 'MyHome',
  beforeDestroy() {
    console.log('Home beforeDestroy')
  },
  mounted() {
    console.log('Home mounted')
    console.log(this)
  }
}
</script>

<template>
  <h2>我是Home的内容</h2>
</template>

<style scoped>

</style>