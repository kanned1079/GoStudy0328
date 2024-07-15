<template>
  <div>
    <div class="row">
      <div class="col-xs-offset-2 col-xs-8">
        <div class="page-header"><h2>Vue Router Demo</h2></div>
      </div>
    </div>
    <div class="row">
      <div class="col-xs-2 col-xs-offset-2">
        <div class="list-group">
          <!--原始使用a标签实现页面跳转-->
          <!--                    <a class="list-group-item active" href="./about.html">About</a>-->
          <!--                    <a class="list-group-item" href="./home.html">Home</a>-->
          <!--vue中借助router=link标签实现路由的切换-->
          <router-link class="list-group-item" active-class="active" to="/about">About</router-link>
          <router-link class="list-group-item" active-class="active" to="/home">Home</router-link>

        </div>
      </div>
      <div class="col-xs-6">
        <div class="panel">
          <div class="panel-body">
            <!--router-view确定视图的位置-->
            <!--            组件-->
            <router-view></router-view>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import MyHome from './components/Home.vue'
// import MyAbout from "@/components/About.vue";
export default {
  name: 'App',
  // components: {MyAbout, MyHome},
}
</script>

<style scoped>
</style>
