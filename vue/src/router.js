import Vue from 'vue'
import Router from 'vue-router'
import MainLayout from './components/MainLayout.vue'
import UpdaterView from './components/UpdaterView.vue'
import UpdaterEdit from './components/UpdaterEdit.vue'
import MainPage from './components/MainPage.vue'

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
      component: MainLayout,
      meta: {
        auth: true
      },
      children: [
        {
          path: '/',
          name: 'mainpage',
          component: MainPage,
          meta: {
            title: 'Dashboard',
          },
        },
        {
          path: '/updater',
          name: 'updater',
          component: UpdaterView,
          meta: {
            title: 'Updater',
          },
        },
        {
          path: '/updater/edit',
          name: 'updaterEdit',
          component: UpdaterEdit,
          meta: {
            title: 'Edit',
          },
        },
      ],
    },
    {
      path: '/api*',
      component: () => import('./components/RenderedByBackendPage.vue'),
      meta: {}
    },
    {
      path: '*',
      component: () => import('./components/Error404NotFound.vue'),
      meta: {
        title: '页面未找到',
      },
    },
  ]
})
