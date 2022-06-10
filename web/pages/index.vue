<template>
  <div>
    <NavBar />
    <Header />
    <GamesList />
  </div>
</template>
<script>
export default {
  middleware: 'notauth',
}
</script>
