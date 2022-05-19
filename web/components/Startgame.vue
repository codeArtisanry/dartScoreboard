<template>
  <div>
    <script src="https://unpkg.com/dartboard/dist/dartboard.js"></script>
    <div class="container text-center">
      <div class="bg-white rounded">
        <!-- dart-board -->
        <div class="text-center">
          <div
            id="dartboard"
            class="d-inline-block"
            style="width: 22rem; height: 22rem"
          ></div>
        </div>
        <table class="table table-striped shadow mt-3">
          <tbody>
            <tr>
              <th class="text-center" scope="row">Player Name</th>
              <td scope="col">
                {{ playerName }}
              </td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-center" scope="row">Round, Turn</th>
              <td scope="col">
                {{ getPlayerInfo.round }}, {{ getPlayerInfo.throw }}
              </td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-center" scope="row">{{ gameScore }}</th>
              <td>{{ playerScore }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <!-- Using value -->
        <b-button v-b-toggle="'collapse-2'" class="m-1 px-3"
          ><div class="d-flex justify-content-around">
            <div>
              <img
                height="20"
                src="/scoreboard.svg"
                alt="scoreboard-icon"
                class="text-white mb-1"
              />
            </div>
            <div class="ml-2">Scoreboard</div>
          </div></b-button
        >
        <!-- Element to collapse -->
        <b-collapse id="collapse-2" class="table-responsive">
          <table class="table container-fluid wrap">
            <thead>
              <tr>
                <th scope="col">Name</th>
                <th
                  v-for="roundInfo in totalRounds"
                  :key="roundInfo.round"
                  scope="col"
                >
                  R-{{ roundInfo.round }}
                </th>
                <th scope="col">{{ gameScore }}</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="player in getScoreboard.players_score"
                :key="player.first_name"
              >
                <th scope="row">
                  {{ player.first_name + "  " + player.last_name }}
                </th>
                <td v-for="p in player.rounds" :key="p.total">
                  <mark
                    v-if="p.check_round == 'INVALID'"
                    style="background-color: #ffcccb"
                  >
                    {{ p.throws_score }}
                  </mark>
                  <div v-else>
                    {{ p.throws_score }}
                  </div>
                </td>
                <td>{{ player.total }}</td>
              </tr>
            </tbody>
          </table>
        </b-collapse>
      </div>
    </div>
  </div>
</template>
<script>
/* eslint-disable no-console */
export default {
  data() {
    return {
      playerName: "",
      playerScore: "",
      gameScore: "",
      totalRounds: "",
    };
  },
  computed: {
    getPlayerInfo() {
      return this.$store.state.game.playerInfo;
    },
    getScoreboard() {
      return this.$store.state.game.scoreboard;
    },
  },
  async created() {
    await this.playerInfoApi();
    await this.scoreboardApi();
    this.changeScoreColHeader();
    this.fetchUpdatedData();
  },
  mounted() {
    // eslint-disable-next-line no-undef
    const dartboard = new Dartboard("#dartboard");
    dartboard.render();
    document
      .querySelector("#dartboard")
      .addEventListener("throw", async (d) => {
        await this.postScoreApi(d.detail);
        this.$router.push(`/games/${this.$route.params.gameid}/player`);
      });
  },
  methods: {
    async postScoreApi(dartScore) {
      await this.$store.dispatch("game/postScore", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
        roundId: this.$route.params.roundid,
        turnId: this.$route.params.turnid,
        score: dartScore,
      });
    },
    async playerInfoApi() {
      await this.$store.dispatch("game/getGamePlayerInfo", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
      });
    },
    async scoreboardApi() {
      await this.$store.dispatch(
        "game/getScoreboard",
        this.$route.params.gameid
      );
    },
    fetchUpdatedData() {
      this.playerName =
        this.$store.state.game.playerInfo.active_player_info.first_name +
        " " +
        this.$store.state.game.playerInfo.active_player_info.last_name;
      this.playerScore =
        this.$store.state.game.playerInfo.active_player_info.score;
      this.totalRounds =
        this.$store.state.game.scoreboard.players_score[0].rounds;
    },
    changeScoreColHeader() {
      if (this.getPlayerInfo.game_type === "High Score") {
        this.gameScore = "Total Score";
      } else {
        this.gameScore = "Remaining Score";
      }
    },
  },
};
</script>

<style scoped>
.wrap {
  width: 110%;
  max-width: 110%;
  margin-bottom: 25px;
  white-space: nowrap;
}
</style>
