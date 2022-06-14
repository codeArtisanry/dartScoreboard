<template>
  <div>
    <NavBar />
    <center>
      <div class="text-center container mx-2">
        <p class="font-weight-bolder">Game Name:</p>
        <p class="text text-muted" data-test="game-name">{{ game.name }}</p>
        <hr />
        <p class="font-weight-bolder">Game Type:</p>
        <p class="text text-muted" data-test="game-type">{{ game.type }}</p>
        <hr />
        <p class="font-weight-bolder">Creater Name:</p>
        <p class="text text-muted" data-test="creater-name">
          {{ game.creater_name }}
        </p>
        <hr />
        <p class="font-weight-bolder">Players Names:</p>
        <p
          v-for="player in game.players"
          :key="player.id"
          class="text-muted"
          data-test="players-name"
        >
          {{ player.first_name }}
        </p>
      </div>
      <div class="row">
        <div
          v-if="game.status === 'Not Started' && isOwner"
          class="col text-center"
        >
          <button
            class="btn btn-info"
            data-test="update-button"
            @click="redirectToUpdatePage"
          >
            Update
          </button>
        </div>
        <div v-if="isOwner" class="col text-center">
          <button
            class="btn btn-danger"
            data-test="delete-button"
            @click="DeleteGameById(game.id)"
          >
            Delete
          </button>
        </div>
        <div class="col text-center">
          <button
            class="btn btn-success"
            data-test="game-state-button"
            @click="redirectToCurrentGameState"
          >
            {{ buttonName }}
          </button>
        </div>
      </div>
      <div class="text-center mt-5">
        <button
          class="btn btn-secondary"
          data-test="home-button"
          @click="redirectToHomePage"
        >
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
      buttonName: "",
    };
  },

  computed: {
    game() {
      return this.$store.getters["game/details"];
    },
    isOwner() {
      return this.$store.getters["game/isOwner"];
    },
  },

  async created() {
    await this.getGame();
    this.updateButtonName();
  },

  methods: {
    async getGame() {
      await this.$store.dispatch("game/getGame", this.$route.params.gameid);
    },

    updateButtonName() {
      if (this.game.status === "Not Started") {
        this.buttonName = "Start";
      } else if (this.game.status === "In Progress") {
        this.buttonName = "Resume";
      } else {
        this.buttonName = "Scoreboard";
      }
    },

    redirectToUpdatePage() {
      this.$router.push(`/games/${this.$route.params.gameid}/update`);
    },

    redirectToHomePage() {
      this.$router.push("/");
    },

    redirectToCurrentGameState() {
      if (this.game.status === "Completed") {
        this.$router.push(`/games/${this.$route.params.gameid}/scoreboard`);
      } else {
        this.$router.push(`/games/${this.$route.params.gameid}/player`);
      }
    },

    DeleteGameById(id) {
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
