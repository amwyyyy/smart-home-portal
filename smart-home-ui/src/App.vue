<template>
    <div id="app">
        <router-view></router-view>
    </div>
</template>
<script>
    export default {
        name: 'app',
        components: {}
    }
</script>
<style>
    @import "../node_modules/flex.css/dist/flex.css";

    body {
        margin: 0px;
        padding: 0px;
        background: #1F2D3D;
        font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
        font-weight: 400;
        -webkit-font-smoothing: antialiased;
    }

    #app {
        position: absolute;
        top: 0px;
        bottom: 0px;
        width: 100%;
        background-size: cover;
    }

    .toolbar {
        background: #fff;
        padding: 20px 10px 0px 10px;
        margin-bottom: 10px;
    }

    .toolBtn {
        float: right;
    }

    .cell {
        padding-left: 10px !important;
    }
</style>