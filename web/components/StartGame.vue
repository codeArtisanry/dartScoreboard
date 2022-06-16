<template>
  <div>
    <Speak :text="speak" />
    <div>
      <NavBar />
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
          <div
            class="
              form-text
              text-muted
              ml-4
              my-4
              text-uppercase
              font-weight-bold
            "
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
          <div v-if="undo">
            <button class="btn btn-secondary" @click="undoScore">Undo</button>
          </div>
          <table class="table table-striped shadow mt-3">
            <tbody>
              <tr>
                <th class="text-center" scope="row">Player Name</th>
                <td scope="col">
                  {{ playerInfo.name }}
                </td>
              </tr>
            </tbody>
            <tbody>
              <tr>
                <th class="text-center" scope="row">Round, Turn</th>
                <td scope="col">
                  {{ playerInfo.round }}, {{ playerInfo.throw }}
                </td>
              </tr>
            </tbody>
            <tbody>
              <tr>
                <th class="text-center" scope="row">{{ scoreColName }}</th>
                <td>{{ playerInfo.score }}</td>
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
                    v-for="roundInfo in rounds"
                    :key="roundInfo.round"
                    scope="col"
                  >
                    R-{{ roundInfo.round }}
                  </th>
                  <th scope="col">{{ scoreColName }}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(player, index) in scoreboard.players" :key="index">
                  <th scope="row">
                    {{ player.first_name + " " + player.last_name[0] + "." }}
                  </th>
                  <td v-for="(round, index) in player.rounds" :key="index">
                    <ul class="list-inline">
                      <li
                        v-for="(dart, index) in round.throws_score"
                        :key="index"
                        class="list-inline-item"
                      >
                        <mark
                          v-if="round.check_round == 'INVALID'"
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
  </div>
</template>
<script>
/* eslint-disable no-console */
export default {
  data() {
    return {
      scoreColName: "",
      undo: true,
      speak: "",
    };
  },
  computed: {
    playerInfo() {
      return this.$store.state.game.player;
    },
    scoreboard() {
      return this.$store.state.game.scoreboard;
    },
    turn() {
      return this.$store.state.game.currentTurn;
    },
    rounds() {
      return this.$store.getters["game/rounds"];
    },
    game() {
      return this.$store.getters["game/info"];
    },
  },

  async mounted() {
    await this.getCurrentTurn();
    await this.getPlayerInfo();
    this.speak = "Welcome To" + this.game.type + "Game";
    await this.getScoreboard();
    this.scoreColumnName();
    this.speak = "Now " + this.playerInfo.name + "'s Turn";

    // eslint-disable-next-line no-undef
    const dartboard = new Dartboard("#dartboard");
    dartboard.render();
    document
      .querySelector("#dartboard")
      .addEventListener("throw", async (d) => {
        await this.addScore(d.detail);
        this.speak = d.detail.score;
        await this.getCurrentTurn();
        await this.getPlayerInfo();
        if (this.turn.throw === 1) {
          this.speak = "Now " + this.playerInfo.name + "'s Turn";
        }
        await this.getScoreboard();
      });
  },

  methods: {
    async getPlayerInfo() {
      await this.$store.dispatch("game/getGamePlayerInfo", {
        gameId: this.$route.params.gameid,
        playerId: this.turn.player_id,
      });
      this.checkGameStatusAndRedirect();
    },

    async addScore(turnScore) {
      await this.$store.dispatch("game/addScore", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
        roundId: this.$route.params.roundid,
        turnId: this.$route.params.turnid,
        score: turnScore,
      });
    },

    async getScoreboard() {
      await this.$store.dispatch(
        "game/getScoreboard",
        this.$route.params.gameid
      );
    },

    async getCurrentTurn() {
      await this.$store.dispatch(
        "game/getCurrentTurn",
        this.$route.params.gameid
      );
    },

    checkGameStatusAndRedirect() {
      if (this.game.status === "Completed") {
        this.$router.push(
          "/games/" + this.$route.params.gameid + "/scoreboard"
        );
      } else {
        this.$router.push(
          `/games/${this.$route.params.gameid}/player/${this.turn.player_id}/round/${this.turn.round}/turns/${this.turn.throw}`
        );
      }
    },

    scoreColumnName() {
      if (this.game.type === "High Score") {
        this.scoreColName = "Total Score";
      } else {
        this.scoreColName = "Remaining Score";
      }
    },

    async undoScore() {
      await this.$store.dispatch("game/undoScore", {
        gameId: this.$route.params.gameid,
        playerId: this.$route.params.playerid,
        roundId: this.$route.params.roundid,
        turnId: this.$route.params.turnid,
      });
      await this.getCurrentTurn();
      await this.getPlayerInfo();
      await this.getScoreboard();
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
