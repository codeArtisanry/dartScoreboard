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
        <a
          :href="$config.loginURL"
          class="btn btn-primary mb-5 py-2 px-5"
          data-test="login-button"
          type="button"
        >
          Sign In With Google
        </a>
      </b-card>
    </center>
  </div>
</template>
