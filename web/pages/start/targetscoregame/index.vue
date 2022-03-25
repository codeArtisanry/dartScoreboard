<template>
  <Startgame :game-type="gameType" />
</template>
<script>
import Startgame from '@/components/Startgame.vue'
export default {
  components: { Startgame },
  data() {
    return {
      gameType: {
        type: 'Target-Score Game',
        score: 'Remaining-Score',
      },
    }
  },
}
</script>
