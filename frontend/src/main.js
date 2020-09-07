// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import App from './App';
import router from './router';
import store from './store';
import './plugins';
import {routes} from './router/routes';
import { sync } from 'vuex-router-sync';
import BootstrapVue from "bootstrap-vue";
import InfiniteLoading from 'vue-infinite-loading'

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import 'assets/css/global.css'

Vue.use(BootstrapVue);
Vue.use(InfiniteLoading, { /* options */ });

Vue.prototype.$routes = routes;

sync(store, router);

Vue.config.productionTip = false;

/* eslint-disable no-new */
const vue = new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
}).$mount('#app');
