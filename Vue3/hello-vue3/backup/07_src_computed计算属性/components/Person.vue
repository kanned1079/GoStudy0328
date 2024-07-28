// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
import { ref ,computed } from 'vue';
let firstName = ref('zhang');
let lastName = ref('san');
//计算属性 有缓存
// 这样定义的计算属性它是只读的不可修改
// let fullName = computed(() => {
//   console.log('computed');
//   return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
// })
// 定义可以读写的计算属性
let fullName = computed({
  get() {
    console.log('computed.get');
    return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
  },
  set(val) {
    console.log('computed.set', val);
    let [str1, str2] = val.split('-');
    firstName.value = str1;
    lastName.value = str2;
  }
})

// 方法 注意方法没有缓存
let fullName2 = () => {
  console.log('function');
  return firstName.value.slice(0, 1).toUpperCase() + firstName.value.slice(1) + '-' + lastName.value;
}
// 点击事件
let changeFullName = () => {
  console.log(fullName);
  fullName.value = 'li-si'
}
</script>

<template>
  <div class="person">
    姓： <input type="text" v-model="firstName"><br>
    名： <input type="text" v-model="lastName"><br>
    <button @click="changeFullName">将全名改为li-si</button>

    全名：<span>{{ fullName }}</span>
    全名：<span>{{ fullName }}</span>
    全名：<span>{{ fullName }}</span>
    <hr>
    全名：<span>{{ fullName2() }}</span>
    全名：<span>{{ fullName2() }}</span>
    全名：<span>{{ fullName2() }}</span>
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