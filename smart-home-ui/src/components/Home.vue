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
    <el-container>
        <el-header>
            <el-menu class="el-menu-vertical" mode="horizontal" default-active="/main" unique-opened
                     router background-color="#2C3F53" active-text-color="#ffd04b" text-color="#F7F7F7">
                <el-divider style="margin: 0px 0;"></el-divider>
                <template v-for="(item, index) in $router.options.routes" v-if="!item.hidden">
                    <el-submenu :index="index+''" v-if="!item.leaf && item.children.length > 0">
                        <i :class="item.iconCls"></i>
                        <span slot="title">{{ item.name }}</span>
                        <el-menu-item v-for="child in item.children" :index="child.path" v-if="!child.hidden">
                            <i :class="child.iconCls"></i>
                            <span slot="title">{{ child.name }}</span>
                        </el-menu-item>
                    </el-submenu>
                    <el-menu-item v-if="item.leaf && !item.hidden" :index="item.children[0].path">
                        <i :class="item.children[0].iconCls"></i>
                        <span slot="title">{{ item.children[0].name }}</span>
                    </el-menu-item>
                </template>
            </el-menu>
        </el-header>

        <el-main style="margin-top: 5px">
            <router-view></router-view>
        </el-main>
    </el-container>
</template>

<script>
    export default {
        data() {
            return {
                currentPathName: this.$route.name,
                currentPathNameParent: this.$route.matched[0].name,
            };
        },
        watch: {
            $route(to) {
                //监听路由改变
                this.currentPathName = to.name;
                this.currentPathNameParent = to.matched[0].name;
            }
        }
    };
</script>
<style>
    .el-menu-vertical:not(.el-menu--collapse) {
        width: 100%;
    }

    .el-menu--collapse {
        border-right: 1px solid #475669;
        border-top: 1px solid #475669;
    }

    .el-container {
        height: 100%;
        background-color: #EDEDED;
    }

    .el-aside {
        background-color: #2C3F53;
        width: auto !important;
        overflow: visible;
    }

    .el-col {
        box-sizing: border-box;
    }

    .el-divider--horizontal {
        margin: 0px 0;
    }
</style>