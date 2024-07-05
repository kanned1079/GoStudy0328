<script>
// 导入混入(mixin)
import {mixin_showName} from '../mixin'
export default {
  name: 'School',
  data() {
    return {
      name: 'NNU',
      addr: 'Nanjing',
      x: 600,
      // 如果这里已经有 这个优先级高
    }
  },
  mixins:[mixin_showName],
  mounted() {
    console.log('mounted in School.')
  }
}
</script>

<template>
  <div class="school">
    <h2 @click="showName">学校名称：{{ name }}</h2>
    <h2>地址：{{ addr }}</h2>
  </div>
</template>

<style>
.school {
  background-color: white;
}
</style>