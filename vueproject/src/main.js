import Vue from 'vue'
import App from './App.vue'
import 'element-ui/lib/theme-chalk/index.css'
import element from './element/index'
import axios from 'axios'
import router from './router'




Vue.use(element)

Vue.config.productionTip = false
Vue.prototype.$ajax=axios
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

