<template>
  <div>
    <br />
    <div class="alert alert-danger">
      * You want to score 0 touch on outer ring
    </div>
    <div class="container text-center mt-4">
      <div class="bg-white rounded">
        <!-- dart-board -->
        <div class="text-center">
          <div
            id="dartboard"
            class="d-inline-block"
            style="width: 22rem; height: 22rem"
          ></div>
        </div>
        <!-- /*  class="video-btn vidbtn"-->
        <div
          class="form-text text-muted ml-4 my-4 text-uppercase font-weight-bold"
        >
          <ins>Rules for Use Dartboard</ins>
          <a v-b-modal.modalPopover class="bg-white border-0 mt-4 mb-4"
            ><ins> Read...</ins>
          </a>
          <b-modal id="modalPopover" title="Rules for Use Dartboard" ok-only>
            <div class="col sm-12">
              <img
                class="img-responsive"
                width="100%"
                src="/scoreboard.png"
                alt="scoreboard-icon"
              />
            </div>
          </b-modal>
        </div>
        <div v-if="hideUndo == true">
          <button class="btn btn-secondary" @click="undoLastScore">Undo</button>
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
                  <ul class="list-inline">
                    <li
                      v-for="(dart, index) in p.throws_score"
                      :key="index"
                      class="list-inline-item"
                    >
                      <mark
                        v-if="p.check_round == 'INVALID'"
                        style="background-color: #ffcccb"
                      >
                        {{ dart }}
                      </mark>
                      <div v-else>
                        {{ dart }}
                      </div>
                    </li>
                  </ul>
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
      hideUndo: Boolean,
    };
  },
  computed: {
    getPlayerInfo() {
      return this.$store.state.game.playerInfo;
    },
    getScoreboard() {
      return this.$store.state.game.scoreboard;
    },
    getTurn() {
      return this.$store.state.game.currentTurn;
    },
  },
  async mounted() {
    await this.currentTurnApi();
    await this.playerInfoApi();
    const scoreSpeak = new SpeechSynthesisUtterance(
      "welcome to " + this.getPlayerInfo.game_type + " game"
    );
    window.speechSynthesis.speak(scoreSpeak);
    this.speakPlayerName();
    await this.scoreboardApi();
    this.changeScoreColHeader();
    this.fetchUpdatedData();
    // eslint-disable-next-line no-undef
    const dartboard = new Dartboard("#dartboard");
    dartboard.render();
    document
      .querySelector("#dartboard")
      .addEventListener("throw", async (d) => {
        await this.postScoreApi(d.detail);
        const scoreSpeak = new SpeechSynthesisUtterance(d.detail.score);
        window.speechSynthesis.speak(scoreSpeak);
        await this.currentTurnApi();
        await this.playerInfoApi();
        await this.scoreboardApi();
        if (this.getTurn.throw === 1) {
          this.speakPlayerName();
        }
        this.fetchUpdatedData();
        this.hideUndo = true;
      });
  },
  methods: {
    async playerInfoApi() {
      await this.$store.dispatch("game/getGamePlayerInfo", {
        gameId: this.$route.params.gameid,
        playerId: this.getTurn.player_id,
      });
    },
    async postScoreApi(dartScore) {
      await this.$store.dispatch("game/postScore", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
        roundId: this.$route.params.roundid,
        turnId: this.$route.params.turnid,
        score: dartScore,
      });
    },
    async scoreboardApi() {
      await this.$store.dispatch(
        "game/getScoreboard",
        this.$route.params.gameid
      );
    },
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
    async undoLastScore() {
      await this.$store.dispatch("game/undoScore", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
        roundId: this.$route.params.roundid,
        turnId: this.$route.params.turnid,
      });
      await this.currentTurnApi();
      await this.playerInfoApi();
      await this.scoreboardApi();
      this.fetchUpdatedData();
    },
    speakPlayerName() {
      const playerNameSpeak = new SpeechSynthesisUtterance(
        "Now " + this.getPlayerInfo.active_player_info.first_name + "'s Turn"
      );
      window.speechSynthesis.speak(playerNameSpeak);
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
