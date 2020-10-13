// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import App from './App';
import router from './router';
import store from './store';
import './plugins';
import {routes} from './config/routes';
import {sync} from 'vuex-router-sync';
import BootstrapVue from "bootstrap-vue";
import InfiniteLoading from 'vue-infinite-loading'
import {ws} from 'api/ws'

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import 'assets/css/global.css'

Vue.use(BootstrapVue);
Vue.use(InfiniteLoading, { /* options */});

Vue.prototype.$routes = routes;

sync(store, router);

Vue.config.productionTip = false;

Vue.use(ws, {
    store,
    url: `ws://${document.location.host}${routes.service_chat.ws}`
});

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App),
    created() {
        this.$store.watch(() => this.$store.getters['user/getUser'], el => {
            if (el === null) {
                this.$ws.disconnect()
            } else {
                this.$ws.connect()
            }
        })
    }
}).$mount('#app');
