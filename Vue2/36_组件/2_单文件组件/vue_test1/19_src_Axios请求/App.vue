<template>
  <div>
    <button @click="getStudents">获取Student信息</button>
    <button @click="getCars">获取汽车信息</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {}
  },
  methods: {
    getStudents() {
      axios.get('http://localhost:8080/prefix1/students').then(
          response => {
            console.log('请求成功， 学生信息：', response.data);
          },
          error => {
            console.log('请求失败', error.message);
          }
      )
    },
    getCars() {
      axios.get('http://localhost:8080/prefix2/cars').then(
          response => {
            console.log('请求成功，汽车信息：', response.data);
          },
          error => {
            console.log('请求失败', error.message);
          }
      );

    }
  }

}
</script>

<style>

</style>
