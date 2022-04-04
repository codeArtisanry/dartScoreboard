<template>
  <main class="container">
    <div class="text-center pt-5">
      <h1 class="font-weight-bolder">Welcome</h1>
      <p class="p-md-2">Let's create a new Game</p>
    </div>
    <div class="">
      <input
        v-model="gameName"
        class="form-control ml-1"
        type="text"
        placeholder="Game Name..."
      />

      <div class="mt-5 mb-3 form-control ml-1">
        <select class="border-0 bg-white w-100" @change="onChange($event)">
          <option value="0">--select--</option>
          <option value="Highest Score Game">Highest Score Game</option>
          <option id="101" value="Target Score Game-101">
            Target Score Game 101
          </option>
          <option id="301" value="Target Score Game-301">
            Target Score Game 301
          </option>
          <option id="501" value="Target Score Game-501">
            Target Score Game 501
          </option>
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
                    To play high score darts, each player throws three darts per
                    turn and races to reach a predetermined target score. In
                    highest score game, completion of 3 rounds for all players
                    marks the game as finished and the player with highest score
                    wins.
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
                  After starting game first you want to select one from this If
                  you are choosing 101 then To play 101 darts the rules are
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
    <div class="d-inline-flex w-100 my-4">
      <input
        v-model="newPlayer"
        class="form-control w-75 ml-1"
        type="text"
        placeholder="Add a new Player"
        @keypress.enter="addPlayer"
      />
      <span class="btn btn-secondary ml-2 w-25" @click="addPlayer"> Add </span>
    </div>
    <div>
      <Player
        v-for="(player, i) in $store.state.players"
        :key="i"
        :player="player"
      />
    </div>
    <div class="row">
      <div class="col text-right">
        <button class="btn btn-secondary" @click="register">
          Register Game
        </button>
      </div>
      <div class="col text-left">
        <button class="btn btn-secondary"  @click="startgame(id)">
          Start Game
        </button>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  data() {
    return {
      registerGames: '',
      gameType: '',
      gameName: '',
      registerGame: {
        gameName: '',
        PlayersNames: [],
        gameType: '',
        gameTargetScore: 0,
      },
      newPlayer: '',
    }
  },

  methods: {
    addPlayer() {
      if (this.newPlayer) {
        this.$store.commit('ADD_PLAYER', this.newPlayer)
        this.newPlayer = ''
      }
    },
    register() {
      this.gamenamefunc()
      const playerArr = this.$store.state.players
      for (let i = 0; i <= playerArr.length - 1; i++) {
        this.registerGame.PlayersNames.push(playerArr[i].content)
      }
      this.postgamedata()
      if (
        this.registerGame.PlayersNames.length === 0 ||
        this.registerGame.PlayersNames.length === 1
      ) {
        alert('please enter players name more then one')
      } else {
      this.$router.push('/home')
    }
    },
    
  

    startgame() {
      this.gamenamefunc()
      const playerArr = this.$store.state.players
      for (let i = 0; i <= playerArr.length - 1; i++) {
        this.registerGame.PlayersNames.push(playerArr[i].content)
      }
      this.postgamedata()
      this.getGameData()
     
     
      if (
        this.registerGame.PlayersNames.length === 0 ||
        this.registerGame.PlayersNames.length === 1
      ) {
        alert('please enter players name more then one')
      } else {
        console.log(this.registerGame)
        if (this.registerGame.gameType === 'Highest Score Game') {
          this.$router.push('/start/highscoregame/')
        } else {
           this.$router.push('/start/highscoregame/')
        }
      }
    },
    onChange(event) {
      this.gameType = event.target.value
      this.gameType = this.gameType.split('-')
      this.registerGame.gameType = this.gameType[0]
      this.registerGame.gameTargetScore = Number(this.gameType[1])
    },

    gamenamefunc() {
      if (this.gameName === '') {
        this.registerGame.gameName = new Date().toLocaleString()
      } else {
        this.registerGame.gameName = this.gameName
      }
    },
    async postgamedata() {
      await this.$axios.$post(
       `${process.env.base_URL}/registerGame`,
        this.registerGame
      )
    },
    // clicked(id) {
    //   this.$router.push('/start/highscoregame' + id)
    // },
    async getGameData() {
      const res = await this.$axios.$get(
        `${process.env.base_URL}/registerGame` + this.$route.params.id
      )
      this.registerGames = res

    },
  },
}
</script>
