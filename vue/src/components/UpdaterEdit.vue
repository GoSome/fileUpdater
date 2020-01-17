<template>
  <div class="">
    <p>Editing updater {{ $route.query.name }}</p>
    <div class="my-3" v-if="initialized">
      <editor v-model="content"
              @init="editorInit"
              :lang="lang"
              class="rounded"
              theme="dracula"
              height="500"></editor>
    </div>
    <b-button variant="primary" @click="save">Save</b-button>
    <b-button class="ml-2" @click="cancel">Cancel</b-button>
    <b-modal id="updated-modal"
             ok-title="Got It"
             hide-header
             ok-only>
      <p class="my-4">Update successfully!</p>
    </b-modal>
  </div>
</template>

<script>
export default {
  components: {
    editor: require('vue2-ace-editor'),
  },
  data() {
    return {
      initialized: false,
      content: '',
      lang: 'text',
    }
  },
  created() {
    this.refresh()
  },
  methods: {
    editorInit() {
      require('brace/ext/language_tools')
      require('brace/mode/json')
      require('brace/theme/dracula')
    },
    save() {
      this.$http.post('/api/content', {
        name: this.$route.query.name,
        content: this.content,
      }).then(() => {
        this.refresh()
        this.$bvModal.show('updated-modal')
      }, r => {
        console.log(r)
        alert('Networking Error')
      })
    },
    refresh() {
      this.$http.get('/api/content?name=' + this.$route.query.name).then(r => {
        this.content = r.data.content
        this.initialized = true
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
