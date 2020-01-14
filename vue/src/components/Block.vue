<template>
  <div class="d-flex flex-column py-2 px-3 bg-white block">
    <div v-if="title" class="pb-1 mb-2 font-weight-bold title">{{ title }}</div>
    <slot></slot>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String
    }
  }
}
</script>

<style lang="scss" scoped>
.block {
  $shadow-depth: 3px;
  border-radius: .5rem;
  box-shadow: $shadow-depth $shadow-depth 9px 2px $gray-400;
  margin-top: 1rem;
  margin-bottom: 1rem;

  .title {
    border-bottom: 1px solid #eee;
  }

  &:first-child {
    margin-top: 0;
  }

  &:last-child {
    margin-bottom: 0;
  }
}
</style>
