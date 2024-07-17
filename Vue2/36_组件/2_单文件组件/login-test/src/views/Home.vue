<script>
export default {
  name: '',
  data: function () {
    return {
      tableData: [
        {
          name: 'aaa',
          todayBuy: 344,
          monthBuy: 564,
          totalBuy: 987,
        },
        {
          name: 'bbb',
          todayBuy: 236,
          monthBuy: 560,
          totalBuy: 350,
        },
        {
          name: 'ccc',
          todayBuy: 780,
          monthBuy: 750,
          totalBuy: 810,
        },{
          name: 'ddd',
          todayBuy: 175,
          monthBuy: 546,
          totalBuy: 924,
        }
      ],
      tableLabel: {
        name: '课程',
        todayBuy: '今日购买',
        monthBuy: '当月购买',
        totalBuy: '总购买',
      },
      countData: [
        {
          name: '今日支付订单',
          value: 1234,
          icon: 'success',
          color: '#2ec7c9',
        },
        {
          name: '今日收藏订单',
          value: 210,
          icon: 'star-on',
          color: '#ffb980',
        },
        {
          name: '今日未支付订单',
          value: 1204,
          icon: 's-goods',
          color: '#5ab1ef',
        },
        {
          name: '本月支付订单',
          value: 134,
          icon: 'success',
          color: '#2ec7c9',
        },
        {
          name: '本月收藏订单',
          value: 2130,
          icon: 'star-on',
          color: '#ffb980',
        },
        {
          name: '本月未支付订单',
          value: 1204,
          icon: 's-goods',
          color: '#5ab1ef',
        },
      ]
    }
  },
}
</script>

<template>
  <el-row>
    <el-col :span="8">
      <el-card class="box-card">
        <div class="user">
          <img src="../assets/images/user1.jpg" alt="">
          <div class="userinfo">
            <p class="name">Admin</p>
            <p class="access">超级管理员</p>
          </div>
        </div>
        <div class="login-info">
          <p>上次登录时间：<span>2024-01-01</span></p>
          <p>上次登录地点：<span>武汉</span></p>
        </div>
      </el-card>
      <el-card class="box-card" style="margin-top: 20px; height: 460px" >
        <el-table
            :data="tableData"
            style="width: 100%">
          <el-table-column v-for="(key, value) in tableLabel" :key="key" :prop="value" :label="key">
          </el-table-column>

        </el-table>
      </el-card>
    </el-col>
    <el-col :span="16">
      <div class="num">
        <el-card v-for="item in countData" :key="item.name">
          <i class="icon" :class="`el-icon-${item.icon}`"></i>
        </el-card>
      </div>
    </el-col>
  </el-row>
</template>

<style lang="less" scoped>
.user {
  display: flex;
  align-items: center;
  border-bottom: 1px solid #ccc;
  padding-bottom: 20px;
  margin-bottom: 20px;

  img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    margin-right: 40px;
  }
}

.userinfo {
  .name {
    font-size: 32px;
    margin-bottom: 10px;
    color: #000000;
  }

  .access {
    color: #999999;
  }
}

p {
  line-height: 28px;
  font-size: 14px;
  color: #999999;

  span {
    color: #666666;
    margin-left: 60px;
  }
}
</style>