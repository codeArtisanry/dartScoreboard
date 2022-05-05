<template>
  <div>
    <CommonNavBar />
    <div class="container text-center mt-4">
      <div class="bg-white pb-4 px-5 rounded">
        <h5 class="heading">Welcome to</h5>
        <h4 class="font-weight-bolder">{{ currentgame.game_type }}</h4>
        <table class="table table-striped shadow mt-3">
          <tbody>
            <tr>
              <th class="text-left" scope="row">Round</th>
              <td scope="col">{{ currentgame.round }}</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">Player Name</th>
              <td scope="col">
                {{ playername }}
              </td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">{{ gameScore }}</th>
              <td>{{ scoreofplayer }}</td>
            </tr>
          </tbody>
        </table>
        <form class="pt-3">
          <div class="form-group">
            <label for="enterThrow" class="text-muted font-weight-bolder"
              >Enter Points in Dart Throw :{{ currentgame.throw }}</label
            >
            <input
              id="enterThrow"
              v-model.number="throwscore"
              type="number"
              class="form-control"
              min="0"
              autofocus
              @keydown="noSpecialchar($event)"
              @keypress="onlyNumber($event)"
            />
          </div>
          <button
            type="button"
            class="btn btn-success"
            @click="postgamescore"
            @keyup.enter="postgamescore"
          >
            Submit
          </button>
        </form>
      </div>
      <div>
        <!-- Using value -->
        <b-button v-b-toggle="'collapse-2'" class="m-1" @click="getCurrentGame"
          >To see ScoreBoard Click Here</b-button
        >
        <!-- Element to collapse -->
        <b-collapse id="collapse-2">
          <b-card>
            <table class="table">
              <thead>
                <tr>
                  <th scope="col">Name</th>
                  <th
                    v-for="throwscore in rounddata"
                    :key="throwscore"
                    scope="col"
                  >
                    Round {{ throwscore.round }}
                  </th>
                  <th scope="col">{{ gameScore }}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="player in players" :key="player">
                  <th scope="row">
                    {{ player.first_name + '  ' + player.last_name }}
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
          </b-card>
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
      scoreboard: '',
      dart: {
        score: 0,
      },
      throwscore: '',
      checkplayer: '',
      scoreofplayer: 0,
      playername: '',
      registerGame: '',
      gameScore: '',
      currentgame: '',
      players_score: [],
    }
  },
  async created() {
    await this.checkplayerid()
    await this.getCurrentGame()
    if (
      this.currentgame.game_type === 'Target Score-101' ||
      this.currentgame.game_type === 'Target Score-301' ||
      this.currentgame.game_type === 'Target Score-501'
    ) {
      this.gameScore = 'Remaining Score'
    } else {
      this.gameScore = 'Score'
    }
    this.players = this.scoreboard.players_score
    this.rounddata = this.scoreboard.players_score[0].rounds
    this.playername =
      this.currentgame.active_player_info.first_name +
      ' ' +
      this.currentgame.active_player_info.last_name
  },
  methods: {
    async getCurrentGame() {
      const res = await this.$axios.$get(
        `api/v1/games/` +
          this.$route.params.gameid +
          `/players/` +
          this.$route.params.playerid +
          `/player-info`
      )
      this.currentgame = res
      const res1 = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/scoreboard`
      )
      this.scoreboard = res1
      for (let i = 0; i <= this.scoreboard.players_score.length - 1; i++) {
        if (
          this.currentgame.active_player_info.first_name ===
          this.scoreboard.players_score[i].first_name
        ) {
          this.scoreofplayer = this.scoreboard.players_score[i].total
        }
      }
    },
    async checkplayerid() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/active-status`
      )
      this.checkplayer = res
      if (res.player_id === 0) {
        this.$router.push(`/games/` + this.$route.params.gameid + `/scoreboard`)
      } else {
        const res = await this.$axios.$get(
          `api/v1/games/` + this.$route.params.gameid + `/active-status`
        )
        this.checkplayer = res
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
    onlyNumber($event) {
      const keyCode = $event.keyCode ? $event.keyCode : $event.which
      if (keyCode < 0 || keyCode > 60) {
        $event.preventDefault()
      }
    },
    noSpecialchar(e) {
      if (/^\W$/.test(e.key)) {
        e.preventDefault()
      }
    },
    async postgamescore() {
      this.dart.score = Number(this.throwscore)
      if (
        this.dart.score < 0 ||
        this.dart.score > 60 ||
        !Number.isInteger(this.dart.score)
      ) {
        alert('please enter valid score')
      } else {
        await this.$axios.$post(
          `api/v1/games/` + this.$route.params.gameid + `/score`,
          this.dart
        )
        console.log(this.dart, 'responce of post')
        this.$router.push(`/games/` + this.checkplayer.game_id + `/player/`)
      }
    },
  },
}
</script>
