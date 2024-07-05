<script>
import Student from './components/Student.vue'

export default {
  name: 'App',
  components: {
    Student,
  }
}
</script>

<template>
  <div>
    <Student name="kanna" sex="男"></Student>
    <!--    :age=18 变为了语句-->
<!--    <hr>-->
<!--    <Student sex="女"></Student>-->
<!--    <hr>-->
<!--    <Student></Student>-->
  </div>
</template>

<style>
</style>