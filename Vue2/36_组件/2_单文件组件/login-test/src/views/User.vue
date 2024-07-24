<script>
import {addUser, delUser, editUser, getUsers} from "@/api";

export default {
  name: '',
  data: function () {
    return {
      dialogVisible: false,
      form: {
        // id: '',
        name: '',
        age: '',
        sex: '',
        birth: '',
        addr: '',
      },
      userform: {
        name: '',
      },
      rules: {
        name: [
          {required: true, message: '姓名为必填'},
        ],
        age: [
          {required: true, message: '年龄为必填'},
        ],
        sex: [
          {required: true, message: '性别为必选项'},
        ],
        birth: [
          {required: true, message: '请选择出生日期'},
        ],
        // addr: [
        //   {required: true, message: '地址为必填'},
        // ],
      },
      tableData: [],
      modalType: 0, // 0表示新增 1表示编辑
      total: 0,   // 当前的记录总条数
      pageData: {
        page: 1,
        limit: 10,
      },
    }
  },
  methods: {
    handleClose() {
      // this.$confirm('确认关闭？')
      //     .then(_ => {
      //       done();
      //     })
      //     .catch(_ => {
      //     });
      this.dialogVisible = false; // 关闭窗口
    },
    handleCancel() {
      this.handleClose()
      console.log('handleClose 清空表单');
      this.$refs.form.resetFields(); // 清空表单
    },
    handleSubmit() {
      console.log('handleSubmit');
      this.$refs.form.validate((valid) => {
        if (valid) {  // valid是bool校验通过返回true
          // 后续是对表单的处理
          if (this.modalType === 0) {
            // 如果是modalType是0 将调用新增用户的接口
            console.log('this.form', this.form)
            addUser(this.form).then(() => {
              // 更新后要刷新用户列表
              this.getList()
            })
          } else {
            // 如果是modalType是1 将调用修改用户的接口
            editUser(this.form).then(() => {
              this.getList()
            })
          }
          console.log('用户表单', this.form)
          // this.dialogVisible = false; // 填写完成后关闭弹窗
        }
      })
      this.handleClose()
    },
    handleEdit(row) { // 每一项的修改按钮
      this.modalType = 1;
      this.dialogVisible = true;
      console.log('edit user function')
      // 回显数据到表单
      this.form = JSON.parse(JSON.stringify(row))
      console.log(this.form)
    },
    handleDel(row) {  // 每一项的删除按钮
      console.log('del user function')
      console.log(row)

      this.$alert(`确定删除${row.name}的记录吗`, '警告', {
        confirmButtonText: '确定',
        callback: action => {
          delUser(row).then(() => {
            this.getList()
          })
          this.$message({
            type: 'success',
            message: `删除成功`
          });
        }
      });
    },
    getList() {
      // 专门用户更新列表
      getUsers({
        params: {
          ...this.userform, ...this.pageData
        }
      }).then(({data}) => {
        console.log(data)
        this.tableData = data.data;
        this.total = data.count ? data.count : 0; // 列表的记录数
        console.log('数据条数：', this.total);
      }, error => {
      })
    },
    handleAdd() { // 顶部的新增按钮
      this.modalType = 0;
      this.dialogVisible = true
    },
    handlePage(page) {  // 选择页码的回调函数
      console.log('页码：', page);
      this.pageData.page = page;
      this.getList();
    },
    onSearch() {  // 列表查询条件
      this.getList()
    }

  },
  mounted() {
    this.getList()
  }

}
</script>

<template>
  <div class="manage">
    <el-dialog
        title="新增用户"
        :visible.sync="dialogVisible"
        width="40%"
        :before-close="handleClose">
      <!--      用户的表单信息-->
      <el-form ref="form" :inline="true" :model="form" label-width="80px" :rules="rules">
        <el-form-item label="姓名" prop="name">
          <el-input placeholder="请输入姓名" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="年龄" prop="age">
          <el-input placeholder="请输入年龄" v-model.number="form.age"></el-input>
        </el-form-item>
        <el-form-item label="性别" prop="sex">
          <el-select v-model.number="form.sex" placeholder="请选择性别">
            <el-option label="男" :value=1></el-option>
            <el-option label="女" :value=0></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="出生日期" prop="birth">
          <el-date-picker v-model="form.birth" type="date" placeholder="选择日期"></el-date-picker>
        </el-form-item>
        <el-form-item label="地址" prop="addr">
          <el-input placeholder="请输入地址" v-model="form.addr"></el-input>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
              <el-button @click="handleCancel">取 消</el-button>
              <el-button type="primary" @click="handleSubmit">确 定</el-button>
          </span>
    </el-dialog>

    <div class="manager_header">
      <el-button type="primary" @click="handleAdd">
        +新增
      </el-button>
      <!--      form搜索区域-->
      <el-form :model="userform" :inline="true">
        <el-form-item>
          <el-input v-model="userform.name" placeholder="请输入过滤的姓名"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch" size="medium">查询</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="common-table">
      <el-table
          :data="tableData"
          stripe
          style="width: 100%; border-radius: 10px;"
          height="90%"
      >
        <!--      <el-table-column-->
        <!--          prop="id"-->
        <!--          label="ID"-->
        <!--          width="120px">-->
        <!--      </el-table-column>-->
        <el-table-column
            prop="name"
            label="姓名"
            width="140px">
        </el-table-column>
        <el-table-column
            prop="age"
            label="年龄"
            width="120px">
        </el-table-column>
        <el-table-column
            prop="sex"
            label="性别">
          <template slot-scope="scope">
            <span style="margin-left: 10px"> {{ scope.row.sex === 1 ? '男' : '女' }} </span>
          </template>
        </el-table-column>
        <el-table-column
            prop="birth"
            label="生日">
        </el-table-column>
        <el-table-column
            prop="addr"
            label="家庭住址">
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="info" @click="handleEdit(scope.row)" size="mini">修改</el-button>
            <el-button type="danger" @click="handleDel(scope.row)" size="mini">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pager">
        <el-pagination
            layout="prev, pager, next"
            :total="total"
            @current-change="handlePage">

        </el-pagination>
      </div>
    </div>


  </div>
</template>

<style lang="less" scoped>
.manage {
  height: 100%;
  .manager_header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .common-table {
    position: relative;
    height: 700px;
    .pager {
      position: absolute;
      bottom: 0;
      right: 20px;

    }
  }
}



</style>