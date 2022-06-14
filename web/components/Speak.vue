<template>
  <div></div>
</template>
<script>
export default {
  props: {
    text: {
      type: String,
    },
  },
  watch: {
    text() {
      this.speak(this.text);
    },
  },
  methods: {
    speak(text) {
      const speakText = new SpeechSynthesisUtterance(text);
      window.speechSynthesis.speak(speakText);
    },
  },
};
</script>
