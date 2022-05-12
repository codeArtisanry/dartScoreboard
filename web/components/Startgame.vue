<template>
  <div>
    <NavBar />
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
                {{ playername }}
              </td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-center" scope="row">Round, Turn</th>
              <td scope="col">
                {{ currentgame.round }}, {{ currentgame.throw }}
              </td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-center" scope="row">{{ gameScore }}</th>
              <td>{{ scoreofplayer }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <!-- Using value -->
        <b-button
          v-b-toggle="'collapse-2'"
          class="m-1 px-3"
          @click="getCurrentGame"
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
                  v-for="throwscore in rounddata"
                  :key="throwscore"
                  scope="col"
                >
                  R-{{ throwscore.round }}
                </th>
                <th scope="col">{{ gameScore }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="player in players" :key="player">
                <th scope="row">
                  {{ player.first_name + "  " + player.last_name }}
                </th>
                <td v-for="p in player.rounds" :key="p">
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
      rounddata: "",
      players: "",
      scoreboard: "",
      checkplayer: "",
      scoreofplayer: 0,
      playername: "",
      registerGame: "",
      gameScore: "",
      currentgame: "",
      players_score: [],
    };
  },
  async created() {
    await this.checkplayerid();
    await this.getCurrentGame();

    if (
      this.currentgame.game_type === "Target Score-101" ||
      this.currentgame.game_type === "Target Score-301" ||
      this.currentgame.game_type === "Target Score-501"
    ) {
      this.gameScore = "Remaining Score";
    } else {
      this.gameScore = "Score";
    }
    this.players = this.scoreboard.players_score;
    this.rounddata = this.scoreboard.players_score[0].rounds;
    this.playername =
      this.currentgame.active_player_info.first_name +
      " " +
      this.currentgame.active_player_info.last_name;
  },
  mounted() {
    // eslint-disable-next-line no-undef
    const dartboard = new Dartboard("#dartboard");
    console.log(dartboard, "Hello World");
    dartboard.render();
    document.querySelector("#dartboard").addEventListener("throw", (d) => {
      this.$axios.$post(
        `/api/v1/games/` +
          this.$route.params.gameid +
          `/players/` +
          this.$route.params.playerid +
          `/rounds/` +
          this.$route.params.roundid +
          `/turns/` +
          this.$route.params.turnid +
          `/score`,
        d.detail
      );
      this.$router.push(`/games/` + this.checkplayer.game_id + `/player/`);
    });
  },
  methods: {
    async getCurrentGame() {
      const res = await this.$axios.$get(
        `api/v1/games/` +
          this.$route.params.gameid +
          `/players/` +
          this.$route.params.playerid +
          `/player-info`
      );
      this.currentgame = res;
      const res1 = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/scoreboard`
      );
      this.scoreboard = res1;
      for (let i = 0; i <= this.scoreboard.players_score.length - 1; i++) {
        if (
          this.currentgame.active_player_info.first_name ===
          this.scoreboard.players_score[i].first_name
        ) {
          this.scoreofplayer = this.scoreboard.players_score[i].total;
        }
      }
    },
    async checkplayerid() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/active-status`
      );
      this.checkplayer = res;
      if (res.player_id === 0) {
        this.$router.push(
          `/games/` + this.$route.params.gameid + `/scoreboard`
        );
      } else {
        const res = await this.$axios.$get(
          `api/v1/games/` + this.$route.params.gameid + `/active-status`
        );
        this.checkplayer = res;
        this.$router.push(
          `/games/` +
            res.game_id +
            `/player/` +
            res.player_id +
            `/round/` +
            res.round +
            `/turns/` +
            res.throw
        );
      }
    },
    onlyNumber($event) {
      const keyCode = $event.keyCode ? $event.keyCode : $event.which;
      if (keyCode < 0 || keyCode > 60) {
        $event.preventDefault();
      }
    },
    noSpecialchar(e) {
      if (/^\W$/.test(e.key)) {
        e.preventDefault();
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
