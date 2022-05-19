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
            v-for="game in getGames.game_responses"
            :key="game.id"
            @click="clicked(game.id)"
          >
            <td>{{ game.game_name }}</td>
            <td>{{ game.game_type }}</td>
            <td>{{ game.game_status }}</td>
          </tr>
        </tbody>
      </table>
      <button
        v-if="getGames.pre_page_link == 'cross limits'"
        variant="outline-primary"
        class="btn btn-sm btn-secondary invisible"
        @click="prepage"
      >
        Back
      </button>
      <button
        v-else
        variant="outline-primary"
        class="btn btn-sm btn-secondary col my-2"
        @click="prepage"
      >
        Back
      </button>
      <button
        v-if="getGames.post_page_link == 'cross limits'"
        variant="outline-primary"
        class="btn btn-sm btn-secondary invisible"
        @click="postpage"
      >
        Next
      </button>

      <button
        v-else
        variant="outline-primary"
        class="btn btn-sm btn-secondary col"
        @click="postpage"
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
    getGames() {
      return this.$store.state.game.games;
    },
  },
  async created() {
    await this.gamesApi();
  },
  methods: {
    clicked(id) {
      this.$router.push("/games/" + id);
    },
    prepage() {
      this.$store.dispatch("game/getPrePage", this.getGames.pre_page_link);
    },
    postpage() {
      this.$store.dispatch("game/getPostPage", this.getGames.post_page_link);
    },
    async gamesApi() {
      await this.$store.dispatch("game/getGames");
    },
  },
};
</script>
