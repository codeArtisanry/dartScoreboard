<template>
  <div>
    <div class="col float-right col-lg-4 mt-lg-5 mb-lg-3 my-md-2 text-center">
      <div class="col pt-2 text-center">
        <h1 class="text-muted">Welcome to</h1>
        <h3 class="pt-2 pb-4 bold font-weight-bold">Dart Scoreboard</h3>
      </div>
      <div class="px-3">
        <button
          class="
            col col-lg-12
            shadow
            p-3
            mb-3
            border border-light border-start-0
            rounded-pill
          "
          @click="creategame"
        >
          <b @click="fetch">Create Game</b>
        </button>
        <br />
      </div>
      <hr />
    </div>
  </div>
</template>

<script>
import Cookies from 'js-cookie'
import jwtDecode from 'jwt-decode'
export default {
  mounted() {
    const auth = Cookies.get('user')
    const token = jwtDecode(auth)
    this.$store.dispatch('cookie/getAndSetCookie', { auth, token })
  },
  // updated() {
  //   const auth = Cookies.get('user')
  //   const token = jwtDecode(auth)
  //   this.$store.dispatch('cookie/getAndSetCookie', { auth, token })
  // },
  methods: {
    creategame() {
      this.$router.push('/home/creategame')
    },
  },
}
</script>
