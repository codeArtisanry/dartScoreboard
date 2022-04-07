<template>
  <div>
    <div class="container text-center mt-4">
      <div class="bg-white pb-4 px-5 rounded">
        <h5 class="heading">Welcome to</h5>
        <h4 class="font-weight-bolder">{{ registerGame.gameType }}</h4>

        <table class="table table-striped shadow mt-3">
          <tbody>
            <tr>
              <th class="text-left" scope="row">Round</th>
              <td scope="col">{{ points.round }}</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">Player Name</th>
              <td scope="col">{{ points.playername }}</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">{{ gameScore }}</th>
              <td>{{ total }}</td>
            </tr>
          </tbody>
        </table>
        <form class="pt-3">
          <div class="form-group">
            <label for="enterThrow" class="text-muted font-weight-bolder"
              >Enter Points in Dart Throw :{{ points.throw + 1 }}</label
            >
            <input
              id="enterThrow"
              v-model="points.point"
              type="Number"
              class="form-control"
              autofocus
            />
          </div>
          <button type="button" class="btn btn-success" @click="increment">
            Submit
          </button>
        </form>
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
      registerGame: '',
      gameScore: '',
      // p1: this.registerGame.PlayersNames[0],
      count: 0,
      counter: 0,
      total: 0,
      points: {
        round: 1,
        playername: '',
        throw: 0,
        point: 0,
      },
    }
  },
  async created() {
    await this.getGameData()
    if (this.registerGame.gameType === 'Target Score Game') {
      this.gameScore = 'Remaining Score'
    } else {
      this.gameScore = 'Score'
    }
    //
  },
  methods: {
    increment() {
      console.log(this.registerGame.PlayersNames[0])
      // console.log(this.p1);
      if (this.count === 0) {
        this.points.playername = this.registerGame.PlayersNames[0]
      }
      this.count++
      this.points.throw++
      console.log(this.points.throw)
      if (this.count === 9 * this.registerGame.PlayersNames.length) {
        this.$router.push('/highestscoreboard')
      }
      this.postgamedata()
      if (this.count % (3 * this.registerGame.PlayersNames.length) === 0) {
        this.points.round = this.points.round + 1
      }
      if (this.count % 3 === 0) {
        this.counter = this.counter + 1
        this.counter = this.counter % this.registerGame.PlayersNames.length
        this.points.playername = this.registerGame.PlayersNames[this.counter]
        console.log(this.points.playername)
      }
      this.total = Number(this.total) + Number(this.points.point)
      if (this.count % 3 === 0) {
        console.log(this.total)
        this.total = 0
      }
      if (this.count % 3 === 0) {
        this.points.throw = 0
      }
    },
    async postgamedata() {
      await this.$axios.$post(`${process.env.base_URL}/points`, this.points)
    },
    async getGameData() {
      const res = await this.$axios.$get(
        `${process.env.base_URL}/registerGame/` + this.$route.params.gameid
      )
      this.registerGame = res
    },
    playname() {
      // this.playname = this.registerGame.PlayersNames[0]
    },
  },
}
</script>
