<script>
import pubsub from 'pubsub-js'

export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
    }
  },
  methods: {
    demo(msgName, data) {
      console.log(this)
      // 第一个参数为消息名 第二个才是消息内容
      console.log('hello消息的回调成功')
      console.log('消息名：', msgName, 'data: ', data);
    }
  },
  mounted() {
    // this.$bus.$on('hello', (data) => {
    //   console.log('我是School组件', data)
    // })

    // 一挂载上就开启订阅消息
    // 返回有一个消息订阅id 需要将其附加到此vc上
    // this.pubId = pubsub.subscribe('hello', (msgName, data) => {
    // 不能使用普通函数需要使用箭头函数否则此处的this将undefined
    //   console.log(this)
    //   // 第一个参数为消息名 第二个才是消息内容
    //   console.log('hello消息的回调成功')
    //   console.log('消息名：', msgName, 'data: ', data);
    // })

    // 或将其封装到函数中
    this.pubId = pubsub.subscribe('hello', this.demo)
  },
  beforeDestroy() {
    // 谁创建的谁销毁时解绑事件
    // this.$bus.$off('')

    // 同样的消息事件也需要解绑
    pubsub.unsubscribe(this.pubId)  // 注意这里面需要的参数不是消息名称而是消息Is
  }
}
</script>

<template>
  <div class="school">
    <h2>学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
  </div>
</template>

<style scoped>
.school {
  background-color: skyblue;
  padding: 5px;
  margin: 5px;
}
</style>