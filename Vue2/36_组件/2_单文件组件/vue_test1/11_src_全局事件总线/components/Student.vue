<script>
export default {
  name: 'Student',
  data() {
    return {
      name: 'kanna',
      sex: 'male',
      number: 0,
    }
  },
  methods: {
    sendStudentName() {
       this.$bus.$emit('hello', 666)
    }
  }
}
</script>

<template>
  <div class="school">
    <h2 class="qwe">学生名称：{{ name }}</h2>
    <h2>学生性别：{{ sex }}</h2>
    <button @click="sendStudentName">把学生名给School组件</button>
  </div>
</template>

<style scoped>
.school {
  background-color: orange;
  padding: 5px;
  margin: 5px;
}
</style>