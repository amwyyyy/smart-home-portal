<template>
    <section>
        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-form :inline="true">
                <el-form-item>
                    <el-select v-model="queryModel.serviceName" filterable allow-create placeholder="服务名称">
                        <el-option label="请选择" value=""></el-option>
                        <el-option v-for="item in serviceNameList" :label="item.name" :value="item.name"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="el-icon-search" @click="query">查找</el-button>
                </el-form-item>
            </el-form>
        </el-col>

        <!--列表-->
        <el-table :data="mailSubList" stripe border style="width: 100%">
            <el-table-column type="index" label="序号" width="50"></el-table-column>
            <el-table-column prop="serviceName" label="服务名称" width="250"></el-table-column>
            <el-table-column prop="mailAddress" label="订阅邮箱地址" width="350"></el-table-column>
            <el-table-column prop="endpoints" label="订阅接口地址" width="350"></el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button size="mini" type="primary" plain @click="openDialog(scope.row)">更新邮箱地址</el-button>
                    <el-button size="mini" type="success" plain @click="openEndpointDialog(scope.row)">更新接口地址</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!--更新邮箱地址弹框-->
        <el-dialog title="更新邮箱地址" :visible.sync="showDialog" :close-on-click-modal="false" width="50%">
            <el-table :data="mailList" stripe border style="width: 100%">
                <el-table-column type="index" label="序号" width="50"></el-table-column>
                <el-table-column prop="addr" label="邮箱地址" width="200"></el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button size="mini" type="danger" plain @click="delMailAddr(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-row :gutter="24" style="margin-top: 10px">
                <el-col :span="12"><el-input v-model="mailAddr" placeholder="邮箱地址"/></el-col>
                <el-col :span="6"><el-button @click="addMailAddr()">添加</el-button></el-col>
            </el-row>

            <span slot="footer" class="dialog-footer">
                    <el-button type="primary" @click="showDialog=false;query()">确 定</el-button>
                </span>
        </el-dialog>

        <!--更新接口地址弹框-->
        <el-dialog title="更新接口地址" :visible.sync="showEndpointDialog" :close-on-click-modal="false" width="50%">
            <el-table :data="endpointList" stripe border style="width: 100%">
                <el-table-column type="index" label="序号" width="50"></el-table-column>
                <el-table-column prop="endpoint" label="接口地址" width="200"></el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button size="mini" type="danger" plain @click="delEndpoint(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-row :gutter="24" style="margin-top: 10px">
                <el-col :span="12"><el-input v-model="endpoint" placeholder="接口地址"/></el-col>
                <el-col :span="6"><el-button @click="addEndpoint()">添加</el-button></el-col>
            </el-row>

            <span slot="footer" class="dialog-footer">
                    <el-button type="primary" @click="showEndpointDialog=false;query()">确 定</el-button>
                </span>
        </el-dialog>
    </section>
</template>

<script>
import util from '../../common/util'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css';

