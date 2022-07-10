import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import PrimeVue from "primevue/config";

import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";

const app = createApp(App);

app.use(router);
app.use(PrimeVue);
app.use(store);

router.isReady().then(() => {
  app.mount("#app");
});
