<script>
export default {
  middleware: 'notauth',

  created() {
    this.checkplayerid()
  },
  methods: {
    async checkplayerid() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/active-status`
      )
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
    },
  },
}
</script>
