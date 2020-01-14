<template>
  <div class="p-5">
    <p>Updater {{ $route.query.name }}</p>
    <template v-if="updater">
      <p>name: {{ updater.name }}</p>
      <p>path: {{ updater.path }}</p>
      <b-button>Edit</b-button>
    </template>
  </div>
</template>

<script>
export default {
  data() {
    return {
      updater: null,
    }
  },
  created() {
    this.$http.get('/api/updater?name=' + this.$route.query.name).then(r => {
      this.updater = r.data
    }, r => {
      console.log(r)
      alert('Networking error')
    })
  }
}
</script>
