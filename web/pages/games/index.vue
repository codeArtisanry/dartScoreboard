<template>
  <div></div>
</template>
<script>
export default {
  middleware: 'notauth',
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
