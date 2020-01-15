<template>
  <li class="nav-item">
    <router-link class="nav-link" :to="to">
      <i :class="'fas fa-fw '+ faIcon"></i>
      <span>{{ text }}</span></router-link>
  </li>
</template>

<script>
export default {
  props: {
    to: {
      type: Object,
    },
    faIcon: {
      type: String,
    },
    text: {
      type: String,
    },
  },
  data() {
    return {}
  },
}
</script>