export default {
        data () {
            return {
                queryModel: {
                    serviceName: ''
                },
                serviceNameList: [],
                mailSubList: [],
                showDialog: false,
                showEndpointDialog: false,
                currentServiceName: '',
                currentServiceId: null,
                currentMailAddr: '',
                currentEndpoint: '',
                mailList: [],
                mailAddr: '',
                endpoints: [],
                endpointList: [],
                endpoint: ''
            }
        },
        methods: {
            query() {
                NProgress.start();
                let _this = this;
                util.get('/mailSubList', {serviceName: this.queryModel.serviceName}, result => {
                    NProgress.done();
                    _this.mailSubList = result;
                    for (let idx in _this.mailSubList) {
                        if (_this.mailSubList[idx].endpointList != null) {
                            _this.mailSubList[idx].endpoints = _this.mailSubList[idx].endpointList.join(",")
                        }
                    }
                });
            },
            openDialog(row) {
                this.showDialog = true;
                this.currentServiceName = row.serviceName;
                this.currentServiceId = row.serviceId;
                this.currentEndpoint = row.endpointList;

                this.mailList = [];
                if (row.mailAddress !== undefined && row.mailAddress != null && row.mailAddress.length > 0) {
                    let addrs = row.mailAddress.split(";");
                    for (var idx in addrs) {
                        this.mailList.push({'addr': addrs[idx]});
                    }
                }
            },
            openEndpointDialog(row) {
                this.showEndpointDialog = true;
                this.currentServiceName = row.serviceName;
                this.currentServiceId = row.serviceId;
                this.currentMailAddr = row.mailAddress;

                this.endpointList = [];
                for (let ep in row.endpointList) {
                    this.endpointList.push({'endpoint': row.endpointList[ep]});
                }
            },
            delMailAddr(row) {
                NProgress.start();
                remove(this.mailList, row);

                let mailAddrs = [];
                for (var idx in this.mailList) {
                    mailAddrs.push(this.mailList[idx].addr);
                }

                let _this = this;
                let param = {
                    serviceName: this.currentServiceName,
                    serviceId: this.currentServiceId,
                    endpointList: this.currentEndpoint,
                    mailAddress: mailAddrs.join(';')
                };
                util.post('/updateMailSub', param, () => {
                    NProgress.done();
                    _this.$message('更新邮箱成功');
                });
            },
            addMailAddr() {
                if (!this.mailAddr.endsWith("@iboxpay.com")) {
                  this.$message('只能使用 iboxpay 邮箱');
                  return;
                }
                NProgress.start();
                this.mailList.push({'addr': this.mailAddr});
                let mailAddrs = [];
                for (let idx in this.mailList) {
                    mailAddrs.push(this.mailList[idx].addr);
                }

                let _this = this;
                let param = {
                    serviceName: this.currentServiceName,
                    serviceId: this.currentServiceId,
                    endpointList: this.currentEndpoint,
                    mailAddress: mailAddrs.join(';')
                };
                util.post('/updateMailSub', param, () => {
                    NProgress.done();
                    _this.$message('更新邮箱成功');
                    _this.mailAddr = '';
                });
            },
            addEndpoint() {
                NProgress.start();
                this.endpointList.push({'endpoint': this.endpoint});

                let endpoints = [];
                for (let idx in this.endpointList) {
                    endpoints.push(this.endpointList[idx].endpoint);
                }

                let _this = this;
                let param = {
                    serviceName: this.currentServiceName,
                    serviceId: this.currentServiceId,
                    mailAddress: this.currentMailAddr,
                    endpointList: endpoints
                };
                util.post('/updateMailSub', param, () => {
                    NProgress.done();
                    _this.$message('更新接口地址成功');
                    _this.endpoint = '';
                });
            },
            delEndpoint(row) {
                NProgress.start();
                remove(this.endpointList, row);

                let endpoints = [];
                for (let idx in this.endpointList) {
                    endpoints.push(this.endpointList[idx].endpoint);
                }

                let _this = this;
                let param = {
                    serviceName: this.currentServiceName,
                    serviceId: this.currentServiceId,
                    mailAddress: this.currentMailAddr,
                    endpointList: endpoints
                };
                util.post('/updateMailSub', param, () => {
                    NProgress.done();
                    _this.$message('更新接口地址成功');
                });
            }
        },
        computed: {

        },
        mounted() {
            NProgress.start();
            let _this = this;
            util.get('/serviceNameList', '', (result) => {
                _this.serviceNameList = result;
            });

            util.get('/mailSubList', {serviceName: this.queryModel.serviceName}, (result) => {
                NProgress.done();
                _this.mailSubList = result;
                for (let idx in _this.mailSubList) {
                    if (_this.mailSubList[idx].endpointList != null) {
                        _this.mailSubList[idx].endpoints = _this.mailSubList[idx].endpointList.join(",")
                    }
                }
            });
        }
    }

    function remove(arr, val) {
        let index = arr.indexOf(val);
        if (index > -1) {
            arr.splice(index, 1);
        }
    }
</script>
<style scoped>
    .toolbar {
        background: #fff;
        padding: 10px 10px 0px 10px;
    }
</style>