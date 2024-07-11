<script>

export default {
  name: 'Test2',
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
    <transition-group name="hello" appear>
      <!--      如果有name 那么下面就要写成 hello-enter-active-->
      <h1 v-show="isShow" key="1">Hello</h1>
      <h1 v-show="isShow" key="2">Hello2</h1>
    </transition-group>
  </div>
</template>

<style scoped>
h1 {
  background-color: orange;
  transition: 0.5s;
}

/*进入的起点*/
.hello-enter, .hello-leave-to {
  transform: translateX(-100%);
}

/*进入的终点*/
.hello-enter-to, .hello-leave {
  transform: translateX(0);
}


.hello-enter-active {

}

.hello-leave-active {

}


</style>