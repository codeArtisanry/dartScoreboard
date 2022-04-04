<template>
  <div class="">
    <div class="container my-5 mt-5 mb-3 text-center">
      <h2 class="font-weight-bolder">Game Info</h2>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Type</th>
          </tr>
        </thead>
        <tbody>
          <tr @click="clicked">
            <th scope="row">alpha</th>
            <td>TargetScore</td>
          </tr>

          <tr @click="clicked">
            <th scope="row">delta</th>
            <td>HighScore</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  methods: {
    clicked() {
      this.$router.push('/home/creategame')
    },
  },
}
</script>
