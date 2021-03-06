import { createApp } from "vue";
import { OpenAPI } from "@/services/openapi";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import PrimeVue from "primevue/config";
import ConfirmationService from "primevue/confirmationservice";
import ToastService from "primevue/toastservice";

import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";
import { ItemType } from "@/utils";

export const API_ENDPOINT_URL = process.env.NODE_ENV === "production" ? window.origin + "/api" : "http://localhost:5000/api";
export const WEBSOCKET_ENDPOINT_URL = API_ENDPOINT_URL.replace("http", "ws") + "/orders/ws";
OpenAPI.BASE = API_ENDPOINT_URL;
OpenAPI.WITH_CREDENTIALS = true;

store
  .dispatch("getOrderItems", ItemType.Food)
  .then(() => store.commit("login"))
  .catch(() => store.commit("logout"))
  .finally(() => {
    const app = createApp(App);

    app.use(router);
    app.use(PrimeVue, {
      locale: {
        startsWith: "Beginnt mit",
        contains: "enthält",
        notContains: "enthält nicht",
        endsWith: "endet mit",
        equals: "entspricht",
        notEquals: "entspricht nicht",
        noFilter: "Kein Filter",
        lt: "Weniger als",
        lte: "Weniger als oder gleich viel",
        gt: "Mehr als",
        gte: "Mehr als oder gleich viel",
        dateIs: "Datum ist",
        dateIsNot: "Datum ist nicht",
        dateBefore: "Datum liegt vor",
        dateAfter: "Datum liegt nach",
        clear: "Löschen",
        apply: "Anwenden",
        matchAll: "Alle abgleichen",
        matchAny: "Mit jedem abgleichen",
        addRule: "Regel hinzufügen",
        removeRule: "Regel entfernen",
        accept: "Ja",
        reject: "Nein",
        choose: "Auswählen",
        upload: "Hochladen",
        cancel: "Abbrechen",
        dayNames: ["Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"],
        dayNamesShort: ["Son", "Mon", "Die", "Mit", "Don", "Fre", "Sam"],
        dayNamesMin: ["So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"],
        monthNames: ["Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"],
        monthNamesShort: ["Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"],
        today: "Heute",
        weekHeader: "Wk",
        firstDayOfWeek: 1,
        dateFormat: "dd.mm.yy",
        weak: "Schwach",
        medium: "Medium",
        strong: "Stark",
        passwordPrompt: "Passwort eingeben",
        emptyFilterMessage: "Keine Ergebnisse gefunden",
        emptyMessage: "Keine verfügbaren Optionen",
      },
    });
    app.use(store);
    app.use(ConfirmationService);
    app.use(ToastService);

    router.isReady().then(() => {
      app.mount("#app");
    });
  });
