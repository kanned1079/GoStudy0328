<script>
export default {
  name: 'MyAbout',
  beforeDestroy() {
    console.log('About beforeDestroy')
  },
  mounted() {
    console.log('About mounted')
  }
}
</script>

<template>
  <h2>我是About的内容</h2>
</template>

<style scoped>

</style>