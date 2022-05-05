<template>
  <div>
    <LoadingSpinner v-if="showHideSpinner" />

    <Nuxt />
  </div>
</template>


<script>
export default {
  data() {
    return {
      showHideSpinner: true,
    }
  },
  beforeCreate() {
    this.showHideSpinner = true
  },
  mounted() {
    this.showHideSpinner = false
  },
}
</script>
