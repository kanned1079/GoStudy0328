<script>
import {mapState} from 'vuex'
import router from "@/router";

export default {
  name: 'CommonTag',
  computed: {
    ...mapState('tab', ['tabList'])
  },
  mounted() {
    // console.log('tabList', this.tabList)
  },
  methods: {
    changeMenu(item) {
      console.log('item', item)
      console.log(this.$route.path, item.path)
      this.$router.push({
        name: item.name,
      })
    },
    closeMenu(item) {
      console.log('关闭Tag')
    }
  }
}
</script>

<template>
  <div class="tabs">
    <el-tag style="margin-right: 10px; height: 30px; line-height: 30px; align-items: center"
        v-for="item in tabList"
        :key="item.path"
        :closable="item.name !== 'home'"
        :effect="$route.name === item.name?'dark':'plain'"
        :type="item.type"
        @click="changeMenu(item)"
        @close="closeMenu(item)">
      {{ item.label }}
    </el-tag>
  </div>
</template>

<style lang="less" scoped>
.tabs {
  el-tag {
    padding-right: 20px;
  }
}
</style>