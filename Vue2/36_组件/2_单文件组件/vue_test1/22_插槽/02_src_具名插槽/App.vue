<template>
  <div class="container">
    <Search></Search>
    <MyCategory title="美食">
      <!--      这里面的属于是属性标签的标签体内容-->
      <img :src="this.$imgAddr" alt="" slot="center">
      <a href="http://baidu.com/" slot="footer">更多美食</a>
    </MyCategory>
    <MyCategory title="游戏">
        <ul slot="center">
          <li v-for="(item,index) in games" :key="index">{{item}}</li>
        </ul>
      <div slot="footer">
        <a href="http://baidu.com/">单机游戏</a>
        <!--      注意这里用&nbsp;不行-->
        &nbsp;
        <a href="http://baidu.com/">网络游戏</a>
      </div>


    </MyCategory>
    <MyCategory title="电影">
      <video slot="center" controls :src="this.$videoAddr"></video>
      <template v-slot:footer>
        <div>
          <a href="http://baidu.com/">经典</a>&nbsp;
          <a href="http://baidu.com/">推荐</a>&nbsp;
          <a href="http://baidu.com/">热门</a>
        </div>
        <h4 slot="footer">欢迎来影院观看</h4>
      </template>
    </MyCategory>
  </div>
</template>

<script>
import MyCategory from "@/components/Count.vue";

export default {
  name: 'App',
  components: {MyCategory},
  data() {
    return {
      foods: ['火锅', '烧烤', '小龙虾', '牛排'],
      games: ['红色警戒', '穿越火线', '劲舞团', '超级玛丽'],
      films: ['《教父》', '《拆弹专家》', '你好，李焕英', '《星际穿越》'],
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: space-around;
}
</style>
