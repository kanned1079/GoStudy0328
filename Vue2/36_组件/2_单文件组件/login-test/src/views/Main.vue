<script>
import CommonAside from "@/components/CommonAside.vue";
import CommonHeader from "@/components/CommonHeader.vue";
import CommonTag from "@/components/CommonTag.vue";

export default {
  name: '',
  components: {CommonAside, CommonHeader, CommonTag}
}
</script>

<template>
  <div>
    <!--    <h1>Main</h1>-->
    <!--    <router-view></router-view>-->
    <el-container>
      <el-aside width="auto">
        <CommonAside></CommonAside>
      </el-aside>
      <el-container>
        <el-header height="50px">
          <CommonHeader></CommonHeader>
        </el-header>
        <div style="padding: 10px 10px; background-color: #eef1f7">
          <CommonTag></CommonTag>
        </div>
        <el-main style="background-color: #eef1f7">
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.el-header {
  padding: 0;
}
</style>