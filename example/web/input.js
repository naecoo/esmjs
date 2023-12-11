import { ref, createApp } from "vue";

const app = createApp({
  setup() {
    return {
      count: ref(0),
    };
  },
});

app.mount("#app");

console.log("Hello World!");
