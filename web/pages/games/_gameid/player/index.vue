<template>
  <div></div>
</template>

<script>
export default {
  middleware: 'notauth',

  async created() {
    await this.checkPlayerTurn()
  },
  methods: {
    async checkPlayerTurn() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/active-status`
      )
      if (res.player_id === 0) {
        this.$router.push(`/games/` + res.game_id + `/scoreboard`)
      } else {
        this.$router.push(
          `/games/` +
            res.game_id +
            `/player/` +
            res.player_id +
            `/round/` +
            res.round +
            `/turns/` +
            res.throw
        )
      }
    },
  },
}
</script>
