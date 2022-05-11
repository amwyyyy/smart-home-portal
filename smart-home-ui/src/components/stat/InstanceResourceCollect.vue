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
    <!--工具条-->
    <el-form :inline="true" class="toolbar">
      <el-form-item>
        <el-select style="width: 350px" v-model="businessLine" collapse-tags filterable allow-create placeholder="请选择业务线">
          <el-option v-for="item in businessList" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" @click="query">搜索</el-button>
      </el-form-item>
    </el-form>

    <!--列表-->
    <el-table :data="instanceResourceList" border style="width: 100%;" :row-class-name="tableRowClassName">
      <el-table-column prop="service_name" label="服务名称" width="250"></el-table-column>
      <el-table-column prop="ip_addr" label="IP 地址" width="150"></el-table-column>
      <el-table-column prop="avg_cpu" label="avg cpu (%)" width="120"></el-table-column>
      <el-table-column prop="max_cpu" label="max cpu (%)" width="120"></el-table-column>
      <el-table-column prop="avg_physical_cpu" label="avg system cpu(%)" width="155"></el-table-column>
      <el-table-column prop="max_physical_cpu" label="max system cpu(%)" width="155"></el-table-column>
      <el-table-column prop="avg_memory" label="平均内存使用率(%)" width="150"></el-table-column>
      <el-table-column prop="max_memory" label="最大内存使用率(%)" width="150"></el-table-column>
      <el-table-column prop="total_memory" label="总堆内存(MB)" width="120"></el-table-column>
      <el-table-column prop="max_cpm" label="峰值请求量(分钟)" width="130"></el-table-column>
      <el-table-column prop="total_request" label="总请求量(日)" width="100"></el-table-column>
      <el-table-column prop="current" label="当前配置" width="100"></el-table-column>
      <el-table-column prop="recommend" label="建议配置" width="100"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import util from "../../common/util";

export default {
  data() {
    return {
      businessLine: '',
      businessList: [],
      instanceResourceList: []
    }
  },
  methods: {
    query() {
      let _this = this;
      util.post('/report/resource_collect', {'businessLine': _this.businessLine}, (result) => {
        _this.instanceResourceList = result;
      });
    },
    tableRowClassName({row, rowIndex}) {
      if (row.max_cpu >= 85 || row.max_memory >= 85) {
        return 'warning-row';
      }
      if ((row.max_cpu <= 10 && row.max_memory <= 80) || row.max_cpm < 150 || row.total_request < 200000) {
        return 'success-row';
      }

      return '';
    }
  },
  mounted() {
    let _this = this;
    util.get('/report/business_line', '', (result) => {
      _this.businessList = result;
    });
  }
}
</script>

<style>
.toolbar {
  background: #fff;
}

.el-table .warning-row {
  background-color: #fd8585;
}

.el-table .success-row {
  background-color: #b4fa8e;
}
</style>