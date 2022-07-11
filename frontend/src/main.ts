import { createApp } from "vue";
import { TablesService, OpenAPI } from "@/services/openapi";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import PrimeVue from "primevue/config";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCheese } from "@fortawesome/free-solid-svg-icons";
import { faChampagneGlasses } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
library.add(faCheese, faChampagneGlasses);

import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";

export const API_ENDPOINT_URL = process.env.NODE_ENV === "production" ? window.origin + "/api" : "http://localhost:5000/api";
export const WEBSOCKET_ENDPOINT_URL = API_ENDPOINT_URL.replace("http", "ws") + "/system/ws";
OpenAPI.BASE = API_ENDPOINT_URL;
OpenAPI.WITH_CREDENTIALS = true;

TablesService.getTables()
  .then((res) => {
    store.commit("setTables", res);
    store.commit("login");
  })
  .catch(() => store.commit("logout"))
  .finally(() => {
    const app = createApp(App);

    app.use(router);
    app.use(PrimeVue);
    app.use(store);
    app.component("font-awesome-icon", FontAwesomeIcon);

    router.isReady().then(() => {
      app.mount("#app");
    });
  });
