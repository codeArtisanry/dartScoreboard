<template>
  <div>
    <NavBar />
    <main class="container">
      <div>
        <b-alert
          v-model="alert"
          class="position-fixed fixed-top py-3 rounded-0"
          style="z-index: 2000"
          variant="warning"
          dismissible
        >
          Please, add at least one player!!!
        </b-alert>
      </div>
      <div class="text-center pt-5">
        <h1 class="font-weight-bolder">Welcome</h1>
        <p class="p-md-2">Let's create a new Game</p>
      </div>
      <div>
        <small
          style="color: red; position: absolute; top: 180px; right: 20px"
          class="pl-5"
          >Note : * Indicates mandatory field.</small
        >
      </div>
      <div class="ml-1">
        <label class="text-muted">Enter Game Name:</label>
        <input
          v-model="name"
          class="form-control"
          type="text"
          placeholder="Game Name..."
        />

        <div class="mt-5 mb-3">
          <label class="text-muted"
            >Select Game Type<span style="color: red"> *</span></label
          >
          <select v-model="type" class="form-control w-100">
            <option value="High Score">High Score</option>
            <option value="Target Score-101">Target Score-101</option>
            <option value="Target Score-301">Target Score-301</option>
            <option value="Target Score-501">Target Score-501</option>
          </select>
        </div>
        <small class="form-text text-muted"
          >Rules of the Game
          <span>
            <a v-b-modal.modalPopover class="bg-white text-primary border-0"
              >Read...</a
            >
            <b-modal id="modalPopover" title="Rules For Game" ok-only>
              <div class="card-deck">
                <div class="card">
                  <div class="card-body">
                    <h2 class="text-center">
                      <b><u>Highest Score Game</u></b>
                    </h2>
                    <br /><br />
                    <p class="card-text">
                      <b>How to play high score darts?</b><br />
                      To play high score darts, each player throws three darts
                      per turn and races to reach a predetermined target score.
                      In highest score game, completion of 3 rounds for all
                      players marks the game as finished and the player with
                      highest score wins.
                    </p>
                    <br /><br />
                  </div>
                </div>
                <div class="card">
                  <div class="card-body">
                    <h2 class="text-center">
                      <b><u>Target Score Game</u></b>
                    </h2>
                    <br />
                    <p class="card-text">
                      <b>How to play target score darts?</b>
                    </p>

                    <h6>There are three types in target score game</h6>
                    <ul>
                      <li>101</li>
                      <li>301</li>
                      <li>501</li>
                    </ul>
                    After starting game first you want to select one from this
                    If you are choosing 101 then To play 101 darts the rules are
                    simple, both players or teams start with a score of 101
                    points. Each player then takes alternating turns at throwing
                    their darts at the dartboard. The points scored are removed
                    from the total, and then the opposing player/team does the
                    same. The first to reach zero wins the game. For 301 and 501
                    rules are same but you want to start with 301 or 501 Score
                    <br /><br />
                  </div>
                </div>
              </div>
            </b-modal>
          </span>
        </small>
      </div>
      <br />
      <div>
        <label class="typo__label text-muted mt-3"
          >Select PlayersNames<span style="color: red"> *</span></label
        >
        <multiselect
          v-model="players"
          :options="usersList"
          :custom-label="optionsFormat"
          :multiple="true"
          :close-on-select="false"
          :clear-on-select="true"
          :preserve-search="true"
          placeholder="Search Player Name"
          label="first_name"
          track-by="first_name"
        >
          <template slot="selection" slot-scope="{ values, isOpen }"
            ><span
              v-if="values.length &amp;&amp; !isOpen"
              class="multiselect__single"
              >{{ values.length }} options selected</span
            ></template
          >
        </multiselect>
        <pre
          class="language-json"
        ><code v-for="player in players" :key="player.email">{{ player.first_name }} {{player.last_name}}<br></code><br></pre>
        <br />
      </div>

      <div class="row">
        <div class="col text-center">
          <button
            v-if="$route.params.gameid !== undefined"
            class="btn btn-secondary"
            @click="update"
          >
            Update
          </button>
          <button v-else class="btn btn-secondary" @click="registerNewGame">
            Register
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import Multiselect from "vue-multiselect";
export default {
  components: { Multiselect },
  data() {
    return {
      alert: false,
      usersList: [],
      name: "",
      type: "High Score",
      players: [],
      game: {
        name: "",
        type: "",
        players: [],
      },
    };
  },
  computed: {
    gameDetails() {
      return this.$store.getters["game/details"];
    },
    users() {
      return this.$store.getters["users/list"];
    },
  },

  async created() {
    await this.getUsers();
    if (this.$route.params.gameid) {
      await this.getGameById(this.$route.params.gameid);
      this.name = this.gameDetails.name;
      this.type = this.gameDetails.type;
      this.players = this.gameDetails.players;
    }
  },

  methods: {
    // eslint-disable-next-line camelcase
    optionsFormat({ first_name, last_name, email }) {
      // eslint-disable-next-line camelcase
      return `${first_name} ${last_name} — [${email}]`;
    },

    // get users list from users api
    async getUsers() {
      await this.$store.dispatch("users/getUsers");
      this.usersList = this.users;
    },

    // get game details by game id from game api
    async getGameById(id) {
      await this.$store.dispatch("game/getGame", id);
      this.name = this.gameDetails.name;
      this.type = this.gameDetails.type;
      this.players = this.gameDetails.players;
    },

    // redirect to home page
    redirectToHome() {
      this.$router.push("/");
    },

    // fill game details in game object to create a game
    createGameObject() {
      if (!this.name) {
        this.generateName();
      }
      this.game.name = this.name;
      this.game.type = this.type;
      this.players.forEach((player) => {
        this.game.players.push(player.id);
      });
    },

    // generate game name for users don't give any name of game
    generateName() {
      const totalPlayers = this.players.length;
      this.name = `${this.players[0].first_name} (+${totalPlayers - 1} others)`;
    },

    // create a game using game details
    async createGame() {
      this.createGameObject();
      await this.$store.dispatch("game/createGame", this.game);
    },

    // register a new game
    async registerNewGame() {
      if (!this.players.length) {
        this.alert = true;
        setTimeout(() => (this.alert = false), 3000);
      } else {
        await this.createGame();
        this.redirectToHome();
      }
    },

    // update a game details
    async update() {
      if (!this.players.length) {
        this.alert = true;
        setTimeout(() => (this.alert = false), 3000);
      } else {
        this.createGameObject();
        await this.$store.dispatch("game/updateGame", {
          id: this.gameDetails.id,
          update: this.game,
        });
        this.redirectToHome();
      }
    },
  },
};
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css" />
