<template>
  <div>
    <div class="row my-2 mx-1">
      <div class="col col-lg-1 col-sm-12">
        <img
          class="rounded-circle"
          height="60"
          src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmG5p8UupmMVKs875O7bXQYyptCrL_lhnMC6YoIvjalCiN8mkKiYCtJJxFqP2RS4r7k-A&usqp=CAU"
          alt="Profile pic"
        />
      </div>
      <div class="col pt-2 col-lg-6 col-6 col-md-12">
        <div class="text-right pt-2">
          <a
            :href="logoutURL"
            type="button"
            class="col-12 col-lg-6 shadow border border-light border-start-0 rounded-pill"
          >
            Log Out
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      logoutURL: process.env.LOGOUT_URL,
    }
  },
}
</script>
