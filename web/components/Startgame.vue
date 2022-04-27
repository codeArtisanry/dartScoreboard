<template>
  <div>
    <div class="container text-center mt-4">
      <div class="bg-white pb-4 px-5 rounded">
        <h5 class="heading">Welcome to</h5>
        <h4 class="font-weight-bolder">{{ currentgame.game_type }}</h4>
        {{ rounddata }}
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
              v-model="throwscore"
              type="Number"
              class="form-control"
              autofocus
            />
          </div>
          <button type="button" class="btn btn-success" @click="postgamescore">
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
                  <th scope="col">total</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="player in players" :key="player">
                  <th scope="row">
                    {{ player.first_name + '  ' + player.last_name }}
                  </th>
                  <td v-for="p in player.rounds" :key="p">
                    {{ p.throws_score }}
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
export default {
  props: {
    gameType: Object,
    // registerGame: Object,
  },
  data() {
    return {
      scoreboardobj: '',
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
    this.players = this.scoreboardobj.players_score
    this.rounddata = this.scoreboardobj.players_score[0].rounds
    this.playername =
      this.currentgame.active_player_info.first_name +
      ' ' +
      this.currentgame.active_player_info.last_name
    this.checkplayerid()
  },
  methods: {
    async getCurrentGame() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/current-game`
      )
      this.currentgame = res
      const res1 = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/scoreboard`
      )
      this.scoreboardobj = res1
      console.log(this.scoreboardobj)
      for (let i = 0; i <= this.scoreboardobj.players_score.length - 1; i++) {
        if (
          this.currentgame.active_player_info.first_name ===
          this.scoreboardobj.players_score[i].first_name
        ) {
          this.scoreofplayer = this.scoreboardobj.players_score[i].total
        }
      }
    },
    async checkplayerid() {
      if (this.currentgame.round === 0) {
        // this.$router.push(`/games/` + this.$route.params.gameid + `/scoreboard`)
        console.log('hii  ')
      } else {
        const res = await this.$axios.$get(
          `api/v1/games/` + this.$route.params.gameid + `/active-status`
        )
        this.checkplayer = res
        console.log(this.checkplayer)
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
    async postgamescore() {
      this.dart.score = Number(this.throwscore)
      await this.$axios.$post(
        `api/v1/games/` + this.$route.params.gameid + `/score`,
        this.dart
      )
      console.log(this.dart, 'responce of post')
      this.$router.push(
        `/games/` +
          this.checkplayer.game_id +
          `/player/` +
          this.checkplayer.player_id
      )
    },
  },
}
</script>
