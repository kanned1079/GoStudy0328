<script>
export default {
  name: '',
  data() {
    return {

    }
  },
  props: ['keyWord']
}
</script>

<template>
  <div>
    我是搜索
    <h2>params: {{$route.params.keyWord}}</h2>
    <h2>query: {{$route.query.k}}</h2>
    <h2>{{ keyWord }}</h2>
  </div>
</template>

<style scoped>

</style>