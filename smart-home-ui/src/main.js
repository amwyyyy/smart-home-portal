import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Resource from 'vue-resource';
import VueRouter from 'vue-router';
import VueCodemirror from 'vue-codemirror'

import App from './App';
import Home from './components/Home.vue';

Vue.use(ElementUI);
Vue.use(VueRouter);
Vue.use(Resource);
Vue.use(VueCodemirror);

var routes = [{
    path: '/',
    redirect: '/main',
    component: Home,
    name: '控制台',
    hidden: true,
    children: [{
        path: '/main',
        name: '控制台',
        component: require('./components/Main.vue')
    }]
}, {
    path: '/alarm',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/alarm',
        name: '告警配置',
        iconCls: 'el-icon-phone-outline',
        component: require('./components/helper/alarm.vue')
    }]
}, {
    path: '/mail',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/mail',
        name: '报表配置',
        iconCls: 'el-icon-message',
        component: require('./components/helper/mail.vue')
    }]
}, {
    path: '/segment',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/segment',
        name: '异常链路管理',
        iconCls: 'el-icon-circle-close',
        component: require('./components/helper/segment.vue')
    }]
}, {
    path: '/apdex',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/apdex',
        name: 'Apdex配置',
        iconCls: 'el-icon-s-data',
        component: require('./components/helper/apdex.vue')
    }]
}, {
    path: '/slowdb',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/slowdb',
        name: '数据库慢查询配置',
        iconCls: 'el-icon-s-platform',
        component: require('./components/helper/slowDB.vue')
    }]
}, {
    path: '/stat',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/stat',
        name: '资源报表',
        iconCls: 'el-icon-data-line',
        component: require('./components/stat/InstanceResource.vue')
    }]
}, {
    path: '/collect',
    component: Home,
    name: '',
    hidden: false,
    leaf: true,
    children: [{
        path: '/collect',
        name: '资源统计',
        iconCls: 'el-icon-pie-chart',
        component: require('./components/stat/InstanceResourceCollect.vue')
    }]
}];

const router = new VueRouter({
    routes
});

new Vue({
    el: '#app',
    template: '<App/>',
    router,
    components: {
        App
    }
}).$mount('#app');