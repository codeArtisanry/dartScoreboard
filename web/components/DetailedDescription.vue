<template>
  <div class="">
    <div class="container my-5 mt-5 mb-3 text-center">
      <h2 class="font-weight-bolder">Game Info</h2>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Type</th>
            <th scope="col">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="gameData in registerGame"
            :key="gameData.id"
            @click="clicked(gameData.id)"
          >
            <th scope="row">{{ gameData.gameName }}</th>
            <td>{{ gameData.gameType }}</td>
            <td>{{ gameData.gameStatus }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      registerGame: '',
    }
  },
  created() {
    this.getGameData()
  },
  methods: {
    clicked(id) {
      this.$router.push('/games/' + id + '/verify')
    },
    async getGameData() {
      const res = await this.$axios.$get(`${process.env.base_URL}/registerGame`)
      this.registerGame = res
    },
  },
}
</script>
