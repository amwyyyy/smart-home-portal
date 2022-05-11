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
    <section>
        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-form :inline="true">
                <el-form-item>
                    <el-select v-model="queryModel.serviceId" @change="fetchEndpoint" filterable allow-create placeholder="应用名称">
                        <el-option label="请选择" value=""></el-option>
                        <el-option v-for="item in serviceList" :label="item.name" :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-select v-model="queryModel.endpointId" filterable allow-create placeholder="接口名称">
                        <el-option label="请选择" value=""></el-option>
                        <el-option v-for="item in endpointList" :label="item.name" :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-select v-model="queryModel.error" filterable allow-create placeholder="是否异常">
                        <el-option label="请选择" value=""></el-option>
                        <el-option label="Success" value="false"></el-option>
                        <el-option label="Error" value="true"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-select v-model="queryModel.status" filterable allow-create placeholder="是否处理">
                        <el-option label="请选择" value=""></el-option>
                        <el-option label="未处理" value="0"></el-option>
                        <el-option label="处理中" value="1"></el-option>
                        <el-option label="已完成" value="2"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="el-icon-search" @click="query">搜索</el-button>
                </el-form-item>
            </el-form>
        </el-col>

        <!--列表-->
        <el-table :data="segmentList" stripe border style="width: 100%">
            <el-table-column type="index" label="序号" width="50"></el-table-column>
            <el-table-column prop="segmentObject.service_" label="应用名称" width="150"></el-table-column>
            <el-table-column prop="endpoint_name" label="接口名称" width="300"></el-table-column>
            <el-table-column prop="latency" label="耗时" width="70"></el-table-column>
            <el-table-column label="是否异常" width="100">
                <template slot-scope="scope">
                    <el-tag v-if="scope.row.is_error===0" type="success">Success</el-tag>
                    <el-tag v-if="scope.row.is_error===1" type="danger">Error</el-tag>
                </template>
            </el-table-column>
            <el-table-column prop="errorKind" label="异常类型" width="200"></el-table-column>
            <el-table-column label="状态" width="70">
                <template slot-scope="scope">
                    {{ statusMap[scope.row.status] }}
                </template>
            </el-table-column>
            <el-table-column prop="plan" label="解决方案" width="150"></el-table-column>
            <el-table-column prop="remark" label="备注" width="150"></el-table-column>
            <el-table-column prop="insertTimeDate" label="新增时间" width="160"></el-table-column>
            <el-table-column prop="updateTimeDate" label="更新时间" width="160"></el-table-column>
            <el-table-column label="操作" width="280">
                <template slot-scope="scope">
                    <el-button size="mini" type="info" round plain @click="openDialog(scope.row)">查看详情</el-button>
                    <el-button v-if="scope.row.status===0" size="mini" type="primary" round plain @click="handler(scope.row)">处理</el-button>
                    <el-button v-if="scope.row.status===1" size="mini" type="primary" round plain @click="finish(scope.row)">完成</el-button>
                    <el-button v-if="scope.row.status===0" size="mini" type="danger" round plain @click="finish(scope.row)">不处理</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!--分页-->
        <page :query-param="queryModel" :page-size="pageSize" :total="pageTotal" @pagination="query"></page>

        <!--编辑弹框-->
        <el-dialog title="链路详情" :visible.sync="showDialog" width="70%">
            <el-form label-width="100px" v-if="showDialog" size="mini">
                <div v-for="span in showRow.segmentObject.spans_">
                    <el-form-item size="mini" label="接口: ">
                        <div class="text">
                            {{ span.operationName_ }}
                        </div>
                    </el-form-item>
                    <el-form-item size="mini" label="跨度类型: ">
                        <div class="text">
                            {{ spanType[span.spanType_] }}
                        </div>
                    </el-form-item>
                    <el-form-item size="mini" label="组件: ">
                        <div class="text">
                            {{ spanLayer[span.spanLayer_] }}
                        </div>
                    </el-form-item>
                    <el-form-item size="mini" label="Peer: ">
                        <div class="text">
                            {{ span.peer_ }}
                        </div>
                    </el-form-item>
                    <el-form-item size="mini" label="是否异常: ">
                        <el-tag v-if="span.isError_" type="danger">Error</el-tag>
                        <el-tag v-else type="success">Success</el-tag>
                    </el-form-item>
                    <el-form-item size="mini" v-for="tag in span.tags_" :label="tag.key_ + ': '">
                        <codemirror v-if="tag.key_ === 'db.statement'" v-model="tag.value_" :options="cmOptions" />
                        <div v-else class="text">
                            {{ tag.value_ }}
                        </div>
                    </el-form-item>
                    <div v-for="item in span.logs_">
                        <hr style="border-top:1px dashed black;">
                        <el-form-item size="mini" v-for="log in item.data_" :label="log.key_ + ': '">
                            <codemirror v-if="log.key_ === 'stack' || log.key_ === 'message'" v-model="log.value_" :options="cmOptions" />
                            <div v-else class="text">
                                {{ log.value_ }}
                            </div>
                        </el-form-item>
                    </div>
                    <hr/>
                </div>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="showDialog = false">确 定</el-button>
            </span>
        </el-dialog>
    </section>
