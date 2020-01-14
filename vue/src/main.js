import Vue from 'vue'
import VueResource from 'vue-resource'
import BootstrapVue from 'bootstrap-vue'

import App from './App'
import Block from './components/Block'
import router from './router'

Vue.router = router // Just for VueResource

router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? to.meta.title :
      (to.name ? to.name.charAt(0).toUpperCase() + to.name.slice(1) : 'Untitled')
  next()
})

// use 顺序有要求
Vue.use(VueResource)
Vue.use(BootstrapVue)

Vue.component('block', Block)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
