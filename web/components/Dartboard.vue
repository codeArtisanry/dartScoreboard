<template>
  <div
    id="dartboard"
    class="d-inline-block"
    style="width: 22rem; height: 22rem"
  ></div>
</template>
<script>
export default {
  mounted() {
    // eslint-disable-next-line no-undef
    const dartboard = new Dartboard("#dartboard");
    dartboard.render();
    document.querySelector("#dartboard").addEventListener("throw", (d) => {
      this.$emit("onAddScore", d.detail);
    });
  },
};
</script>
