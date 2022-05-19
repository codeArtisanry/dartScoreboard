<template>
  <client-only>
    <div class="text-center mt-5 mb-5">
      <div class="container text-center">
        <div class="text-center">
          <h4>
            Congratulations
            <b
              ><u>{{ getScoreboardStore.winner }}</u></b
            >
          </h4>
          <h5>You Win This Game</h5>
          <br />

          <h3 class="text-center">{{ getGame.game_name }}</h3>
          <h4 class="text-center">{{ getGame.game_type }}</h4>
          <br />
          <h5 class="text-center">ScoreBoard</h5>
        </div>
        <div class="text-center row">
          <div class="table-responsive col-sm-12">
            <table class="table-hover table showtable">
              <tr>
                <td colspan="2">
                  <b><u>Name</u></b>
                </td>
                <td>
                  <b><u>Total Score</u></b>
                </td>
              </tr>
              <tbody
                v-for="(player, i) in getScoreboardStore.players_score"
                :key="i"
                v-b-toggle="`${i}collapse`"
              >
                <tr>
                  <td>
                    <div class="d-flex justify-content-center">
                      <div>
                        <img
                          height="10"
                          src="/downarrow.svg"
                          alt="click here"
                          class="text-white mb-1 mr-2"
                        />
                      </div>
                    </div>
                  </td>
                  <td colspan="1">
                    {{ player.first_name + " " + player.last_name }}
                  </td>
                  <td v-if="getGame.game_type == 'High Score'">
                    {{ player.total }}
                  </td>
                  <td v-else-if="getGame.game_type == 'Target Score-101'">
                    {{ 101 - player.total }}
                  </td>
                  <td v-else-if="getGame.game_type == 'Target Score-301'">
                    {{ 301 - player.total }}
                  </td>
                  <td v-else>{{ 501 - player.total }}</td>
                </tr>
                <b-collapse :id="`${i}collapse`">
                  <tr>
                    <td>Round</td>
                    <td>Darts</td>
                    <td>Total</td>
                  </tr>
                  <tr v-for="(round, j) in player.rounds" :key="j" scope="row">
                    <td scope="row">
                      <mark
                        v-if="round.check_round == 'INVALID'"
                        style="background-color: #ffcccb"
                      >
                        {{ round.round }}
                      </mark>
                      <div v-else>{{ round.round }}</div>
                    </td>
                    <td>
                      <mark
                        v-if="round.check_round == 'INVALID'"
                        style="background-color: #ffcccb"
                        >{{ round.throws_score }}</mark
                      >
                      <div v-else>{{ round.throws_score }}</div>
                    </td>
                    <td>
                      <mark
                        v-if="round.check_round == 'INVALID'"
                        style="background-color: #ffcccb"
                        >{{ round.round_total }}</mark
                      >
                      <div v-else>{{ round.round_total }}</div>
                    </td>
                  </tr>
                </b-collapse>
              </tbody>
            </table>
            <hr />
          </div>
        </div>
        <div class="text-center"></div>
        <br />
        <div class="d-grid gap-2 col-6 mx-auto">
          <button class="btn btn-secondary" type="button" @click="homepage">
            Home Page
          </button>
        </div>
      </div>
    </div>
  </client-only>
</template>
<script>
export default {
  computed: {
    getGame() {
      return this.$store.state.game.game;
    },
    getScoreboardStore() {
      return this.$store.state.game.scoreboard;
    },
  },
  async created() {
    await this.gameApi();
    await this.scoreboardApi();
  },
  methods: {
    homepage() {
      this.$router.push("/");
    },
    async gameApi() {
      await this.$store.dispatch("game/getGame", this.$route.params.gameid);
    },
    async scoreboardApi() {
      await this.$store.dispatch(
        "game/getScoreboard",
        this.$route.params.gameid
      );
    },
  },
};
</script>

<style scoped>
.showtable .show {
  display: contents !important;
}
</style>
