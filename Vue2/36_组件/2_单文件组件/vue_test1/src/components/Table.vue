<script>
export default {
  name: 'PersonTable',
  data() {
    return {
      personList: [
        {id: '001', age: 18, sex: '男', name: '张三'},
        {id: '002', age: 19, sex: '女', name: '张四'},
        {id: '003', age: 21, sex: '男', name: '张五'},
      ]
    }
  }
}
</script>

<template>
<div>
  <el-container style="height: 500px; border: 1px solid #eee">


    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 15px"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>查看</el-dropdown-item>
            <el-dropdown-item>新增</el-dropdown-item>
            <el-dropdown-item>删除</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span>王小虎</span>
      </el-header>

      <el-main>
        <el-table :data="personList">
          <el-table-column prop="id" label="日期" width="140">
          </el-table-column>
          <el-table-column prop="age" label="姓名" width="120">
          </el-table-column>
          <el-table-column prop="sex" label="地址">
          </el-table-column>
          <el-table-column prop="name" label="地址">
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </el-container>
</div>
</template>

<style scoped>
.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>