// 或者使用一个插件来设置 标签中需要写入name="xxx"
<script lang="ts" setup name="Person">
// 导入限制
import { type Persons } from "@/types";
import { withDefaults } from "vue";
// 只接收a
// 但是这样的接收可能会引起类型不匹配问题
// defineProps(['a'])

// 接收a 同时将a保存
// let x = defineProps(['a', 'list'])

// 接收a + 限制类型
// defineProps<{list:Persons}>()

// 接收list 限制类型 限制必要性 指定默认值
withDefaults(defineProps<{ list?: Persons }>(), {
  list: () => {
    return [
      { id: '10002', name: 'kanna', age: 21 },
      { id: '10003', name: 'kinggyo', age: 16 },
      // ...
    ]
  }
})

</script>


<template>
  <div class="person">
    <!-- <h2>{{ .a }}</h2> -->
    <ul>
      <li v-for="item in list" :key="item.id">
        {{ item.name }} - {{ item.age }}
      </li>
    </ul>
  </div>
</template>

<style scoped lang="less">
.person {
  background-color: skyblue;
  border-radius: 10px;
  padding: 20px;
  margin-bottom: 5px;

  h1 {
    color: white;
  }

  button {
    margin: 0 5px;
  }
}
</style>