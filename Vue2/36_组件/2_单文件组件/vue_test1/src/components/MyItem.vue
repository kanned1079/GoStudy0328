<script>
import pubsub from "pubsub-js";

export default {
  name: 'MyItem',
  // 声明接收todo对象
  props: ['todo', 'checkTodo', 'deleteTodo'],
  mounted() {
    console.log(this.todo)
  },
  methods: {
    handleCheck(id) {
      console.log(id);
      // 通知App组件取反对应值
      this.checkTodo(id);
    },
    handleDelete(id) {
      if (confirm('确定要删除吗？')) {
        console.log(id);
        // this.deleteTodo(id);

        // 改为消息订阅
        pubsub.publish('deleteTodoMsg', id);
      }
    },
    // 编辑
    handleEdit(todo) {
      // this.$set(todo, 'isEdit', true);
      // 如果有isEdit就只修改 否则就追加一个
      // 注意不能用 if(todo.isEdit) 因为它可能为false
      todo.hasOwnProperty('isEdit') ? todo.isEdit = true : this.$set(todo, 'isEdit', true);
      // 放在这里会导致语句执行了但是此时input框还没有出来
      this.$refs.inputTitle.focus();
      // 解决方法1 使用setTimeout
      // setTimeout(() => {
      //   this.$refs.inputTitle.focus();
      // }, 0)
      // 官方解决方案 this.$nextTick
      // nextTick指定的回调会在DOM解析后被执行
      this.$nextTick(function () {
        // 当数据改变后 要基于更新后的新的DOM(此时是新的input框)进行某些操作时 要在nextTick所指定的回调函数中执行
          this.$refs.inputTitle.focus();
      })
    },
    // 失去焦点时才真正实现修改的逻辑
    handleBlur(todo, e) {
      todo.isEdit = false;
      if (!e.target.value.trim())
        return alert('输入不能为空');
      this.$bus.$emit('updateTodo', todo.id, e.target.value.trim());
      console.log(todo.id, e.target.value.trim())


    }
  }
}
</script>

<template>
  <li>
    <label>
      <input type="checkbox" :checked="todo.done" @change="handleCheck(todo.id)">
      <span v-show="!todo.isEdit">{{ todo.title }}</span>
      <input v-show="todo.isEdit"
             type="text"
             :value="todo.title"
             @blur="handleBlur(todo, $event)"
             ref="inputTitle"
      >
    </label>
    <button class="btn btn-danger" @click="handleDelete(todo.id)">删除</button>
    <button v-show="!todo.isEdit" class="btn btn-safe" style="margin-right: 10px" @click="handleEdit(todo)">编辑</button>


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