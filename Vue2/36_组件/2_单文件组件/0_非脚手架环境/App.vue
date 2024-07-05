<script>
  // 引入组件
  import School from "../../0_非脚手架环境/Student.vue";
  import Student from "../../0_非脚手架环境/Student.vue";

export default {
  name: 'App',
  components: {
    School,
    Student,
  }

}
</script>

<template>
  <div>
    <School></School>
    <Student></Student>
  </div>
</template>
