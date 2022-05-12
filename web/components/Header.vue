<template>
  <div>
    <div class="text-center">
      <div class="col pt-2 text-center">
        <h5 class="text-muted">Welcome to</h5>
        <h4 class="pb-4 bold font-weight-bold">Dart Scoreboard</h4>
      </div>
      <div class="px-3">
        <button
          class="
            shadow
            px-5
            py-3
            font-weight-bolder
            border border-light border-start-0
            rounded
          "
          @click="creategame"
        >
          New Game
        </button>
        <br />
      </div>
      <hr />
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    creategame() {
      this.$router.push("/games/new");
    },
  },
};
</script>
