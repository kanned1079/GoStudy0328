<script>
export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
    }
  },
  mounted() {
    this.$bus.$on('hello', (data) => {
      console.log('我是School组件', data)
    })
  },
  beforeDestroy() {
    // 谁创建的谁销毁时解绑事件
    this.$bus.$off('')
  }
}
</script>

<template>
  <div class="school">
    <h2>学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
  </div>
</template>

<style scoped>
.school {
  background-color: skyblue;
  padding: 5px;
  margin: 5px;
}
</style>