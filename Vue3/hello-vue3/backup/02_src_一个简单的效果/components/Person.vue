<script lang="ts">
export default {
  name: 'Person',
  data() {
    return {
      name: '张三',
      age: 19,
      tel: '1243211131'
    }
  },
  methods: {
    showTel() {
      alert(this.tel)
    },
    modifyName() {
      this.name = 'zhangsan'
    },
    modifyAge() {
      this.age += 1
    }
  }
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