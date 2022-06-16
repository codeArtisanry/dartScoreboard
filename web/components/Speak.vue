<template>
  <div>
    <div v-if="speaking" class="blur"></div>
    <div v-if="speaking" class="pt-5 loader">
      <ScaleLoader class="pt-5 mt-5" />
    </div>
  </div>
</template>
<script>
import ScaleLoader from "vue-spinner/src/ScaleLoader.vue";
export default {
  components: {
    ScaleLoader,
  },
  props: {
    text: {
      default: "",
      type: String,
    },
  },
  data() {
    return {
      speaking: false,
    };
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
      this.listenForSpeechEvents(speakText);
    },
    listenForSpeechEvents(speak) {
      speak.onstart = () => {
        this.speaking = true;
      };
      speak.onend = () => {
        this.speaking = false;
      };
    },
  },
};
</script>
<style scoped>
.loader {
  z-index: 50;
  width: 100vw;
  height: 100vh;
  position: absolute;
}
.blur {
  width: 100vw;
  height: 100vh;
  background-color: white;
  opacity: 0.6;
  position: absolute;
  z-index: 25;
}
</style>
