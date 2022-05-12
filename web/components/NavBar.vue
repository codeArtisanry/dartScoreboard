<template>
  <nav class="bg-light py-1 row">
    <div class="col col-4">
      <b-button v-b-toggle.sidebar-no-header class="bg-light ml-4 border-0">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          height="30"
          viewBox="0 0 448 512"
        >
          <path
            d="M0 96C0 78.33 14.33 64 32 64H416C433.7 64 448 78.33 448 96C448 113.7 433.7 128 416 128H32C14.33 128 0 113.7 0 96zM0 256C0 238.3 14.33 224 32 224H416C433.7 224 448 238.3 448 256C448 273.7 433.7 288 416 288H32C14.33 288 0 273.7 0 256zM416 448H32C14.33 448 0 433.7 0 416C0 398.3 14.33 384 32 384H416C433.7 384 448 398.3 448 416C448 433.7 433.7 448 416 448z"
          />
        </svg>
      </b-button>
    </div>
    <div class="col col-7 text-right mt-1">
      <img src="/dart-icon.png" height="35" alt="dart-icon" />
    </div>
    <b-sidebar
      id="sidebar-no-header"
      aria-labelledby="sidebar-no-header-title"
      no-header
      shadow
    >
      <template #default="{ hide }">
        <div class="p-3 text-center">
          <img
            :src="$store.state.auth.token.picture"
            class="rounded-circle my-4"
            height="50"
            alt="Profile Pic"
          />
          <h6 class="text-muted font-weight-bolder">
            {{ $store.state.auth.token.name }}
          </h6>
          <h6 class="text-muted">{{ $store.state.auth.token.email }}</h6>
          <hr />
          <b-nav vertical>
            <b-button
              class="
                bg-light
                text-dark
                mt-3
                border-0
                hover
                font-weight-bolder
                text-dark
              "
              @click="home"
              >Home</b-button
            >
            <b-link
              type="button"
              class="
                btn btn-light btn-link
                mt-1
                mb-4
                border-0
                font-weight-bolder
                text-dark text-decoration-none
              "
              :href="logoutURL"
              @click="hide"
            >
              Sign Out
            </b-link>
          </b-nav>
          <b-button variant="secondary" block @click="hide">Close</b-button>
        </div>
      </template>
    </b-sidebar>
  </nav>
</template>
<script>
export default {
  data() {
    return {
      logoutURL: process.env.LOGOUT_URL,
    };
  },
  methods: {
    home() {
      this.$router.push("/");
    },
  },
};
</script>
