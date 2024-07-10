<script>
export default {
  name: 'Student',
  data() {
    return {
      name: 'kanna',
      sex: 'male',
      number: 0,
    }
  },
  methods: {
    sendStudentName() {
      // 触发student组件上的aaaaa事件
      this.$emit('aaaaa', this.name, 66, 777, 888);
      // this.$off('demo');
      this.$emit('click')
    },
    unbind() {
      // 解绑一个自定义事件
      this.$off('aaaaa');
      // 解绑多个事件
      // this.$off(['aaaaa', 'demo'])
      // 解绑全部自定义shijian
      this.$off();
    },
    death() {
      // 销毁后所有实例的自定义事件都不可用了
      this.$destroy();
    },
    add() {
      console.log('add回调');
      this.number++;
    },
    mounted() {
      this.$on('haha', (data) => {
        console.log('hhhhhhhhhhh', data)
      })
    }
  }
}
</script>

<template>
  <div class="school">
    <h2 class="qwe">学生名称：{{ name }}</h2>
    <h2>学生性别：{{ sex }}</h2>
    <h2>当前求和为 {{ number }}</h2>
    <button @click="add">number++</button>
    <button @click="sendStudentName">sendStudentName</button>
    <button @click="unbind">解绑aaaa事件</button>
    <button @click="death">销毁当前的Student实例对象vc</button>
  </div>
</template>

<style scoped>
.school {
  background-color: orange;
  padding: 5px;
  margin: 5px;
}
</style>