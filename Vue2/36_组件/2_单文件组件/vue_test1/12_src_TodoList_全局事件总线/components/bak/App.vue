<template>
  <div id="root">
    <div class="todo-container">
      <div class="todo-wrap">
        <MyHeader :addTodo="addTodo"></MyHeader>
        <MyList :todos="todos" :checkTodo="checkTodo"></MyList>
        <MyFooter></MyFooter>
      </div>
    </div>

  </div>
</template>

<script>
import MyHeader from "@/components/MyHeader.vue";
import MyFooter from "@/components/MyFooter.vue";
import MyList from "@/components/MyList.vue";

export default {
  name: 'App',
  components: {
    MyFooter,
    MyHeader,
    MyList
  },
  data() {
    return {
      todos: [
        {id: '001', title: '吃饭', done: true},
        {id: '002', title: '睡觉', done: false},
        {id: '003', title: '开车', done: true},
      ]
    }
  },
  methods: {
    // 添加一个todo
    addTodo(todoObj) {
      console.log('i am app component, i got data', todoObj);
      this.todos.unshift(todoObj);
    },
    // 取消勾选一个todo
    checkTodo(id) {
      this.todos.forEach((todo) => {
        todo.id === id ? todo.done = !todo.done : todo.done
      })
    }

  }
}

</script>

<style>
/*base*/
body {
  background: #fff;
}

.btn {
  display: inline-block;
  padding: 4px 12px;
  margin-bottom: 0;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  vertical-align: middle;
  cursor: pointer;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2),
  0 1px 2px rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

.btn-danger {
  color: #fff;
  background-color: #da4f49;
  border: 1px solid #bd362f;
}

.btn-danger:hover {
  color: #fff;
  background-color: #bd362f;
}

.btn:focus {
  outline: none;
}

.todo-container {
  width: 600px;
  margin: 0 auto;
}

.todo-container .todo-wrap {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}
</style>
