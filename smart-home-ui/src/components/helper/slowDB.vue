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
    <div>
        <codemirror v-model="code" :options="cmOptions" />
        <el-button style="margin-top: 15px" type="primary" @click="save">保存</el-button>
    </div>
</template>
<script>
    import 'codemirror/mode/yaml/yaml.js'
    import util from '../../common/util'

    import 'codemirror/lib/codemirror.css'
    import 'codemirror/theme/base16-light.css'
    import '../../common/codemirror.css'

    export default {
        data () {
            return {
                code: '',
                cmOptions: {
                    tabSize: 2,
                    mode: 'text/yaml',
                    theme: 'base16-light',
                    lineNumbers: true,
                    line: true
                }
            }
        },
        methods: {
            save() {
                let value = this.code;

                let arr = [];
                let line = value.split(/\r?\n/);
                for (let idx in line) {
                    let pos = line[idx].indexOf("\t");
                    if (pos !== -1) {
                        arr.push("第" + (parseInt(idx) + 1) + "行, 第" + pos + "列, 包含tab字符");
                    }
                }
                if (arr.length > 0) {
                    this.$notify.error({
                        title: '格式有误',
                        dangerouslyUseHTMLString: true,
                        message: arr.join("<br/>")
                    });
                    return;
                }

                let _this = this;
                var settings = {
                    method: 'POST',
                    url: '/update_config/slow',
                    data: value,
                    emulateJSON: true,
                    contentType: 'text/plain; charset=utf-8'
                };
                console.log(settings);
                util.http(settings, result => {
                    _this.$message({
                        message: "保存成功",
                        type: 'success',
                        showClose: true,
                        onClose: function() {
                            location.reload();
                        }
                    });
                });
            }
        },
        computed: {

        },
        mounted() {
            util.get('/get_config/slow', '', result => {
                this.code = result;
            });
        }
    }
</script>
<style scoped>

</style>