<script>
export default {
  name: 'MyBanner',
  methods: {
    back() {
      console.log(this.$router)
      this.$router.back();
    },
    forward() {
      this.$router.forward();
    },
    goStep() {
      this.$router.go(-3);
    }
  }
}
</script>

<template>
  <div class="page-header">
    <h2>Vue Router Demo</h2>
    <button @click="back">后退</button>
    <button @click="forward">前进</button>
    <button @click="goStep">测试go</button>
  </div>

</template>

<style scoped>

</style>