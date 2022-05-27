import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

// 导入 element-ui
import elementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(elementUI)

import axios from "axios"
Vue.prototype.$axios=axios

import qs from 'qs'
Vue.prototype.$qs=qs

require('./mock/LoginMock.js')

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
