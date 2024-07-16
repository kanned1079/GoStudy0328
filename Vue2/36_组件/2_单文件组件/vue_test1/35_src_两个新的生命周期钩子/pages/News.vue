<script>
export default {
  name: 'MyNews',
  data() {
    return {
      a: false,
      opacity: 1,
    }
  },
  // mounted() {
  //   console.log('mounted', this);
  //   setInterval(() => {
  //     console.log('@');
  //     this.opacity -= 0.01;
  //     if (this.opacity <= 0)
  //       this.opacity = 1;
  //   }, 16)
  // },
  // beforeDestroy() {
  //   console.log('组件News即将被销毁');
  //   clearInterval(this.timer);
  // },
  activated() {
    console.log('News activated')
    setInterval(() => {
              console.log('@');
              this.opacity -= 0.01;
              if (this.opacity <= 0)
                this.opacity = 1;
            }, 16)
  },
  deactivated() {
    console.log('News deactivated');
    clearInterval(this.timer);
  }
}
</script>

<template>
  <ul>
    <li :style="{opacity}">欢迎学习vue</li>
    <li>news001 <input type="text"></li>
    <li>news002 <input type="text"></li>
    <li>news003 <input type="text"></li>
  </ul>
</template>

<style scoped>

</style>