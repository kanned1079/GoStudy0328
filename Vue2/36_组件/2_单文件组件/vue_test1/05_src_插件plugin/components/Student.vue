<script>
export default {
  name: 'Student',
  data() {
    return {
      name: 'kanna',
      sex: 'male',
    }
  },
}
</script>

<template>
  <div class="school">
    <h2>学生名称：{{ name }}</h2>
    <h2>学生性别：{{ sex }}</h2>
    <input type="text" v-fbind:value="name">
  </div>
</template>

<style>
.school {
  background-color: white;
}
</style>