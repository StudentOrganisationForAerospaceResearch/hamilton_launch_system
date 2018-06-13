// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

import 'vuetify/dist/vuetify.min.css'
import 'vue-material-design-icons/styles.css'
import VueNativeSock from 'vue-native-websocket'

Vue.use(VueNativeSock, `ws://${document.location.hostname}:8000/ws`, {
  store: store,
  format: 'json',
  reconnection: true, // (Boolean) whether to reconnect automatically (false)
  reconnectionDelay: 3000 // (Number) how long to initially wait before attempting a new (1000)
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
