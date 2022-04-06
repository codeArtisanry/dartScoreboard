export const state = () => ({
  auth: '',
  token: '',
})

export const mutations = {
  setAuthAndToken(state, { auth, token }) {
    state.auth = auth
    state.token = token
  },
}

export const actions = {
  getAndSetCookie({ commit }, payload) {
    commit('setAuthAndToken', payload)
  },
}
