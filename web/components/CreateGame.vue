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
    <br>
   <div>
  <label class="typo__label">Select PlayersNames</label>
  <multiselect v-model="value" :options="options" :custom-label="nameWithLang" :multiple="true" :close-on-select="false" :clear-on-select="false" :preserve-search="true" placeholder="Pick some" label="firstname" track-by="firstname" :preselect-first="true">
    <template slot="selection" slot-scope="{ values, isOpen }" ><span v-if="values.length &amp;&amp; !isOpen" class="multiselect__single" >{{ values.length }} options selected</span></template>
  </multiselect>
  <pre class="language-json"><code v-for="i in value" :key="i">{{ i.firstname }} {{i.lastname}}<br></code><br></pre><br>
</div>

    <div class="row">
      <div class="col text-right">
        <button class="btn btn-secondary" @click="register">
          Register Game
        </button>
      </div>
      <div class="col text-left">
        <button class="btn btn-secondary" @click="startgame(id)">
          Start Game
        </button>
      </div>
    </div>
  </main>
</template>

<script>
import Multiselect from 'vue-multiselect';
export default {
  components: { Multiselect },
  data() {
    return {
      value: [],
      options: [
        {firstname:'payal', lastname:'raviya', email:'payal@improwised.com' },
        {firstname:'jeel', lastname:'rupapara', email:'jeel@improwised.com' },
        {firstname:'vatsal', lastname:'chauhan', email:'vatsal@improwised.com' },
        {firstname:'munir', lastname:'khakhi', email:'munir@improwised.com' },
        {firstname:'tapan', lastname:'bavaliya', email:'tapan@improwised.com' },
       {firstname:'rakshit', lastname:'menpara', email:'rakshit@improwised.com' },
      ],
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
      pname:[]
    }
  },

  methods: {
   
    nameWithLang ({ firstname,lastname, email }) {
      return `${firstname} ${lastname} — [${email}]`
    },
   
    register() {
      this.gamenamefunc()
      if (
        this.value.length === 0 ||
        this.value.length === 1
      ) {
        alert('please enter players name more then one')
      } else {
        this.$router.push('/home')
      console.log(this.value[0].firstname);
      for(let i=0; i<=this.value.length-1;i++){
        this.registerGame.PlayersNames.push(this.value[i].firstname)
      }
      console.log(this.registerGame.PlayersNames)
     this.postgamedata()
        }
      },

    startgame() {
      this.gamenamefunc()
      this.postgamedata()
      // this.getGameData()

      if (
        this.value.length === 0 ||
        this.value.length === 1
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
      console.log(this.registerGame)
    },
    // async getGameData() {
    //   const res = await this.$axios.$get(
    //     `${process.env.base_URL}/registerGame`
    //   )
    //   this.registerGames = res
    //   console.log(this.registerGame)
    // },
  },
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css">
