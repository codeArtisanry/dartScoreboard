<template>
  <div class="text-center mt-5 mb-5">
    <div class="container text-center">
      <h3 class="text-center">{{ gameInfo.game_name }}</h3>
      <h4 class="text-center">{{ gameInfo.game_type }}</h4>
      <h5 class="text-center">ScoreBoard</h5>
    </div>
    <table class="table">
      <thead>
        <tr scope="row">
          <th scope="col">Name</th>
          <th v-for="playerRounds in rounds" :key="playerRounds">
            Round {{ playerRounds.round }}
          </th>
          <th scope="col">{{ totalColHeader }}</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="playerInfo in scoreboard.players_score"
          :key="playerInfo"
          scope="row"
        >
          <th scope="col">
            {{ playerInfo.first_name + ' ' + playerInfo.last_name }}
          </th>
          <td v-for="thows in playerInfo.rounds" :key="thows">
            {{ thows.throws_score }}({{ thows.round_total }})
          </td>
          <th scope="col">
            {{ playerInfo.total }}
          </th>
        </tr>
      </tbody>
    </table>
    <br /><br />
    <div class="text-center">
      <h4>
        Congratulations <b>{{ scoreboard.winner }}</b>
      </h4>
      <h5>You Win This Game</h5>
      <br />
      <div class="d-grid gap-2 col-6 mx-auto">
        <button class="btn btn-success" type="button" @click="homepage">
          Home Page
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      gameInfo: '',
      scoreboard: '',
      rounds: '',
      totalColHeader: '',
    }
  },
  async created() {
    await this.getCurrentGame()
    await this.getScoreboard()
  },
  methods: {
    // call games api for game infomation
    async getCurrentGame() {
      const gameApiRes = await this.$axios.$get(
        `/api/v1/games/` + this.$route.params.gameid
      )
      this.gameInfo = gameApiRes
      this.chanageTotalColHeader()
    },
    // call scoreboard api for perticuler game for get all players scores, total and get winner
    async getScoreboard() {
      const scoreboardApiRes = await this.$axios.$get(
        `/api/v1/games/` + this.$route.params.gameid + `/scoreboard`
      )
      this.scoreboard = scoreboardApiRes
      this.rounds = this.scoreboard.players_score[0].rounds
    },
    // change total column for perticuler game type
    chanageTotalColHeader() {
      if (this.gameInfo.game_type === 'High Score') {
        this.totalColHeader = 'Total Score'
      } else {
        this.totalColHeader = 'Remaining Score'
      }
    },
    homepage() {
      this.$router.push('/')
    },
  },
}
</script>
