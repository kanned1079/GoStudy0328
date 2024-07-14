<script>
import {mapState, mapGetters} from "vuex";

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
    increment() {
      // this.$store.dispatch('jia', this.n);
      this.$store.commit('JIA', this.n);
    },
    decrement() {
      // this.$store.dispatch('jian', this.n);
      this.$store.commit('JIAN', this.n);
    },
    incrementOdd() {
      // this.$store.state.sum % 2 !== 0 ? this.$store.dispatch('jia', this.n) : this.$store.state.sum;
      this.$store.dispatch('jiaOdd', this.n);
    },
    incrementWait() {
      // setTimeout(() => {
      //   this.$store.dispatch('jia', this.n);
      // }, 500);
      this.$store.dispatch('jiaWait', this.n);
    },
  },
  computed: {
    // 手动写计算属性
    // sum() {
    //   return this.$store.state.sum;
    // }
    // -----------------------------------
    // 手动写getter
    // bigSum() {
    //   return this.$store.getters.bigSum;
    // },
    // 借助mapState生成计算属性 从state中读取数据 （对象写法）
    // ...mapState({he: 'sum', xuexiao: 'school', xueke: 'subject',}),
    // 使用数组写法
    // 注意使用数组写法必须保持名称和getters中相同
    ...mapState(['sum', 'school', 'subject']),

    // 借助mapGetter生成getter （对象写法）
    // ...mapGetters({bigSum: 'bigSum',})
    // 数组写法
    ...mapGetters(['bigSum']),

  },
  mounted() {
    const x = mapState({})
    console.log(x);
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
    <button @click="increment">+</button>
    <button @click="decrement">-</button>
    <button @click="incrementOdd">当前求和为奇数再加</button>
    <button @click="incrementWait">等一等再加</button>

  </div>
</template>

<style scoped>
button {
  margin: 5px;
}
</style>