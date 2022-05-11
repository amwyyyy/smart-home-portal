<!--
  - Copyright (C) 2011-2020 ShenZhen iBOXCHAIN Information Technology Co.,Ltd.
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
    <!--分页插件-->
    <el-col :span="24" class="toolbar" style="padding-bottom:10px;">
        <el-pagination :current-page="1" :page-sizes="[10, 20, 30, 50]" :page-size="pageSize"
                       :total="total" @size-change="sizeChange" @current-change="currentChange"
                       style="float:right" background layout="total, sizes, prev, pager, next"/>
    </el-col>
</template>
<script>
    export default {
        data() {
            return {}
        },
        props: ['queryParam', 'pageSize', 'total'],
        methods: {
            sizeChange(pageSize) {
                this.queryParam.pageSize = pageSize;
                this.queryParam.pageNum = 1;

                this.$emit('pagination', this.queryParam);
            },
            currentChange(pageNum) {
                this.queryParam.pageNum = pageNum;

                this.$emit('pagination', this.queryParam);
            }
        }
    }
</script>