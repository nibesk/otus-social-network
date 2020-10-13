import Vue from 'vue';
import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-sugar.css';
import 'assets/css/toast.css'

Vue.use(VueToast, {
  position: 'top-right',
  duration: 5000,
  dismissible: true
});
