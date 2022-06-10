<template>
  <div>
    <div class="container my-5 mt-5 mb-3 text-center">
      <h2 class="font-weight-bolder">Game Info</h2>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Type</th>
            <th scope="col">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="game in gamesList"
            :key="game.id"
            @click="redirectToGameInfoById(game.id)"
          >
            <td>{{ game.name }}</td>
            <td>{{ game.type }}</td>
            <td>{{ game.status }}</td>
          </tr>
        </tbody>
      </table>
      <button
        v-if="previous !== 'null'"
        variant="outline-primary"
        class="btn btn-sm btn-secondary"
        @click="getGameListByPageLink(previous)"
      >
        Back
      </button>
      <button
        v-if="next !== 'null'"
        variant="outline-primary"
        class="btn btn-sm btn-secondary"
        data-test="next"
        @click="getGameListByPageLink(next)"
      >
        Next
      </button>
    </div>
  </div>
</template>
<script>
/* eslint-disable no-console */
export default {
  computed: {
    gamesList() {
      return this.$store.getters["games/list"];
    },
    next() {
      return this.$store.getters["games/next"];
    },
    previous() {
      return this.$store.getters["games/previous"];
    },
  },
  async created() {
    await this.getGamesList("page=1");
  },
  methods: {
    async getGamesList(page) {
      await this.$store.dispatch("games/getGames", page);
    },
    async getGameListByPageLink(pageLink) {
      const page = pageLink.split("?")[1];
      await this.getGamesList(page);
    },
    redirectToGameInfoById(id) {
      this.$router.push(`/games/${id}`);
    },
  },
};
</script>
