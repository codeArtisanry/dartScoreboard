<template>
  <div>
    <div class="container text-center mt-4">
      <div class="bg-white pb-4 px-5 rounded">
        <h5 class="heading">Welcome to</h5>
        <h4 class="font-weight-bolder">{{ gameType.type }}</h4>

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
              <th class="text-left" scope="row">{{ gameType.score }}</th>
              <td>{{ total }}</td>
            </tr>
          </tbody>
        </table>
        <form class="pt-3">
          <div class="form-group">
            <label for="enterThrow" class="text-muted font-weight-bolder"
              >Enter Points in Dart Throw :</label
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
      count: 0,
      counter: 0,
      total: 0,
      points: {
        round: 1,
        playername: this.$store.state.players[0].content,
        throw: 0,
        point: 0,
      },
    }
  },
  created() {
    this.getGameData()
  },
  methods: {
    increment() {
      this.count++
      this.points.throw++
      // console.log(this.points.throw)
      if (this.count === 9 * this.$store.state.players.length) {
        this.$router.push('/highestscoreboard')
      }
      this.postgamedata()
      if (this.count % (3 * this.$store.state.players.length) === 0) {
        this.points.round = this.points.round + 1
      }
      if (this.count % 3 === 0) {
        this.counter = this.counter + 1
        this.counter = this.counter % this.$store.state.players.length
        this.points.playername = this.$store.state.players[this.counter].content
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
        `${process.env.base_URL}/registerGame/` + this.$route.params.id
      )
      this.registerGame = res
      console.log(this.registerGame)
    },
  },
}
</script>
