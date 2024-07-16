<script>
export default {
  name: 'MyDetail',
  mounted() {
    console.log(this.$route)
  },
  computed: {
    // getId() {
    //   return this.$route.params.id;
    // }
  },
  // props: ['a', 'b'],
  props: ['id', 'title'],
}
</script>

<template>
<ul>
  <li>消息编号 {{ id }}</li>
  <li>消息标题 {{ title }}</li>
<!--  <li>a:{{a}}, b:{{b}}</li>-->
</ul>
</template>

<style scoped>

</style>