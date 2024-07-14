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
    // increment() {
    //   // this.$store.dispatch('jia', this.n);
    //   this.$store.commit('JIA', this.n);
    // },
    // decrement() {
    //   // this.$store.dispatch('jian', this.n);
    //   this.$store.commit('JIAN', this.n);
    // },
    // 借助mapMutations生成对应的方法 方法中会调用commit去联系mutations （对象写法）
    ...mapMutations({increment: 'JIA', decrement: 'JIAN'}),
    // 数组写法
    // ...mapMutations(['JIA', 'JIAN']),

    // ------------- 我是分割线 -----------------

    // incrementOdd() {
    //   // this.$store.state.sum % 2 !== 0 ? this.$store.dispatch('jia', this.n) : this.$store.state.sum;
    //   this.$store.dispatch('jiaOdd', this.n);
    // },
    // incrementWait() {
    //   // setTimeout(() => {
    //   //   this.$store.dispatch('jia', this.n);
    //   // }, 500);
    //   this.$store.dispatch('jiaWait', this.n);
    // },
    // 使用mapActions生成对应的方法 方法会调用dispatch去联系action （对象写法）
    ...mapActions({incrementOdd: 'jiaOdd', incrementWait: 'jiaWait'}),
    // 数组写法
    // ...mapActions(['jiaOdd', 'jiaWait']),
  },
  computed: {

    ...mapState(['sum', 'school', 'subject']),
    ...mapGetters(['bigSum']),

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