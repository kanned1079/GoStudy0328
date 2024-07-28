// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {reactive, ref} from "vue";
// 数据
// let car = {brand: '奔驰', price: 100,}
let car = reactive({brand: '奔驰', price: 100,})
let sum = ref(0)
// let games = ref([
//   {id:'asoid001', name: 'game1'},
//   {id:'asoid002', name: 'game2'},
//   {id:'asoid003', name: 'game3'},
// ])
// let obj = ref({
//   a: {
//     b: {
//       c: 6656
//     }
//   }
// })
// 方法
let changePrice = () => {
  car.price += 10;
  console.log('car.price', car.price)
}
let changeBrand = () => car.brand = '宝马'
// let changeCar = () => car = {brand: '奥拓', price: 12}  // 这就叫重新分配一个对象 会因此失去响应式
let changeCar = () => Object.assign(car, {brand: '奥拓', price: 12})  // 可以使用Object.assign
// let changeFirstGame = () => {
//   // games[0].name = 'aaaaa'
//   games.value[0].name = 'aaaaa'
// }
// let changeObjc = () => {
//   // 如果使用的是ref 那么就需要.value
//   obj.value.a.b.c = 999
// }
let changeSum = () => {
  sum.value += 1
}
</script>

<template>
  <div class="person">
    <h2>汽车信息： 一辆{{car.brand}}汽车，价值{{car.price}}w</h2>
    <button @click="changePrice">修改汽车价格</button>
    <button @click="changeBrand">修改汽车品牌</button>
    <button @click="changeCar">修改汽车</button>
<!--    <h2>Game list</h2>-->
<!--    <ul>-->
<!--      <li v-for="item in games" :key="item.id">-->
<!--        {{ item.id }} -&#45;&#45; {{ item.name }}-->
<!--      </li>-->
<!--    </ul>-->
<!--    <button @click="changeFirstGame">修改第一个game的名字</button>-->
<!--    <h2>测试 {{obj.a.b.c}}</h2>-->
<!--    <button @click="changeObjc">修改第一个game的名字</button>-->
    <hr>
    <h2>当前求和: {{ sum }}</h2>
    <button @click="changeSum">sum++</button>


  </div>
</template>

<style scoped lang="less">
.person {
  background-color: skyblue;
  border-radius: 10px;
  padding: 20px;

  h1 {
    color: white;
  }

  button {
    margin: 0 5px;
  }
}
</style>