<template>
  <el-menu default-active="1-4-1"
           class="el-menu-vertical-demo"
           @open="handleOpen"
           @close="handleClose"
           :collapse="isCollapse"
           background-color="#fff"
           text-color="#000"
           active-text-color="#2f477f">
    <h3>{{ !isCollapse?'Nxc Networks':'NXC'}}</h3>
    <el-menu-item v-for="item in noChildren" :key="item.name" :index="item.name" @click="clickMenu(item)">
      <i :class="`el-icon-${item.icon}`"></i>
      <span slot="title">{{ item.label }}</span>
    </el-menu-item>

    <el-submenu v-for="item in hasChildren" :key="item.label" :index="item.name">
      <template slot="title">
        <i :class="`el-icon-${item.icon}`"></i>
        <span slot="title">{{ item.label }}</span>
      </template>
      <!--      子菜单-->
      <el-menu-item-group v-for="subItem in item.children" :key="subItem.path">
        <el-menu-item :index="subItem.path" @click="clickMenu(subItem)">{{ subItem.label }}</el-menu-item>
      </el-menu-item-group>
    </el-submenu>

  </el-menu>
</template>

<script>
import {mapMutations} from "vuex";
export default {
  data() {
    return {
      // isCollapse: false,
      menuData: [
        {
          path: '/',
          name: 'home',
          label: '首页',
          icon: 's-home',
          url: 'Home/Home'
        },
        {
          path: '/mall',
          name: 'mall',
          label: '商品管理',
          icon: 'video-play',
          url: 'MallManage/MallManage'
        },
        {
          path: '/user',
          name: 'user',
          label: '用户管理',
          icon: 'user',
          url: 'UserManage/UserManage'
        },
        {
          label: '其他',
          icon: 'location',
          children: [
            {
              path: '/page1',
              name: 'page1',
              label: '页面1',
              icon: 'setting',
              url: 'Other/PageOne'
            },
            {
              path: '/page2',
              name: 'page2',
              label: '页面2',
              icon: 'setting',
              url: 'Other/PageTwo'
            }
          ]
        }

      ]
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    // 点击菜单
    clickMenu(item) {
      console.log(item.label);
      // 解决方法1
      // this.$router.push(item.path).then(response => {}, error => {})
      // 解决方法2
      if (this.$route.path !== item.path && !(this.$route.path === '/home' && (item.path === '/')))
        this.$router.push(item.path)
      // 修改state中的数据
      // this.$store.commit('selectMenu', item)
      this.$store.commit('tab/selectMenu', item)
    },
    
  },
  computed: {
    // ...mapMutations('tab',{
    //   selectMenu: 'selectMenu'
    // }),
    // 没有子菜单
    noChildren() {
      return this.menuData.filter(item => !item.children);
    },
    // 有子菜单
    hasChildren() {
      return this.menuData.filter(item => item.children);
    },
    isCollapse() {
      return this.$store.state.tab.isCollapse
    }
  }
}
</script>

<style lang="less" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
  text-align: left;
}

.el-menu {
  height: 100vh;
  border-right: none;

  h3 {
    color: white;
    background-color: #2f477f;
    text-align: center;
    line-height: 50px;
    font-size: 16px;
    font-weight: 400;
  }
}
</style>
