<template>
  <b-navbar toggleable type="light" variant="light">
    <b-navbar-brand>NavBar</b-navbar-brand>

    <b-navbar-toggle target="navbar-toggle-collapse">
      <template #default="{ expanded }">
        <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon>
        <b-icon v-else icon="chevron-bar-down"></b-icon>
      </template>
    </b-navbar-toggle>

    <b-collapse id="navbar-toggle-collapse" is-nav>
      <b-navbar-nav class="ml-auto">
        <b-nav-item href="http://localhost:3000/">Home</b-nav-item>
        <b-nav-item href="http://localhost:3000/games/new"
          >Create Game</b-nav-item
        >
        <b-nav-item href="http://localhost:8080/logout/google" type="button"
          >Log Out</b-nav-item
        >
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>

<script>
export default {
  data() {
    return {}
  },
}
</script>
