<template>
  <div></div>
</template>
<script>
export default {
  computed: {
    getTurn() {
      return this.$store.state.game.currentTurn;
    },
  },
  async created() {
    await this.currentTurnApi();
  },
  methods: {
    async currentTurnApi() {
      await this.$store.dispatch(
        "game/getCurrentTurn",
        this.$route.params.gameid
      );
      this.scoreboard();
    },
    scoreboard() {
      if (this.getTurn.round === 0) {
        this.$router.push(
          "/games/" + this.$route.params.gameid + "/scoreboard"
        );
      } else {
        this.$router.push(
          `/games/${this.$route.params.gameid}/player/${this.getTurn.player_id}/round/${this.getTurn.round}/turns/${this.getTurn.throw}`
        );
      }
    },
  },
};
</script>
