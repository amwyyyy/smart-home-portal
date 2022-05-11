import Vue from 'vue';
import {MessageBox} from 'element-ui';
import config from '../../config';

/**
 * 对象转参数形式  name=value&age=aaa
 */
function objectEncode(param, key, encode) {
    if (param == null || param === '') {
        return '';
    }
    var paramStr = '';
    var t = typeof(param);
    if (t === 'string' || t === 'number' || t === 'boolean') {
        paramStr += '&' + key + '=' + ((encode == null || encode) ? encodeURIComponent(param) : param);
    } else {
        for (var i in param) {
            var k = key == null ? i : key + (param instanceof Array ? '[' + i + ']' : '.' + i);
            paramStr += objectEncode(param[i], k, encode);
        }
    }
    return paramStr;
}

// 根据是打包还是开发，判断服务器路径
const uri = process.env.NODE_ENV === 'production' ? config.build.baseUri : config.dev.baseUri;

export default {
    post: (url, data, callback) => {
        Vue.http.post(uri + url, data, {emulateJSON: false}).then(
            result => {
                if (result.data.code === 0) {
                    callback(result.data.data);
                } else {
                    MessageBox.alert(result.data.data, '消息', {
                        confirmButtonText: '确定'
                    });
                }
            },
            error => {
                console.error(error);
                MessageBox.alert('服务器异常！', '消息', {
                    confirmButtonText: '确定'
                });
            }
        );
    },
    get: (url, data, callback) => {
        if (url.indexOf('?') === -1) {
            url += '?';
        }
        url += objectEncode(data);

        Vue.http.get(uri + url).then(
            result => {
                if (result.data.code === 0) {
                    callback(result.data.data);
                } else {
                    MessageBox.alert(result.data.data, '消息', {
                        confirmButtonText: '确定'
                    });
                }
            },
            error => {
                console.error(error);
                MessageBox.alert('服务器异常！', '消息', {
                    confirmButtonText: '确定'
                });
            }
        );
    },
    http: (settings, callback) => {
        Vue.http({
            method: settings.method,
            url: uri + settings.url,
            body: settings.data,
            emulateJSON: settings.emulateJSON,
            headers: {
                'Content-Type': settings.contentType,
            }
        }).then(result => {
            if (result.data.code === 0) {
                callback(result.data.data);
            } else {
                MessageBox.alert(result.data.data, '消息', {
                    confirmButtonText: '确定'
                });
            }
        }, error => {
            console.error(error);
            MessageBox.alert('服务器异常！', '消息', {
                confirmButtonText: '确定'
            });
        });
    },
    objectEncode: (param, key, encode) => {
        return objectEncode(param, key, encode);
    },
    getQueryStringByName: (name) => {
        var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
        var r = window.location.search.substr(1).match(reg);
        var context = "";
        if (r !== null) {
            context = r[2];
        }
        reg = null;
        r = null;
        return context === null || context === "" || context === 'undefined' ? '' : context;
    },
    //复制属性到对象
    copyProperties: (target, obj) => {
        target = target || {};
        if (obj) {
            for (var key in obj) {
                if (!obj.hasOwnProperty(key)) {
                    continue;
                }

                var value = obj[key];
                if (typeof(value) === 'undefined') {
                    continue;
                }

                target[key] = value;
            }
        }
    },
    //清空对象属性
    clearProperties: (target) => {
        target = target || {};
        for (var key in target) {
            if (!target.hasOwnProperty(key)) {
                continue;
            }

            var val = target[key];
            if (typeof(val) === 'undefined') {
                continue;
            }

            target[key] = null;
        }
    }
};