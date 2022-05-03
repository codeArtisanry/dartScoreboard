<template>
  <div></div>
</template>

<script>
export default {
  middleware: 'notauth',

  created() {
    this.goToPlayer()
  },
  methods: {
    goToPlayer() {
      this.$router.push(
        `/games/` +
          this.$route.params.gameid +
          `/player/` +
          this.$route.params.playerid
      )
    },
  },
}
</script>
