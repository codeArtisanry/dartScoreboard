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
            v-for="game in registerGame.game_responses"
            :key="game.id"
            @click="clicked(game.id)"
          >
            <td>{{ game.game_name }}</td>
            <td>{{ game.game_type }}</td>
            <td>{{ game.game_status }}</td>
          </tr>
        </tbody>
      </table>
      <button
        v-if="registerGame.pre_page_link == 'cross limits'"
        variant="outline-primary"
        disabled
        @click="prepage"
      >
        Back
      </button>
      <button v-else variant="outline-primary" @click="prepage">Back</button>
      <button
        v-if="registerGame.post_page_link == 'cross limits'"
        variant="outline-primary"
        disabled
        @click="postpage"
      >
        Next
      </button>
      <button v-else variant="outline-primary" @click="postpage">Next</button>
    </div>
  </div>
</template>
<script>
/* eslint-disable no-console */
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
      this.$router.push('/games/' + id)
    },
    async getGameData() {
      const res = await this.$axios.$get(`/api/v1/games?page=1`)
      this.registerGame = res
      console.log(this.registerGame)
    },
    async prepage() {
      const res = await this.$axios.$get(this.registerGame.pre_page_link)
      this.registerGame = res
      console.log(this.registerGame)
    },
    async postpage() {
      const res = await this.$axios.$get(this.registerGame.post_page_link)
      this.registerGame = res
      console.log(this.registerGame)
    },
  },
}
</script>
