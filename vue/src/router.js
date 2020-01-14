import Vue from 'vue'
import Router from 'vue-router'
import Layout from './components/Layout.vue'
import Dashboard from './components/Dashboard.vue'

Vue.use(Router)

export default new Router({
  // @see https://router.vuejs.org/en/essentials/history-mode.html
  mode: 'history',
  base: process.env.BASE_URL,
  scrollBehavior: () => ({
    // Scroll logic when jumping between routes
    // @see https://router.vuejs.org/en/advanced/scroll-behavior.html
    y: 0
  }),
  routes: [
    {
      path: '',
      component: Layout,
      meta: {
        auth: true
      },
      children: [
        {
          path: '/',
          name: 'dashboard',
          component: Dashboard,
          meta: {
            auth: true,
            title: '首页',
          },
        },
      ],
    },
    {
      path: '/api*',
      component: () => import('./components/RenderedByBackendPage.vue'),
      meta: {
        auth: true
      }
    },
    {
      path: '*',
      component: () => import('./components/error/404.vue'),
      meta: {
        auth: true,
        title: '页面未找到',
      },
    },
  ]
})
