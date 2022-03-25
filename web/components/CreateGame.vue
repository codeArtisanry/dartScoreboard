<template>
  <main class="container">
    <div class="text-center pt-5">
      <h1 class="font-weight-bolder">Welcome</h1>
      <p class="p-md-2">Let's create a new Game</p>
    </div>
    <div class="">
      <input
        class="shadow border form-control border-light rounded-pill ml-1"
        type="text"
        placeholder="Game Name..."
      />
    </div>
    <div class="d-inline-flex w-100 my-4">
      <input
        v-model="newPlayer"
        class="form-control w-75 shadow border border-light rounded-pill"
        type="text"
        placeholder="Add a new Player"
        @keypress.enter="addPlayer"
      />
      <span
        class="btn btn-secondary btn-sm shadow border border-light rounded-pill ml-2 w-25 fw-bolder"
        @click="addPlayer"
      >
        Add
      </span>
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
        <button class="btn btn-secondary">Register Game</button>
      </div>
      <div class="col text-left">
        <button class="btn btn-secondary">Start Game</button>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  data() {
    return {
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
  },
}
</script>
