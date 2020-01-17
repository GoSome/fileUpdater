<template>
  <div class="">
    <p>Updater {{ $route.query.name }}</p>
    <template v-if="updater">
      <p>name: {{ updater.name }}</p>
      <p>path: {{ updater.path }}</p>
      <pre class="p-3 bg-dark text-white rounded">{{ content }}</pre>
      <b-button variant="primary"
                @click="$router.push({name: 'updaterEdit', query: {name: $route.query.name}})">
        Edit
      </b-button>
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
  computed: {
    name() {
      return this.$route.query.name
    },
  },
  watch: {
    name() {
      this.refresh() // vue won't call `created` when routing in same route (only query changed)
    }
  },
  created() {
    this.refresh()
  },
  methods: {
    refresh() {
      this.$http.get('/api/updater?name=' + this.$route.query.name).then(r => {
        this.updater = r.data.updater
        this.content = r.data.content
      }, r => {
        console.log(r)
        alert('Networking error')
      })
    }
  }
}
</script>
