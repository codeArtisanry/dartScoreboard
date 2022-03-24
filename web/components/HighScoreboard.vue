<template>
  <div>
    <center>
      <b-card tag="article" style="max-width: 90rem" class="mb-2 mt-5">
        <h3>{{ title }}</h3>
        <h4>{{ subtitle }}</h4>
        <h5>ScoreBoard</h5>
        <hr />

        <table class="table">
          <thead class="thead-dark">
            <tr>
              <th scope="col">#</th>
              <th scope="col">Team Member Name</th>
              <th scope="col">{{ score }}</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <th scope="row">1</th>
              <td>Jeel</td>
              <td>120</td>
            </tr>
            <tr>
              <th scope="row">2</th>
              <td>Vatsal</td>
              <td>134</td>
            </tr>
            <tr>
              <th scope="row">3</th>
              <td>Payal</td>
              <td>125</td>
            </tr>
          </tbody>
        </table>
        <br /><br />

        <h4>Congratulations <b>Jeel</b></h4>
        <h5>You Win This Game</h5>
        <br />
        <div class="d-grid gap-2 col-6 mx-auto">
          <button class="btn btn-success" type="button">Show DeskBoard</button
          ><br /><br />
          <button class="btn btn-success" type="button">Play Again</button>
        </div>
      </b-card>
    </center>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    subtitle: String,
    score: String,
  },
}
</script>
