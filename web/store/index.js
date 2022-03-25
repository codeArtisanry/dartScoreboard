export const state = () => ({
  players: [],
})

export const mutations = {
  ADD_PLAYER(state, player) {
    state.players = [{ content: player }, ...state.players]
  },
  REMOVE_PLAYER(state, player) {
    state.players.splice(state.players.indexOf(player), 1)
  },
}
