<template>
  <div class="text-center mt-5">
    <div class="container text-center">
      <h3 class="text-center">{{ title }}</h3>
      <h4 class="text-center">{{ subtitle }}</h4>
      <h5 class="text-center">ScoreBoard</h5>
    </div>
    <br /><br />
    <div class="accordion">
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Name</th>
            <th scope="col">Total Score</th>
          </tr>
        </thead>
        <tbody>
          <tr v-b-toggle.accordion-2 block variant="info">
            <th scope="row">Tapan</th>
            <td>99</td>
          </tr>
        </tbody>
      </table>

      <!-- Element to collapse -->
      <b-collapse
        id="accordion-2"
        visible
        accordion="my-accordion"
        role="tabpanel"
      >
        <b-card>
          <table class="table">
            <tr>
              <td>#</td>
              <th>Round 1</th>
              <th>Round 2</th>
              <th>Round 3</th>
            </tr>
            <tr>
              <th>throw 1</th>
              <td>11</td>
              <td>11</td>
              <td>11</td>
            </tr>
            <tr>
              <th>throw 2</th>
              <td>11</td>
              <td>11</td>
              <td>11</td>
            </tr>
            <tr>
              <th>throw 3</th>
              <td>11</td>
              <td>11</td>
              <td>11</td>
            </tr>
            <tr>
              <th>Total</th>
              <td>33</td>
              <td>33</td>
              <td>33</td>
            </tr>
          </table>
        </b-card>
      </b-collapse>
    </div>

    <br /><br />

    <div class="text-center">
      <h4>Congratulations <b>Jeel</b></h4>
      <h5>You Win This Game</h5>
      <br />
      <div class="d-grid gap-2 col-6 mx-auto">
        <button class="btn btn-success" type="button" @click="homepage">
          Home Page
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    subtitle: String,
    score: String,
  },
  methods: {
    homepage() {
      this.$router.push('/')
    },
  },
}
</script>
