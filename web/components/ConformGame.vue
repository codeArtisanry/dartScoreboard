<template>
  <div>
    <center>
      <div class="text-center container mx-2">
        <p class="font-weight-bolder">Game Name:</p>
        <p class="text text-muted">{{ registerGame.gameName }}</p>
        <hr />
        <p class="font-weight-bolder">Game Type:</p>
        <p class="text text-muted">{{ registerGame.gameType }}</p>
        <hr />
        <p class="font-weight-bolder">Players Names:</p>
        <p
          v-for="player in registerGame.PlayersNames"
          :key="player"
          class="text-muted"
        >
          {{ player }}
        </p>
      </div>

      <div class="row">
        <div class="col col-4 text-right">
          <button class="btn btn-info" @click="updateGame">Update</button>
        </div>
        <div class="col col-4 text-center">
          <button class="btn btn-danger">Delete</button>
        </div>
        <div class="col col-4 text-left">
          <button class="btn btn-success" @click="startgame()">Start</button>
        </div>
      </div>
      <div class="text-center mt-5">
        <button class="btn btn-secondary">back to home</button>
      </div>
    </center>
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
    async getGameData() {
      const res = await this.$axios.$get(
        `${process.env.base_URL}/registerGame/` + this.$route.params.id
      )
      this.registerGame = res
    },
    updateGame() {
      this.$router.push('/home/creategame')
    },
    startgame() {
      this.$router.push('/start/highscoregame/' + this.$route.params.id)
    },
  },
}
</script>
