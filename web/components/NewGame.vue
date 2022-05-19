<template>
  <div>
    <NavBar />
    <main class="container">
      <div class="text-center pt-5">
        <h1 class="font-weight-bolder">Welcome</h1>
        <p class="p-md-2">Let's create a new Game</p>
      </div>
      <div class="ml-1">
        <label class="text-muted">Enter Game Name:</label>
        <input
          v-model="game_responses.game_name"
          class="form-control"
          type="text"
          placeholder="Game Name..."
        />

        <div class="mt-5 mb-3">
          <label class="text-muted">Select Game Type:</label>
          <select
            v-model="game_responses.game_type"
            class="form-control bg-white w-100"
          >
            <option value="High Score">High Score</option>
            <option value="Target Score-101">Target Score-101</option>
            <option value="Target Score-301">Target Score-301</option>
            <option value="Target Score-501">Target Score-501</option>
          </select>
        </div>
        <small class="form-text text-muted ml-4"
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
        <label class="typo__label text-muted">Select PlayersNames:</label>
        <multiselect
          v-model="value"
          :options="options"
          :custom-label="nameWithLang"
          :multiple="true"
          :close-on-select="false"
          :clear-on-select="false"
          :preserve-search="true"
          :options-limit="5"
          placeholder="Pick some"
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
        ><code v-for="player in value" :key="player.email">{{ player.first_name }} {{player.last_name}}<br></code><br></pre>
        <br />
      </div>

      <div class="row">
        <div class="col text-center">
          <button
            v-if="$route.params.gameid !== undefined"
            class="btn btn-secondary"
            @click="updatedata"
          >
            Update
          </button>
          <button v-else class="btn btn-secondary" @click="register">
            Register Game
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
      value: [],
      options: [],
      registerGames: "",
      game_responses: {
        game_name: "",
        game_type: "",
        players: [],
      },
      newPlayer: "",
      player: [],
      nameofgame: [],
      totalplayers: 0,
    };
  },
  computed: {
    getGame() {
      return this.$store.state.game.game;
    },
    getUsers() {
      return this.$store.state.game.users;
    },
  },
  async created() {
    await this.users();
    if (this.$route.params.gameid !== undefined) {
      await this.game();
      this.withUpdate();
    }
  },
  methods: {
    // eslint-disable-next-line camelcase
    nameWithLang({ first_name, last_name, email }) {
      // eslint-disable-next-line camelcase
      return `${first_name} ${last_name} — [${email}]`;
    },
    async users() {
      await this.$store.dispatch("game/getUsers");
      this.options = this.getUsers.user_responses;
    },
    register() {
      if (this.value.length === 0) {
        alert("please enter players name at list one");
      } else {
        this.$router.push("/");
        for (let i = 0; i <= this.value.length - 1; i++) {
          this.game_responses.players.push(this.value[i].id);
          this.nameofgame.push(this.value[i].first_name);
        }
        if (this.game_responses.game_name === "") {
          this.totalplayers = this.nameofgame.length - 1;
          this.game_responses.game_name =
            this.nameofgame[0] + ` +${this.totalplayers}  others`;
        }
        this.postGame();
      }
    },
    async postGame() {
      await this.$store.dispatch("game/postGame", this.game_responses);
    },
    async game() {
      await this.$store.dispatch("game/getGame", this.$route.params.gameid);
      this.game_responses = this.getGame;
    },
    async updatedata() {
      this.options = this.getUsers.user_responses;
      this.game_responses.players = [];
      for (let j = 0; j <= this.value.length - 1; j++) {
        for (let k = 0; k <= this.options.length - 1; k++) {
          if (this.value[j].id === this.options[k].id) {
            this.game_responses.players.push(this.options[k].id);
          }
        }
      }
      await this.$store.dispatch("game/updateGame", {
        gameId: this.$route.params.gameid,
        update: this.game_responses,
      });
      this.$router.push("/");
    },
    withUpdate() {
      this.value = this.getGame.players;
    },
  },
};
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css">
