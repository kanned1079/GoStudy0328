<script>
export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
    }
  },
}
</script>

<template>
  <div class="school">
    <h2 class="title">学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
  </div>
</template>

<style scoped>
.school {
  background-color: skyblue;
}
</style>