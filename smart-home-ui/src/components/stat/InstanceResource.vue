<!--
  - Copyright (C) 2011-2021 ShenZhen iBOXCHAIN Information Technology Co.,Ltd.
  -
  - All right reserved.
  -
  - This software is the confidential and proprietary
  - information of iBOXCHAIN Company of China.
  - ("Confidential Information"). You shall not disclose
  - such Confidential Information and shall use it only
  - in accordance with the terms of the contract agreement
  - you entered into with iBOXCHAIN inc.
  -->

<template>
  <div>
    <el-form :inline="true" class="toolbar">
      <el-form-item>
        <el-select style="width: 350px" v-model="queryModel.serviceIds" multiple collapse-tags filterable allow-create placeholder="应用名称">
          <el-option v-for="item in serviceList" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-date-picker
            v-model="queryModel.dateTime"
            type="daterange"
            align="right"
            unlink-panels
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :picker-options="pickerOptions">
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" @click="query">搜索</el-button>
      </el-form-item>
    </el-form>
    <div id="cpuAvg" class="graph"></div>
    <div id="cpuMax" class="graph"></div>
    <div id="memoryAvg" class="graph"></div>
    <div id="memoryMax" class="graph"></div>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import util from "../../common/util";

function line(data, id, title) {
  let chart = echarts.init(document.getElementById(id), 'light');

  const num = data[0].length;
  let seriesArr = [];
  for (let i = 0; i < num - 1; i++) {
    seriesArr[i] = {type: 'line'};
  }

  chart.setOption({
    title: {
      text: title
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      itemWidth: 5,
      formatter: function (name) {
        if (!name) return '';
        if (name.length > 8) {
          name =  name.slice(0,8) + '...';
        }
        return name;
      },
      tooltip: {
        show: true
      }
    },
    grid: {
      left: '3%',
      right: '3%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false
    },
    yAxis: {},
    dataset: {
      source: data
    },
    series: seriesArr
  }, true);
}

function bar(data, id, title) {
  const num = data[0].length;
  let seriesArr = [];
  for (let i = 0; i < num - 1; i++) {
    seriesArr[i] = {type: 'bar'};
  }

  let chart = echarts.init(document.getElementById(id), 'light');
  chart.setOption({
    title: {
      text: title
    },
    tooltip: {
      trigger: 'axis'
    },
    dataset: {
      source: data
    },
    grid: {
      left: '3%',
      right: '3%',
      bottom: '3%',
      containLabel: true
    },
    legend: {
      itemWidth: 3,
      formatter: function (name) {
        if (!name) return '';
        if (name.length > 8) {
          name =  name.slice(0,8) + '...';
        }
        return name;
      },
      tooltip: {
        show: true
      }
    },
    xAxis: {type: 'category'},
    yAxis: {},
    series: seriesArr
  }, true)
}

export default {
  data () {
    return {
      queryModel: {
        serviceIds: [],
        dateTime: ''
      },
      serviceList: [],
      pickerOptions: {
        shortcuts: [{
          text: '最近一周',
          onClick(picker) {
            const end = new Date();
            end.setTime(end.getTime() - 3600 * 1000 * 24);
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
            picker.$emit('pick', [start, end]);
          }
        }, {
          text: '最近一个月',
          onClick(picker) {
            const end = new Date();
            end.setTime(end.getTime() - 3600 * 1000 * 24);
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
            picker.$emit('pick', [start, end]);
          }
        }, {
          text: '最近三个月',
          onClick(picker) {
            const end = new Date();
            end.setTime(end.getTime() - 3600 * 1000 * 24);
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
            picker.$emit('pick', [start, end]);
          }
        }]
      }
    }
  },
  methods: {
    query() {
      let _this = this;
      let params = {
        serviceIds: _this.queryModel.serviceIds,
        beginTime: _this.queryModel.dateTime[0],
        endTime: _this.queryModel.dateTime[1]
      }
      util.post('/report/resource_stat', params, (result) => {
        line(result['cpuAvg'], 'cpuAvg', 'CPU 平均使用率');
        bar(result['cpuMax'], 'cpuMax', 'CPU 最大使用率');
        line(result['memoryAvg'], 'memoryAvg', '内存平均使用率');
        bar(result['memoryMax'], 'memoryMax', '内存最大使用率');
      })
    }
  },
  mounted() {
    let _this = this;
    util.get('/report/service_list', '', (result) => {
      _this.serviceList = result;
    });
  }
}
</script>
<style scoped>
.toolbar {
  background: #fff;
}

.graph {
  width: 49.8%;
  height: 400px;
  background: #fff;
  display: inline-block;
}
</style>