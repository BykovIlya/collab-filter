// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
//import 'bootstrap-vue/dist/bootstrap-vue.css
import VueResorce from 'vue-resource'

Vue.use(VueResorce);

Vue.use(BootstrapVue);
Vue.config.productionTip = false;

const API = 'http://localhost:5001';
const API_URL = API + '/api/c/';
Vue.use(API);
//Vue.use(API_URL);

Vue.http.options.root = API_URL;

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
});
