<script>
import {mixin_showName} from "@/mixin";

export default {
  name: 'Student',
  data() {
    return {
      name: 'kanna',
      sex: 'male',
    }
  },
  mixins: [mixin_showName],

}
</script>

<template>
  <div class="school">
    <h2 @click="showName">学生名称：{{ name }}</h2>
    <h2>学生性别：{{ sex }}</h2>
  </div>
</template>

<style>
.school {
  background-color: white;
}
</style>