<script>
import Student from './components/Student.vue'
import School from './components/School.vue'

export default {
  name: 'App',
  components: {
    Student,
    School,
  }
}
</script>

<template>
  <div>
    <School></School>
    <Student></Student>
  </div>
</template>

<!--<style lang="css">-->
<!--.title {-->
<!--  color: #42b983;-->
<!--}-->
<!--</style>-->

<style lang="less">
.demo {
  background-color: pink;
  .qwe {
    font-size: 40px;
  }
}
</style>