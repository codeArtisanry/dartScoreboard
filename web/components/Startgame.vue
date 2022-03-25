<template>
  <div>
    <div class="container-fluid text-center mt-4">
      <div class="shadow bg-white py-5 px-5 rounded">
        <h5 class="heading">Welcome to</h5>
        <h4 class="font-weight-bolder">{{ gameType.type }}</h4>

        <table class="table table-striped shadow mt-3">
          <tbody>
            <tr>
              <th class="text-left" scope="row">Round</th>
              <td scope="col">1</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">Player Name</th>
              <td scope="col">jeel</td>
            </tr>
          </tbody>
          <tbody>
            <tr>
              <th class="text-left" scope="row">{{ gameType.score }}</th>
              <td>100</td>
            </tr>
          </tbody>
        </table>
        <div>
          <div class="mt-2">
            <input
              class="form-control shadow border border-light rounded-pill p-3 bg-white"
              type="number"
              placeholder="First try point..."
            />
          </div>
          <div class="mt-2">
            <input
              class="form-control shadow border border-light rounded-pill p-3 bg-white"
              type="number"
              placeholder="Second try point..."
            />
          </div>
          <div class="mt-2">
            <input
              class="form-control shadow border border-light rounded-pill p-3 bg-white"
              type="number"
              placeholder="Third try point..."
            />
          </div>
        </div>

        <div class="mt-4 mb-2">
          <button class="btn btn-secondary">Next player</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    gameType: Object,
  },
}
</script>
