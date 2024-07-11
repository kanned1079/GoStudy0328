<script>
import 'animate.css';
export default {
  name: 'Test3',
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
    <transition
        appear
        name="animate__animated animate__bounce"
        enter-active-class="animate__swing"
        leave-active-class="animate__backOutUp"
    >
      <!--      如果有name 那么下面就要写成 hello-enter-active-->
      <h1 v-show="isShow">Hello3</h1>
    </transition>
  </div>
</template>

<style scoped>
h1 {
  background-color: orange;
}

</style>