<script>
import {mapState, mapGetters, mapMutations, mapActions} from "vuex";

export default {
  name: 'MyCount',
  data() {
    return {
      // 移动到index.js
      // sum: 0, // 求和
      n: 0, // 选择的值
    }
  },
  methods: {
    // 借助mapMutations生成对应的方法 方法中会调用commit去联系mutations （对象写法）
    ...mapMutations('countAbout', {increment: 'JIA', decrement: 'JIAN'}),
    // 使用mapActions生成对应的方法 方法会调用dispatch去联系action （对象写法）
    ...mapActions('countAbout', {incrementOdd: 'jiaOdd', incrementWait: 'jiaWait'}),
  },
  computed: {
    ...mapState('countAbout', ['sum', 'school', 'subject']),
    ...mapState('personAbout', ['personList']),
    ...mapGetters('countAbout', ['bigSum']),
  },
  mounted() {
    // const x = mapState({})
    // console.log(x);
  }
}
</script>

<template>
  <div>
    <h1>当前求和为：{{ sum }}</h1>
    <h3>放大10倍后的值：{{ bigSum }}</h3>
    <h3>我在{{ school }}, 学习{{ subject }}</h3>
    <h3 style="color: orange">Person组件的总人数是: {{ personList.length }}</h3>

    <select v-model.number="n">
      <option value=1 selected>1</option>
      <option value=2>2</option>
      <option value=3>3</option>
    </select>
    <button @click="increment(n)">+</button>
    <button @click="decrement(n)">-</button>
    <button @click="incrementOdd(n)">当前求和为奇数再加</button>
    <button @click="incrementWait(n)">等一等再加</button>

  </div>
</template>

<style scoped>
button {
  margin: 5px;
}
</style>