<script>
export default {
  name: 'MyMessage',
  data() {
    return {
      messageList: [
        {id: '001', title: '消息001'},
        {id: '002', title: '消息002'},
        {id: '003', title: '消息003'},
      ],
    }
  }
}
</script>

<template>
  <div>
    <ul>
      <li v-for="m in messageList" :key="m.id">
        <!--        跳转路由并携带params参数 to的字符串写法-->
<!--                <router-link :to=" `/home/message/detail/${m.id}/${m.title}` ">{{ m.title }}</router-link>-->
        <!--        to的对象写法-->
        <router-link :to="{
          // path: '/home/message/detail',
          // 注意如果传递params 那么这里不能使用path 只能用name
          name: 'xiangqing',
          params: {
            id: m.id,
            title: m.title,
          }
        }">
          {{ m.title }}
        </router-link>
      </li>
    </ul>
    <hr>
    <router-view></router-view>
  </div>
</template>

<style scoped>

</style>