<script>
import axios from "axios";
export default {
  name: 'MySearch',
  data () {
    return {
      keyWord: '',
    }
  },
  methods: {
    searchUsers() {
      // 请求前更新List数据
      this.$bus.$emit('updateListData', {isFirst: false, isLoading: true, errMsg: '', users: []});
      axios.get(`https://api.github.com/search/users?q=${this.keyWord}`).then(
          responses => {
            console.log('请求成功', responses.data.items);
            // this.$bus.$emit('getUsers', responses.data.items);
            this.$bus.$emit('updateListData', {isLoading: false, errMsg: '', users: responses.data.items});
          },
          error => {
            console.error('请求失败', error.message);
            this.$bus.$emit('updateListData', {isLoading: false, errMsg: error.message, users: []});
          }
      )

    }
  }
}
</script>

<template>
  <section class="jumbotron">
    <h3 class="jumbotron-heading">Search Github Users</h3>
    <div>
      <input type="text" placeholder="enter the name you search" v-model="keyWord">&nbsp;
      <button @click="searchUsers">Search</button>
    </div>
  </section>
</template>

<style scoped>

</style>