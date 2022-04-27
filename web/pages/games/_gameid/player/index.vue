<template>
  <div></div>
</template>

<script>
export default {
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
