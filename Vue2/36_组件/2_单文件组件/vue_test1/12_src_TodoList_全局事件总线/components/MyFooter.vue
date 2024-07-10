<script>
export default {
  name: 'MyFooter',
  // props: ['todos', 'deleteFinished', 'checkAllTodo'],
  props: ['todos'],
  computed: {
    getTotal() {
      return this.todos.length;
    },
    doneTotal() {
      // 方法1
      // let i = 0;
      // this.todos.forEach((todo) => {
      //   if (todo.done)
      //     i++;
      // })
      // return i
      // 方法2
      // 全部
      // let x = this.todos.reduce((pre, current) => {
      //   console.log('pre =', pre);
      //   console.log('current =', current)
      //   return pre + (current.done ? 1 : 0)
      // }, 0)

      // 精简
      return this.todos.reduce((pre, current) => pre + (current.done ? 1 : 0), 0)

      // console.log('x =', x)
      // return x

    },
    // isAll() {
    //   return this.doneTotal === this.getTotal && this.getTotal > 0;
    // },
    isAll: {
      get() {
        return this.doneTotal === this.getTotal && this.getTotal > 0;
      },
      set(value) {
        this.$emit('checkAllTodo', value)
      }
    }
  },
  methods: {
    checkAll(event) {
      // this.checkAllTodo(event.target.checked)
      this.$emit('checkAllTodo', event)
    },
    clearAll() {
      this.$emit('deleteFinished')
    }
  }
}
</script>

<template>
  <div class="todo-footer" v-show="getTotal">
    <label>
      <input type="checkbox" :checked="isAll" @change="checkAll"/>
      <!--      <input type="checkbox" v-model="isAll"/>-->
    </label>
    <span>
      <span>已完成{{ doneTotal }}</span> / 全部{{ getTotal }}
    </span>
    <button class="btn btn-danger" @click="clearAll">清除已完成任务</button>
  </div>
</template>

<style scoped>
/*footer*/
.todo-footer {
  height: 40px;
  line-height: 40px;
  padding-left: 6px;
  margin-top: 5px;
}

.todo-footer label {
  display: inline-block;
  margin-right: 20px;
  cursor: pointer;
}

.todo-footer label input {
  position: relative;
  top: -1px;
  vertical-align: middle;
  margin-right: 5px;
}

.todo-footer button {
  float: right;
  margin-top: 5px;
}
</style>