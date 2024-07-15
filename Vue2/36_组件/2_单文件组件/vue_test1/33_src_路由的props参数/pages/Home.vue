<script>
export default {
  name: 'MyHome',
  // beforeDestroy() {
  //   console.log('Home beforeDestroy')
  // },
  // mounted() {
  //   console.log('Home mounted')
  //   console.log(this)
  // }
}
</script>

<template>
  <div>
    <h2>我是Home的内容</h2>
    <div>
      <ul class="nav nav-tabs">
        <li>
          <router-link class="list-group-item" active-class="active" to="/home/news">News</router-link>
          <!--                    <router-link class="list-group-item" active-class="active" to="/home/news">News</router-link>-->
        </li>
        <li>
          <router-link class="list-group-item" active-class="active" to="/home/message">Message</router-link>
          <!--          <router-link class="list-group-item" active-class="active" to="/home/message">Message</router-link>-->
        </li>
      </ul>
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>

</style>