</template>

<script>
import util from '../../common/util'
import Page from '../common/Page.vue'
import NProgress from "nprogress"
import 'nprogress/nprogress.css'

import 'codemirror/mode/yaml/yaml.js'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/base16-light.css'
import '../../common/codemirror.css'

export default {
        data () {
            return {
                queryModel: {
                    serviceId: '',
                    endpointId: '',
                    error: null,
                    status: null,
                    pageNum: 1,
                    pageSize: 20
                },
                serviceList: [],
                endpointList: [],
                segmentList: [],
                pageSize: 1,
                pageTotal: 10,
                showDialog: false,
                showRow: {},
                spanType: {
                    0: "Entry",
                    1: "Exit",
                    2: "Local"
                },
                spanLayer: {
                    0: "Unknown",
                    1: "Database",
                    2: "RPCFramework",
                    3: "Http",
                    4: "MQ",
                    5: "Cache",
                },
                statusMap: {
                    0: "未处理",
                    1: "处理中",
                    2: "完成"
                },
                cmOptions: {
                    tabSize: 2,
                    mode: 'text/yaml',
                    theme: 'base16-light',
                    lineNumbers: false,
                    line: true,
                    readOnly: true
                }
            }
        },
        methods: {
            query(queryModel) {
                NProgress.start();
                let params = this.queryModel;
                if (Object.prototype.toString.call(queryModel) === '[object Object]') {
                    params = queryModel;
                }

                loadData(this, params)
            },
            fetchEndpoint(val) {
                if (val === null || val === '' || val === undefined) {
                    return;
                }

                let _this = this;
                util.get('/segment/endpoint_list', {serviceId: val}, (result) => {
                    _this.endpointList = result;
                });
            },
            handler(row) {
                this.$prompt('输入解决方案', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消'
                }).then(({ value }) => {
                    util.post('/segment/update_status', {traceId: row.trace_id, status: 1, plan: value}, () => {
                        this.$message({
                            type: 'success',
                            message: '操作成功'
                        });
                        row.status = 1;
                        row.plan = value;
                    });
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
            },
            finish(row) {
                this.$prompt('输入备注信息', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消'
                }).then(({ value }) => {
                    util.post('/segment/update_status', {traceId: row.trace_id, status: 2, remark: value}, () => {
                        this.$message({
                            type: 'success',
                            message: '操作成功'
                        });
                        row.status = 2;
                        row.remark = value;
                    });
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
            },
            openDialog(row) {
                this.showDialog = true;
                this.showRow = row;
            }
        },
        components: {
            Page
        },
        mounted() {
            NProgress.start();
            let _this = this;
            util.get('/segment/service_list', '', (result) => {
                _this.serviceList = result;
            });

            loadData(this, this.queryModel);
        }
    }

    function loadData(vue, params) {
        let _this = vue;
        if (params.status === "") {
            params.status = null;
        }

        if (params.error === "") {
            params.error = null;
        }

        util.post('/segment/segment_list', params, (result) => {
            NProgress.done();
            _this.segmentList = result.list;
            _this.pageSize = result.pageSize;
            _this.pageTotal = result.total;
        });
    }
</script>
<style scoped>
    .toolbar {
        background: #fff;
        padding: 10px 10px 0px 10px;
    }

    .text {
        width: 100%;
        word-wrap: break-word;
    }
</style>