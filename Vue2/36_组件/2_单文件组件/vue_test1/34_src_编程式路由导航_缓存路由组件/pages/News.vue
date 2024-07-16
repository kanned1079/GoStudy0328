<script>
export default {
  name: 'MyNews',
  beforeDestroy() {
    console.log('组件News即将被销毁')
  }
}
</script>

<template>
  <ul>
    <li>news001 <input type="text"></li>
    <li>news002 <input type="text"></li>
    <li>news003 <input type="text"></li>
  </ul>
</template>

<style scoped>

</style>