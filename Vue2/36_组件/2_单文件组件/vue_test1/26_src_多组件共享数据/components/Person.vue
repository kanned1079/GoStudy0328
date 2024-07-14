<script>
import {mapState} from "vuex";
import {nanoid} from "nanoid";

export default {
  name: 'MyPerson',
  computed: {
    ...mapState(['personList']),
    sum() {
      return this.$store.state.sum;
    }
  },
  methods: {
    add() {
      const personObj = {
        id: nanoid(),
        name: this.name,
      }
      console.log(personObj)
      this.$store.commit('ADD_PERSON', personObj)
    },
  },
  data() {
    return {
      name: '',
    }
  }

}
</script>

<template>
<div>
  <h1>人员列表</h1>
  <h3 style="color: #0c72d2">Count组件的求和为：{{ sum }}</h3>
  <input type="text" placeholder="输入名字" v-model="name">
  <button @click="add">添加</button>
  <ul>
    <li v-for="person in personList" :key="person.id">
      {{ person.id }} - {{ person.name}}
    </li>
  </ul>
</div>
</template>

<style scoped>

</style>