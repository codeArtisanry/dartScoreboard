<template>
  <div>
    <Speak :text="speak" />
    <div class="text-center mt-5 mb-5">
      <div class="container text-center">
        <div class="text-center">
          <h4>
            Congratulations
            <b
              ><u data-test="winner-name">{{ scoreboard.winner }}</u></b
            >
          </h4>
          <h5>You Win This Game</h5>
          <br />

          <h3 class="text-center" data-test="game-title">{{ game.name }}</h3>
          <h4 class="text-center" data-test="game-type">{{ game.type }}</h4>
          <br />
          <h5 class="text-center">ScoreBoard</h5>
        </div>
        <div class="text-center row">
          <div class="table-responsive ml-4">
            <table class="table-hover table showtable">
              <tr>
                <td colspan="4">
                  <b><u>Name</u></b>
                </td>
                <td>
                  <b><u>Total Score</u></b>
                </td>
              </tr>
              <tbody
                v-for="(player, i) in scoreboard.players"
                :key="i"
                v-b-toggle="`${i}collapse`"
              >
                <!-- @click="toggleDetails(player)"   @click="toggle(i)"-->
                <tr data-test="player-details">
                  <td colspan="2">
                    <div class="justify-content-center">
                      <img
                        height="10"
                        src="/downarrow.svg"
                        alt="click here"
                        class="text-white mb-1 mr-2"
                      />
                    </div>
                  </td>
                  <td colspan="2">
                    {{ player.first_name + " " + player.last_name }}
                  </td>
                  <td v-if="game.type === 'High Score'">
                    {{ player.total }}
                  </td>
                  <td v-else-if="game.type === 'Target Score-101'">
                    {{ 101 - player.total }}
                  </td>
                  <td v-else-if="game.type === 'Target Score-301'">
                    {{ 301 - player.total }}
                  </td>
                  <td v-else>{{ 501 - player.total }}</td>
                </tr>

                <td
                  colspan="6"
                  style="padding: inherit"
                  data-test="player-score"
                >
                  <b-collapse :id="`${i}collapse`">
                    <table class="table">
                      <tr>
                        <td>Round</td>
                        <td>D-1</td>
                        <td>D-2</td>
                        <td>D-3</td>
                        <td>Total</td>
                      </tr>
                      <tr v-for="(round, j) in player.rounds" :key="j">
                        <td>
                          <mark
                            v-if="round.check_round === 'INVALID'"
                            style="background-color: #ffcccb"
                          >
                            {{ round.round }}
                          </mark>
                          <div v-else>{{ round.round }}</div>
                        </td>
                        <td
                          v-for="(dart, index) in round.throws_score"
                          :key="index"
                        >
                          <mark
                            v-if="round.check_round === 'INVALID'"
                            style="background-color: #ffcccb"
                            >{{ dart }}</mark
                          >
                          <div v-else>{{ dart }}</div>
                        </td>
                        <td v-if="round.throws_score === null" colspan="3">-</td>
                        <td
                          v-else-if="round.throws_score.length === 1"
                          colspan="2"
                        >
                          -
                        </td>
                        <td v-else-if="round.throws_score.length === 2">-</td>
                        <td>
                          <mark
                            v-if="round.check_round === 'INVALID'"
                            style="background-color: #ffcccb"
                            >{{ round.round_total }}</mark
                          >
                          <div v-else>{{ round.round_total }}</div>
                        </td>
                      </tr>
                    </table>
                  </b-collapse>
                </td>
              </tbody>
            </table>
          </div>
        </div>
        <div class="text-center"></div>
        <br />
        <div class="d-grid gap-2 col-6 mx-auto">
          <button
            class="btn btn-secondary"
            type="button"
            @click="redirectToHomePage"
          >
            Home Page
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      speak: "",
    };
  },
  computed: {
    game() {
      return this.$store.getters["game/details"];
    },
    scoreboard() {
      return this.$store.state.game.scoreboard;
    },
  },
  async created() {
    await this.getGame();
    await this.getScoreboard();
  },
  mounted() {
    this.speak =
      "Congratulations " + this.scoreboard.winner + "You Win This Game";
  },
  methods: {
    redirectToHomePage() {
      this.$router.push("/");
    },

    async getGame() {
      await this.$store.dispatch("game/getGame", this.$route.params.gameid);
    },

    async getScoreboard() {
      await this.$store.dispatch(
        "game/getScoreboard",
        this.$route.params.gameid
      );
    },
  },
};
</script>

<style scoped>
.showtable {
  display: contents !important;
}
</style>
