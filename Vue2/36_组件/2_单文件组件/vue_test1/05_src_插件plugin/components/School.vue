<script>
export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
    }
  },
  methods: {
    test() {
      this.hello();
    }
  }
}
</script>

<template>
  <div class="school">
    <h2>学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
    <button @click="test">测试hello方法</button>
  </div>
</template>

<style>
.school {
  background-color: white;
}
</style>