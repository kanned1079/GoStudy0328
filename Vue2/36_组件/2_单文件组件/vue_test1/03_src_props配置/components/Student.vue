<script>
export default {
  name: 'Student',
  data() {
    return {
      msg: 'I am a student',
      myName: this.name,
      mySex: this.sex,
      myAge: this.age,
    }
  },
  // 接收到的prop是不能修改的
  // props: ['name', 'age', 'sex'],  // 简单声明接收
  // 接收同时对数据进行类型限制
  // props: {
  //   name: String,
  //   sex: String,
  //   age: Number,
  // },
  // 接收同时对数据进行类型限制+必要性限制
  props: {
    name: {
      type: String, // 类型为字符串
      required: true  // name是必须要传递的
    },
    sex: {
      type: String,
      required: true
    },
    age: {
      type: Number,
      default: 99   // 默认值
    },
  },
  methods: {
    changeAge() {
      console.log('changeAge')
      this.myAge++;
    }
  }
}
</script>

<template>
  <div class="school">
    <h1>{{ msg }}</h1>
    <h2>学生名称：{{ myName }}</h2>
    <h2>学生性别：{{ mySex }}</h2>
    <h2>学生年龄：{{ myAge + 1 }}</h2>
    <button @click="changeAge">尝试修改Age</button>
  </div>
</template>

<style>
.school {
  background-color: gray;
}
</style>