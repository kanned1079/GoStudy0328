<script lang="ts">
export default {
  name: 'Person',
  beforeCreate() {
      console.log('beforeCreate')
  },
  setup(props, ctx) {
    // 数据
    let name = '张三' // 注意此时的name不是响应式的
    let age = 18
    let tel = '1231234423'

    // 方法
    function modifyName() {
      console.log('modifyName')
      name = 'ZhangSan' // 注意这里的修改 页面也是不会变化的
    }
    function modifyAge() {
      console.log('modifyName')
      age += 1
    }
    function showTel() {
      alert(tel)
    }

    return {name, age, modifyName, modifyAge, showTel}
  },
}
</script>

<template>
  <div class="person">
    <h2>姓名： {{ name }}</h2>
    <h2>年龄： {{ age }}</h2>
    <button @click="modifyName">修改名字</button>
    <button @click="modifyAge">修改年龄</button>
    <button @click="showTel">显示联系方式</button>
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