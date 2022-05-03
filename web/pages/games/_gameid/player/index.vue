<template>
  <div></div>
</template>

<script>
export default {
  middleware: 'notauth',

  created() {
    this.goToGame()
  },
  methods: {
    goToGame() {
      this.$router.push(`/games/` + this.$route.params.gameid)
    },
  },
}
</script>
