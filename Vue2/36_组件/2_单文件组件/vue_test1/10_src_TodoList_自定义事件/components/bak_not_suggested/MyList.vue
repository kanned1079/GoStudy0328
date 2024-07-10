<script>
import MyItem from "@/components/MyItem.vue";

export default {
  components: {
    MyItem
  },
  name: 'MyList',
  props: ['todos'],

}
</script>

<template>
  <ul class="todo-main">
    <MyItem v-for="(todoObj, index) in todos" :key="todoObj.id" :todo="todoObj"></MyItem>

<!--    <li>-->
<!--    <li>-->
<!--      <label>-->
<!--        <input type="checkbox"/>-->
<!--        <span>xxxxxx</span>-->
<!--      </label>-->
<!--      <button class="btn btn-danger" style="display: none">删除</button>-->
<!--    </li>-->
<!--    <li>-->
<!--      <label>-->
<!--        <input type="checkbox"/>-->
<!--        <span>yyyyyy</span>-->
<!--      </label>-->
<!--      <button class="btn btn-danger" style="display: none">删除</button>-->
<!--    </li>-->
  </ul>
</template>

<style scoped>
/*main*/
.todo-main {
  margin-left: 0px;
  border: 1px solid #ddd;
  border-radius: 2px;
  padding: 0px;
}

.todo-empty {
  height: 40px;
  line-height: 40px;
  border: 1px solid #ddd;
  border-radius: 2px;
  padding-left: 5px;
  margin-top: 10px;
}
</style>