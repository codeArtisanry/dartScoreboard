<template>
  <div>
    <NavBar />
    <center>
      <div class="text-center container mx-2">
        <p class="font-weight-bolder">Game Name:</p>
        <p class="text text-muted">{{ getGame.game_name }}</p>
        <hr />
        <p class="font-weight-bolder">Game Type:</p>
        <p class="text text-muted">{{ getGame.game_type }}</p>
        <hr />
        <p class="font-weight-bolder">Creater Name:</p>
        <p class="text text-muted">{{ getGame.creater_name }}</p>
        <hr />
        <p class="font-weight-bolder">Players Names:</p>
        <p
          v-for="player in getGame.players"
          :key="player.id"
          class="text-muted"
        >
          {{ player.first_name }}
        </p>
      </div>
      <div class="row">
        <div
          v-if="
            getGame.game_status == 'Not Started' &&
            getGame.creater_name == $store.state.auth.token.name
          "
          class="col text-center"
        >
          <button class="btn btn-info" @click="updateGame">Update</button>
        </div>
        <div
          v-if="getGame.creater_name == $store.state.auth.token.name"
          class="col text-center"
        >
          <button class="btn btn-danger" @click="DeleteGame(getGame.id)">
            Delete
          </button>
        </div>
        <div class="col text-center">
          <button class="btn btn-success" @click="startgame()">
            {{ button }}
          </button>
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
      button: "",
    };
  },
  computed: {
    getGame() {
      return this.$store.state.game.game;
    },
  },
  async created() {
    await this.gameApi();
    this.changeButton();
  },
  methods: {
    async gameApi() {
      await this.$store.dispatch("game/getGame", this.$route.params.gameid);
    },
    changeButton() {
      if (this.getGame.game_status === "Not Started") {
        this.button = "Start";
      } else if (this.getGame.game_status === "In Progress") {
        this.button = "Resume";
      } else {
        this.button = "Scoreboard";
      }
    },
    updateGame() {
      if (
        this.getGame.game_status === "In Progress" ||
        this.getGame.game_status === "Completed"
      ) {
        alert("not capable to update");
      } else {
        this.$router.push("/games/" + this.$route.params.gameid + "/update");
      }
    },
    startgame() {
      if (this.getGame.game_status === "Completed") {
        this.$router.push(
          "/games/" + this.$route.params.gameid + "/scoreboard/"
        );
      } else {
        this.$router.push("/games/" + this.$route.params.gameid + "/player/");
      }
    },
    backToHome() {
      this.$router.push("/");
    },
    DeleteGame(id) {
      this.$store.dispatch("game/deleteGame", id);
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
div {
  max-width: 100%;
  overflow-x: hidden;
}
</style>
