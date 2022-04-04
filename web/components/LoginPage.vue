<template>
  <div>
    <center>
      <b-card tag="article" class="card mt-5">
        <h3 class="mt-3"><b>Welcome To Darts Scoreboard</b></h3>
        <hr />
        <br />
        <img
          src="darts.png"
          class="img-fluid mx-auto d-block w-50"
          alt="Responsive image"
        />
        <hr />
        <h4>Not a member? <fa icon="down-long" /></h4>
        <br />
        <button
          class="btn btn-primary mb-5 py-2 px-5"
          type="button"
          @click="login"
        >
          Sign In With Google</button
        ><br />
      </b-card>
    </center>
  </div>
</template>

<script>
export default {
  methods: {
    login() {
      this.$router.push('/home')
    },
  },
}
</script>