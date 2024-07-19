<script>
import {getData} from '../api'
import * as echarts from 'echarts'
export default {
  name: '',
  data: function () {
    return {
      data: {},
      tableData: [],
      tableLabel: {
        name: '课程',
        todayBuy: '今日购买',
        monthBuy: '当月购买',
        totalBuy: '总购买',
      },
      countData: [],
      orderData: {},
    }
  },
  mounted() {
    getData().then(({data}) => {
      // console.log(data)
      const { tableData, orderData, videoData, columnData, piData } = data
      this.tableData = tableData  // ok
      this.countData = orderData  // ok
      // console.log(videoData)  // 在此处可以正常输出
      this.orderData = videoData
      { // 折线图生成
        const chart1 = echarts.init(this.$refs.echaerts1);
        let chart1_option = {}
        let xAxis = Object.keys(videoData.data[0])
        const xAxisData = {
          data: xAxis,
        }
        chart1_option.xAxis = xAxisData
        chart1_option.yAxis = {}
        chart1_option.legend = xAxisData
        chart1_option.tooltip = {}
        chart1_option.series = []
        xAxis.forEach(key => {
          chart1_option.series.push({
            name: key,
            data: videoData.data.map(item => item[key]),
            type: 'line',
          })
        })
        // console.log(chart1_option)
        // 根据配置显示图表
        chart1.setOption(chart1_option)
      }
      { // 柱状图生成
        const chart2 = echarts.init(this.$refs.echaerts2);
        const chart2_option = {
          // title: {
          //   text: '用户使用量',
          //   // subtext: 'Fake Data',
          //   left: 'left'
          // },
          legend: {
            textStyle: {
              color: '#333',
            }
          },
          grid: {
            left: '20%',
          },
          tooltip: {
            trigger: 'axis',
          },
          xAxis: {
            type: 'category',
            data: columnData.map(item => item.date),
            axisLine: {
              lineStyle: {
                color: '#17b3a3',
              }
            },
            axisLabel: {
              interval: 0,
              color: '#333',
            }
          },
          yAxis: [
            {
              type: 'value',
              axisLine: {
                lineStyle: {
                  color: '#17b3a3'
                }
              }
            }
          ],
          color: ['#2ec79c', '#b6a2de'],
          series: [
            {
              name: '新增用户',
              data: columnData.map(item => item.new),
              type: 'bar',
            },
            {
              name: '活跃用户',
              data: columnData.map(item => item.active),
              type: 'bar',
            }
          ]
        }
        chart2.setOption(chart2_option)
      }
      { // 饼状图
        let chart3 = echarts.init(this.$refs.echaerts3);
        let chart3_option = {
          title: {
            text: '用户使用量',
            // subtext: 'Fake Data',
            left: 'center'
          },
          tooltip: {
            trigger: 'item'
          },
          legend: {
            orient: 'vertical',
            left: 'left'
          },
          series: [
            {
              name: 'Access From',
              type: 'pie',
              radius: '50%',
              data: piData,
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              }
            }
          ]
        }
        chart3.setOption(chart3_option)
      }

    }, error => {
      console.log(error.message)
    })
  }
}
</script>

<template>
  <el-row>
    <el-col :span="8" style="padding-right: 10px">
      <el-card class="box-card" style="height: 280px">
        <div class="user">
          <img src="../assets/images/user1.jpg" alt="">
          <div class="userinfo">
            <p class="name">Admin</p>
            <p class="access">超级管理员</p>
          </div>
        </div>
        <div class="login-info">
          <p>上次登录时间：<span>2024-01-01</span></p>
          <p>上次登录地点：<span>中国大陆</span></p>
        </div>
      </el-card>
      <el-card class="box-card" style="margin-top: 20px; height: 460px">
        <el-table
            :data="tableData"
            style="width: 100%">
          <el-table-column v-for="(key, value) in tableLabel" :key="key" :prop="value" :label="key">
          </el-table-column>

        </el-table>
      </el-card>
    </el-col>
    <el-col :span="16" style="padding-left: 10px">
      <div class="num">
        <el-card v-for="item in countData" :key="item.name" :body-style="{ display: 'flex', padding: 0}">
          <i class="icon" :class="`el-icon-${item.icon}`" v-bind:style="{ backgroundColor: item.color }"></i>
          <div class="detail">
            <p class="price"> ¥ {{ item.value }} </p>
            <p class="desc"> {{ item.name }} </p>
          </div>
        </el-card>
      </div>
      <el-card style="height: 280px">
        <!--        折线图的内容-->
        <div ref="echaerts1" style="height: 280px"></div>
      </el-card>
      <div class="graph">
        <el-card style="height: 260px">
          <div ref="echaerts2" style="height: 260px"></div>
        </el-card>
        <el-card style="height: 260px">
          <div ref="echaerts3" style="height: 260px"></div>
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

.num {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  .icon {
    width: 80px;
    height: 80px;
    font-size: 30px;
    color: white;
    align-items: center;
    line-height: 80px;
    text-align: center;
  }

  .detail {
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin-left: 15px;
    margin-right: 30px;

    .price {
      font-size: 30px;
      margin-bottom: 3px;
      color: rgba(0, 0, 0, 0.63);
      line-height: 30px;
      height: 30px;
    }

    .desc {
      font-size: 14px;
      color: #999;
      text-align: center;
    }
  }

  .el-card {
    width: 32%;
    margin-bottom: 20px;
  }
}

.graph {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
  .el-card {
    width: 48.5%;
  }
}


</style>