<template>
  <div class="text-center mt-5">
    <CommonNavBar />
    <div class="container text-center">
      <h3 class="text-center">{{ title }}</h3>
      <h4 class="text-center">{{ subtitle }}</h4>
      <h5 class="text-center">ScoreBoard</h5>
    </div>
    <br /><br />
    <div class="accordion">
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Game Type</th>
          </tr>
        </thead>
        <tbody>
          <tr v-b-toggle.accordion-2 block variant="info">
            <td scope="row">{{ winner }}</td>
            <td>{{ gametype }}</td>
          </tr>
        </tbody>
      </table>
      <b-collapse
        id="accordion-2"
        visible
        accordion="my-accordion"
        role="tabpanel"
      >
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

    <br /><br />

    <div class="text-center">
      <h4>
        Congratulations <b>{{ winner }}</b>
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
  props: {
    title: String,
    subtitle: String,
    score: String,
  },
  data() {
    return {
      winner: ' ',
      gametype: ' ',
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
    this.winner = this.scoreboardobj.winner
    this.gametype = this.currentgame.game_type
  },
  methods: {
    async getCurrentGame() {
      const res = await this.$axios.$get(
        `/api/v1/games/` + this.$route.params.gameid + `/current-game`
      )
      this.currentgame = res
      const res1 = await this.$axios.$get(
        `/api/v1/games/` + this.$route.params.gameid + `/scoreboard`
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
    homepage() {
      this.$router.push('/')
    },
  },
}
</script>
