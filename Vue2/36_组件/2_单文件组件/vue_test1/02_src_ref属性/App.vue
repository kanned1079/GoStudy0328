<script>
import School from './components/School.vue'

export default {
  name: 'App',
  components: {
    School,
  },
  data() {
    return {
      msg: 'Welcome to Vue/CLI'
    }
  },
  methods: {
    showDOM() {
      console.log(this.$refs.title);  // 获取真实DOM元素
      console.log(this.$refs.sch1); // 获取VueComponent组件实例对象
    }
  }
}
</script>

<template>
  <div>
    <h1 v-text="msg"></h1>
    <button @click="showDOM" ref="title">输出上方Dom</button>
    <!--    给组件大ref标签 获取到VueComponent组件实例对象-->
    <School ref="sch1"></School>
    <School></School>
    <School></School>
  </div>
</template>

<style>
</style>