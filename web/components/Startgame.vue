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
              <td scope="col">{{ round }}</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">Player Name</th>
              <td scope="col">{{ name }}</td>
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
              v-model="singlescore"
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
    registerGame: Object,
  },
  data() {
    return {
      count: 0,
      round: 1,
      counter: 0,
      name: this.$store.state.players[0].content,
      singlescore: 0,
      total: 0,
    }
  },
  methods: {
    increment() {
      this.count++
      if (this.count === 9 * this.$store.state.players.length) {
        this.$router.push('/highestscoreboard')
      }
      if (this.count % (3 * this.$store.state.players.length) === 0) {
        this.round = this.round + 1
      }
      if (this.count % 3 === 0) {
        this.counter = this.counter + 1
        this.counter = this.counter % this.$store.state.players.length
        this.name = this.$store.state.players[this.counter].content
      }
      this.total = Number(this.total) + Number(this.singlescore)
    },
  },
}
</script>