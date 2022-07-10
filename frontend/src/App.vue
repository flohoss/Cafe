<template>
  <TheNavigation @logout="logout" />
  <div class="container">
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { AuthorizationService } from "@/services/openapi";
import TheNavigation from "@/components/TheNavigation.vue";

export default defineComponent({
  name: "App",
  components: { TheNavigation },
  setup() {
    const router = useRouter();
    const store = useStore();

    function logout() {
      AuthorizationService.postAuthLogout().finally(() => {
        store.commit("logout");
        router.replace({ name: "Login" });
      });
    }
    return { logout };
  },
});
</script>

<style lang="less">
@import "primevue/resources/themes/saga-blue/theme.css";
@import "primevue/resources/themes/vela-blue/theme.css" screen and (prefers-color-scheme: dark);

@font-face {
  font-family: "roboto";
  src: url(@/assets/fonts/Roboto-Light.ttf);
}

.overflow-ellipsis {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

body {
  margin: 0.2rem;
  padding: 0.2rem;
  background-color: var(--surface-b);
}

html,
body {
  font-size: 1em;
  font-family: "roboto", sans-serif !important;
}
.container {
  --bs-gutter-x: 0;
  --bs-gutter-y: 0;
  width: 100%;
  padding-right: calc(var(--bs-gutter-x) * 0.5);
  padding-left: calc(var(--bs-gutter-x) * 0.5);
  margin-right: auto;
  margin-left: auto;
}

@media (min-width: 576px) {
  .container {
    max-width: 540px;
  }
}
@media (min-width: 768px) {
  .container {
    max-width: 720px;
  }
}
@media (min-width: 992px) {
  .container {
    max-width: 960px;
  }
}
@media (min-width: 1200px) {
  .container {
    max-width: 1140px;
  }
}
@media (min-width: 1400px) {
  .container {
    max-width: 1320px;
  }
}
</style>
