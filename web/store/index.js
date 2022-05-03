import jwtDecode from 'jwt-decode'

export const state = () => {
  return {
    auth: '',
    token: '',
  }
}
export const mutations = {
  setAuth(state, auth, token) {
    state.auth = auth
    state.token = token
  },
}
export const actions = {
  nuxtServerInit({ commit }, { req, app }) {
    let auth = ''
    let token = ''
    if (req.headers.cookie) {
      auth = app.$cookies.get('user')
      console.log(auth, 'hello world')
      try {
        token = jwtDecode(auth)
      } catch (err) {
        console.log(err)
        // this.$router.push('/signin')
        // No valid cookie found
      }
    }
    commit('setAuth', { auth, token })
  },
}
