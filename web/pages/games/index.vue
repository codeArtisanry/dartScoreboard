<template>
  <div></div>
</template>
<script>
export default {
  created() {
    this.goToHome()
  },
  methods: {
    goToHome() {
      this.$router.push('/')
    },
  },
}
</script>
