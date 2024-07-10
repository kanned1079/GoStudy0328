<script>
import Student from './components/Student.vue'
import School from './components/School.vue'

export default {
  name: 'App',
  components: {
    Student,
    School,
  },
  data() {
    return {
      msg: 'Hello.',
      studentName: ''
    }
  },
  methods: {
    getSchoolName(name) {
      console.log('app 收到的学校名：', name);
    },
    getStudentName(name, ...args) {
      console.log('app 收到的学生名：', name, args);
      this.studentName = name;
    },
    m1() {
      console.log('demo is called.')
    }
  },
  mounted() {
    // this.$refs.student.$on('aaaaa', this.getStudentName)
    // setTimeout(() => {
    //   // this.$refs.student.$on('aaaaa', this.getStudentName)
    //   // 如果要设置只能执行一次
    //   this.$refs.student.$once('aaaaa', this.getStudentName)
    //   this.$refs.student.$once('aaaaa', function () {
    //     console.log('app 收到的学生名：', name, args);
    //     this.studentName = name;
    //   })
    // }, 3000)
    // this.$refs.student.$on('aaaaa', (name, ...args) {
    //   console.log(name);
    //   this.studentName = name;
    // })
    // 这里绑定的是子组件 但是回调函数在父组件中 就可以从子组件中获得数据
    this.$refs.student.$on('aaaaa', this.getStudentName)
    // 如果需要只执行一次那就使用once
  }
}
</script>

<template>
  <div class="app">
    <h1>{{ msg }}</h1>
    <h2>学生姓名 {{ studentName }}</h2>
    <!--    通过父组件给子组件传递函数类型的props实现 子给父传递数据-->
    <School :getSchoolName="getSchoolName"></School>
    <Student @demo="m1" @click="show"></Student>
    <Student ref="student"></Student>
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

.app {
  background-color: gray;
}
</style>