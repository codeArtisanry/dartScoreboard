<template>
  <div class="">
    <div class="text-capitalize">
      <table class="table">
        <tbody>
          <tr>
            <th class="pl-5 text-left" scope="row">{{ player.content }}</th>
            <td class="pr-5 text-right" scope="col">
              <button
                class="
                  delete
                  shadow
                  border border-light
                  rounded-pill
                  border-danger
                  px-3
                "
                @click="removePlayer"
              >
                X
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  props: ['player'],
  methods: {
    removePlayer() {
      this.$store.commit('REMOVE_PLAYER', this.player)
    },
  },
}
</script>
