// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import 'vuetify/dist/vuetify.min.css'
import 'vue-material-design-icons/styles.css'
import ThermometerLinesIcon from 'vue-material-design-icons/thermometer-lines.vue'
import WeatherWindyIcon from 'vue-material-design-icons/weather-windy.vue'
import SignDirectionIcon from 'vue-material-design-icons/sign-direction.vue'
import WaterIcon from 'vue-material-design-icons/water.vue'

Vue.config.productionTip = false

Vue.component('thermometer-lines', ThermometerLinesIcon)
Vue.component('weather-windy', WeatherWindyIcon)
Vue.component('sign-direction', SignDirectionIcon)
Vue.component('water', WaterIcon)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
