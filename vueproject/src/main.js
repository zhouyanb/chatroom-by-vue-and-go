import Vue from 'vue'
import App from './App.vue'
import 'element-ui/lib/theme-chalk/index.css'
import element from './element/index'
import axios from 'axios'
import router from './router'
import store from './store'
import cookies from 'vue-cookies'


Vue.use(element)
Vue.config.productionTip = false
Vue.prototype.$cookies=cookies
Vue.prototype.$ajax=axios

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

