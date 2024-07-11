<script>

export default {
  name: 'Test',
  data() {
    return {
      isShow: true,
    }
  }
}
</script>

<template>
  <div>
    <button @click="isShow = !isShow">显示/隐藏</button>
    <transition name="hello" appear>
      <!--      如果有name 那么下面就要写成 hello-enter-active-->
      <h1 v-show="isShow">Hello</h1>
    </transition>
  </div>
</template>

<style scoped>
h1 {
  background-color: orange;
}

.hello-enter-active {
  animation: ani1 1s;
}

.hello-leave-active {
  animation: ani1 1s reverse;
}

@keyframes ani1 {
  from {
    transform: translateX(-100%);
  }
  to {
    transform: translateX(0px);
  }
}
</style>