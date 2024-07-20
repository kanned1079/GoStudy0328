<script>
import {mapState} from "vuex";
export default {
  name: 'CommonHeader',
  methods: {
    handleMenu() {
      this.$store.commit('tab/collapseMenu')
    }
  },
  computed: {
    ...mapState('tab',['tabList'])
  },
  mounted() {
    console.log('tablist ', this.tabList)
    // console.log()
  }
}
</script>

<template>
  <div class="header-container">
    <div class="l-content">
      <el-button icon="el-icon-s-grid" size="mini" @click="handleMenu" style="background-color: #eef1f7"></el-button>
      <!--      面包屑区域-->
      <!--      <span class="text">首页</span>-->
      <el-breadcrumb separator="/">
        <el-breadcrumb-item
            v-for="item in tabList"
            :key="item.path"
            :to="{
              path: item.path,
             }"
        > {{ item.label }} </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="r-content">
      <el-dropdown>
  <span class="el-dropdown-link">
<!--    下拉菜单<i class="el-icon-arrow-down el-icon&#45;&#45;right"></i>-->
    <img class="user" src="../assets/images/user1.jpg" alt="">
  </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>个人中心</el-dropdown-item>
          <el-dropdown-item>退出</el-dropdown-item>

        </el-dropdown-menu>
      </el-dropdown>
    </div>

  </div>
</template>

<style lang="less" scoped>
.header-container {
  background-color: #344f8d;
  height: 50px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .text {
    color: white;
    font-size: 14px;
    margin-left: 14px;
  }

  .el-breadcrumb {
    color: white;
    font-size: 14px;
    margin-left: 14px;
  }

  .r-content {
    .user {
      width: 30px;
      height: 30px;
      border-radius: 50px;
    }
  }
  .l-content {
    display: flex;
    align-items: center;
    /deep/.el-breadcrumb__item {
      .el-breadcrumb__inner {
        font-weight: normal;
        &.is-link {
          color: #5a6e9d;
        }
      }
      &:last-child {
        .el-breadcrumb__inner {
          color: white;
        }
      }
    }
  }
}

.header-container {
  padding: 0 20px;
}


</style>