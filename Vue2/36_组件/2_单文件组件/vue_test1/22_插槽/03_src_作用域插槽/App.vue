<template>
  <div class="container">

    <MyCategory title="游戏">
<!--      作用域插槽必须要使用template标签和scoped-->
       <template slot-scope="aaa">
         <ul>
           <li v-for="(item,index) in aaa.games" :key="index">{{item}}</li>
         </ul>
       </template>
    </MyCategory>

    <MyCategory title="游戏">
      <template slot-scope="aaa">
        <ol>
          <li v-for="(item,index) in aaa.games" :key="index">{{item}}</li>
        </ol>
      </template>>
    </MyCategory>

    <MyCategory title="游戏">
      <template slot-scope="{games}">
        <h4 v-for="(item,index) in games" :key="index">{{item}}</h4>
<!--        <h5>{{ aaa.msg }}</h5>-->
      </template>>
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
      // games: ['红色警戒', '穿越火线', '劲舞团', '超级玛丽'],
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
