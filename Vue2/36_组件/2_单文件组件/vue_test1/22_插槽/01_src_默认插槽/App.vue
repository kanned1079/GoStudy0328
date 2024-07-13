<template>
  <div class="container">
    <Search></Search>
    <MyCategory title="美食">
      <!--      这里面的属于是属性标签的标签体内容-->
      <img :src="this.$imgAddr" alt="">
    </MyCategory>
    <MyCategory title="游戏">
        <ul>
          <li v-for="(item,index) in games" :key="index">{{item}}</li>
        </ul>
    </MyCategory>
    <MyCategory title="电影">
      <video controls :src="this.$videoAddr"></video>
    </MyCategory>
  </div>
</template>

<script>
import MyCategory from "@/components/Catagory.vue";

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
