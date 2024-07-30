// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import {reactive, watch} from "vue";
// 数据
let person = reactive({
  name: '张三',
  age: 18,
  cars: {
    c1: '宝马',
    c2: '奔驰',
  },
})
// 方法
let changeName = () => person.name += '~'
let changeAge = () => person.age += 1
let changeCar1 = () => person.cars.c1 = '奥迪'
let changeCar2 = () => person.cars.c2 = '大众'
// let changeCar = () => person.cars = {c1: '雅迪', c2: '爱玛'}
let changeCar = () => Object.assign(person.cars, {c1: '雅迪', c2: '爱玛'})
//监视
watch([() => person.name, () => person.cars.c1], (newValue, oldValue) => {
  console.log('person.name或person.cars.c1变化了', newValue, oldValue)
}, {deep: true,})

</script>


<template>
  <div class="person">
    <h1>监视多个属性</h1>
    <h2>Name {{ person.name }}</h2>
    <h2>Age {{ person.age }}</h2>
    <h2>Car {{ person.cars.c1 }} - {{ person.cars.c2 }}</h2>
    <button @click="changeName">修改name</button>
    <button @click="changeAge">修改age</button>
    <button @click="changeCar1">修改car1</button>
    <button @click="changeCar2">修改car2</button>
    <button @click="changeCar">修改整个car</button>
  </div>
</template>

<style scoped lang="less">
.person {
  background-color: skyblue;
  border-radius: 10px;
  padding: 20px;
  margin-bottom: 5px;

  h1 {
    color: white;
  }

  button {
    margin: 0 5px;
  }
}
</style>