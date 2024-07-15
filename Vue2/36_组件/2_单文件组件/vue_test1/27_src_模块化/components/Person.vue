<script>
// import {mapState} from "vuex";
import {nanoid} from "nanoid";

export default {
  name: 'MyPerson',
  computed: {
    // ...mapState('personAbout', ['personList']),
    personList() {
      return this.$store.state.personAbout.personList;
    },
    sum() {
      return this.$store.state.countAbout.sum;
    },
    firstPersonName() {
      return this.$store.getters["personAbout/firstPersonName"]
    }
  },
  methods: {
    add() {
      const personObj = {
        id: nanoid(),
        name: this.name,
      }
      console.log(personObj)
      this.$store.commit('personAbout/ADD_PERSON', personObj);
    },
    addWang() {
      this.$store.dispatch('personAbout/addPersonWang', {id:nanoid(), name: this.name});
    },
    addFromServer() {
      this.$store.dispatch('personAbout/addPersonServer');
    }
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
  <h4>列表中第一个人的名字是 {{ firstPersonName }}</h4>
  <input type="text" placeholder="输入名字" v-model="name">
  <button @click="add">添加</button>
  <button @click="addWang">添加一个姓王的人</button>
  <button @click="addFromServer">从服务器请求</button>
  <ul>
    <li v-for="person in personList" :key="person.id">
      {{ person.id }} - {{ person.name}}
    </li>
  </ul>
</div>
</template>

<style scoped>

</style>