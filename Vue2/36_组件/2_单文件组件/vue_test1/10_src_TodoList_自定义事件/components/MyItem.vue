<script>
export default {
  name: 'MyItem',
  // 声明接收todo对象
  // props: ['todo', 'checkTodo','deleteTodo'],
  props: ['todo'],
  mounted() {
    console.log('MyItem', this.todo)
  },
  methods: {
    handleCheck(id) {
      console.log(id);
      // 通知App组件取反对应值
      // this.checkTodo(id);
      this.$bus.$emit('checkTodo', id); // 触发事件总线上的checkTodo方法
    },
    handleDelete(id) {
      if (confirm('确定要删除吗？')) {
        console.log(id);
        // this.deleteTodo(id);
        this.$bus.$emit('deleteTodo', id); // 触发删除方法
      }
    }
  },
}
</script>

<template>
  <li>
    <label>
      <input type="checkbox" :checked="todo.done" @change="handleCheck(todo.id)">
      <span>{{ todo.title }}</span>
    </label>
    <button class="btn btn-danger" @click="handleDelete(todo.id)">删除</button>
  </li>
</template>

<style scoped>
/*item*/
li {
  list-style: none;
  height: 36px;
  line-height: 36px;
  padding: 0 5px;
  border-bottom: 1px solid #ddd;
}

li label {
  float: left;
  cursor: pointer;
}

li label li input {
  vertical-align: middle;
  margin-right: 6px;
  position: relative;
  top: -1px;
}

li button {
  float: right;
  display: none;
  margin-top: 3px;
}

li:before {
  content: initial;
}

li:last-child {
  border-bottom: none;
}

li:hover {
  background-color: #ddd;
}

li:hover button {
  display: block;
}
</style>