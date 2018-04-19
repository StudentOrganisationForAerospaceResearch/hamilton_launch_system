import Vue from 'vue'
import Vuetify from 'vuetify'
import Router from 'vue-router'
import Overview from '@/views/pages/Overview'
import VideoFeed from '@/views/pages/VideoFeed'
import Filling from '@/views/pages/Filling'
import Control from '@/views/pages/Control'

Vue.use(Router)
Vue.use(Vuetify)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/overview'
    },
    {
      path: '/overview',
      name: 'Overview',
      component: Overview
    },
    {
      path: '/video-feed',
      name: 'VideoFeed',
      component: VideoFeed
    },
    {
      path: '/filling',
      name: 'Filling',
      component: Filling
    },
    {
      path: '/control',
      name: 'Control',
      component: Control
    }
  ]
})
