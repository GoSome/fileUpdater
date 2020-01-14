<template>
  <div class="">
    <p>Editing updater {{ $route.query.name }}</p>
    <b-textarea v-model="content" rows="20" class="mb-2"/>
    <b-button variant="primary" @click="save">Save</b-button>
    <b-button class="ml-2" @click="cancel">Cancel</b-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      content: '',
    }
  },
  created() {
    this.refresh()
  },
  methods: {
    save() {
      this.$http.post('/api/content', {
        name: this.$route.query.name,
        content: this.content
      }).then(() => {
        this.refresh()
      }, r => {
        console.log(r)
        alert('Networking Error')
      })
    },
    refresh() {
      this.$http.get('/api/content?name=' + this.$route.query.name).then(r => {
        this.content = r.data
      }, r => {
        console.log(r)
        alert('Networking error')
      })
    },
    cancel() {
      if (confirm('Are you sure?')) {
        this.$router.back()
      }
    }
  }
}
</script>
