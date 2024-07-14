<script>
export default {
  name: 'MyCategory',
  // props: ['listData', 'title'],
  props: ['title'],
  data() {
    return {
      games: ['红色警戒', '穿越火线', '劲舞团', '超级玛丽'],
      msg: 'HELLO',
    }
  }
}
</script>

<template>
<div class="category">
  <h3>{{ title }}</h3>
  <slot :games="games" :msg="msg">我是一些默认值 如果没有传递默认值就会出现1</slot>
</div>
</template>

<style scoped>
.category {
  background-color: skyblue;
  width: 200px;
  height: 300px;
  border-radius: 5px;
}
h3 {
  text-align: center;
  background-color: orange;
}
img, video {
  width: 100%;
}
</style>