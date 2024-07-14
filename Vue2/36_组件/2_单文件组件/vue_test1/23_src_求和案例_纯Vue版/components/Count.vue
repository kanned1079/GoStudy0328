<script>
export default {
  name: 'MyCount',
  data() {
    return {
      sum: 0, // 求和
      n: 0, // 选择的值
    }
  },
  methods: {
    increment() {
      this.sum += this.n;
    },
    decrement() {
      this.sum -= this.n;
    },
    incrementOdd() {
      this.sum % 2 !== 0 ? this.sum += this.n : this.sum;
    },
    incrementWait() {
      setTimeout(() => {
        this.sum += this.n;
      }, 500);
    }


  }
}
</script>

<template>
  <div>
    <h1>当前求和为：{{ sum }}</h1>
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