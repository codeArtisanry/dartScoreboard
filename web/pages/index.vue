<template>
  <div>
    <NavBar />
    <Header />
    <DetailedDescription />
  </div>
</template>
<script>
export default {
  middleware: 'notauth',
}
</script>
