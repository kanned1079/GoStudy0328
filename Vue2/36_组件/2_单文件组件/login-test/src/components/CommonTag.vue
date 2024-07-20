<script>
import index, {mapMutations, mapState} from 'vuex'
import router from "@/router";

export default {
  name: 'CommonTag',
  computed: {
    ...mapState('tab', ['tabList']),
  },
  mounted() {
    // console.log('tabList', this.tabList)
  },
  methods: {
    ...mapMutations('tab', ['closeTag']),
    changeMenu(item) {
      console.log('item', item)
      console.log(this.$route.path, item.path)
      if (item.path !== this.$route.path)
        this.$router.push({
          name: item.name,
        }).then(()=>{}, ()=>{})
    },
    // 点击删除Tag的功能
    handleClose(item, index ) {
      let length = this.tabList.length - 1  // 获取长度
      this.closeTag(item)
      // 删除后的跳转逻辑
      if (item.name !== this.$route.name) {
        return
      }
      // 删除最后一项
      if (index === length) {
        this.$router.push({
          name: this.tabList[index - 1].name
        })
      } else {  // 删除的是中间内容 应该向后跳转
        this.$router.push({
          name: this.tabList[index].name
        })
      }

    }
  }
}
</script>

<template>
  <div class="tabs">
    <el-tag
        v-for="(item, index) in tabList"
        :key="item.path"
        :closable="item.name !== 'home'"
        :effect="$route.name === item.name?'dark':'plain'"
        :type="item.type"
        @click="changeMenu(item)"
        @close="handleClose(item, index)"
        size="small">
      {{ item.label }}
    </el-tag>
  </div>
</template>

<style lang="less" scoped>
.tabs {
  padding: 20px;
  .el-tag {
    margin-right: 15px;
    cursor: pointer;
  }
}
</style>