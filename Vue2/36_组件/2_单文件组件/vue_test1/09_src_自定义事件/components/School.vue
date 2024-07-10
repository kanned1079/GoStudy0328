<script>
export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
    }
  },
  props: ['getSchoolName'],
  methods: {
    sendSchoolName() {
      this.getSchoolName(this.name);
    }
  }
}
</script>

<template>
  <div class="school">
    <h2>学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
    <button @click="sendSchoolName">sendSchoolName</button>
  </div>
</template>

<style scoped>
.school {
  background-color: skyblue;
  padding: 5px;
  margin: 5px;
}
</style>