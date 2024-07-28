<script lang="ts">
import Person from '@/components/Person.vue'
export default {
  name: 'App',
  components: {Person}
}
</script>

<template>
<div class="app">
  <h1>Hello</h1>
  <Person></Person>
</div>
</template>

<style lang="less">
.app {
  background-color: lightgray;
  border-radius: 10px;
  padding: 20px;
}
</style>