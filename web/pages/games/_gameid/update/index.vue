<template>
  <NewGame :game-details="gameDetails" />
</template>
<script>
export default {
  middleware: 'notauth',

  data() {
    return {
      gameDetails: '',
    }
  },
  methods: {
    async getGame() {
      const res = await this.$axios.$get(
        `${process.env.base_URL}/registerGame/` + this.$route.params.gameid
      )
      this.gameDetails = res
    },
  },
}
</script>
