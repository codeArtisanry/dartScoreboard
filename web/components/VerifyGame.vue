<template>
  <div>
    <CommonNavBar />
    <center>
      <div class="text-center container mx-2">
        <p class="font-weight-bolder">Game Name:</p>
        <p class="text text-muted">{{ registerGame.game_name }}</p>
        <hr />
        <p class="font-weight-bolder">Game Type:</p>
        <p class="text text-muted">{{ registerGame.game_type }}</p>
        <hr />
        <p class="font-weight-bolder">Players Names:</p>
        <p
          v-for="player in registerGame.players"
          :key="player.id"
          class="text-muted"
        >
          {{ player.first_name }}
        </p>
      </div>

      <div class="row">
        <div class="col col-4 text-right">
          <button class="btn btn-info" @click="updateGame">Update</button>
        </div>
        <div class="col col-4 text-center">
          <button class="btn btn-danger" @click="DeleteGame(id)">Delete</button>
        </div>
        <div class="col col-4 text-left">
          <button class="btn btn-success" @click="startgame()">Start</button>
        </div>
      </div>
      <div class="text-center mt-5">
        <button class="btn btn-secondary" @click="backToHome">
          back to home
        </button>
      </div>
    </center>
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
    async getGameData() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid
      )
      this.registerGame = res
      console.log(this.registerGame)
    },
    updateGame() {
      if (
        this.registerGame.game_status === 'In Progress' ||
        this.registerGame.game_status === 'Completed'
      ) {
        alert('not capable to update')
      } else {
        this.$router.push('/games/' + this.$route.params.gameid + '/update')
      }
    },
    async startgame() {
      const res = await this.$axios.$get(
        `api/v1/games/` + this.$route.params.gameid + `/active-status`
      )
      this.$router.push(
        '/games/' + this.$route.params.gameid + '/player/' + res.player_id
      )
      console.log(res.player_id)
    },
    backToHome() {
      this.$router.push('/')
    },
    async DeleteGame(id) {
      const res = await this.$axios.$delete(
        `api/v1/games` + '/' + this.$route.params.gameid
      )
      console.log(res)
      this.$router.push('/')
    },
  },
}
</script>
