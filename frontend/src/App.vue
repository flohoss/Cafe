<template>
  <Toast style="width: 90vw" position="bottom-right" group="br" />
  <TheNavigation @logout="logout" />
  <div class="m-2">
    <router-view v-slot="{ Component }" mode="out-in">
      <transition>
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { AuthorizationService } from "@/services/openapi";
import TheNavigation from "@/components/UI/TheNavigation.vue";
import Toast from "primevue/toast";

export default defineComponent({
  name: "App",
  components: { TheNavigation, Toast },
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
@import "assets/css/all.min.css";

@font-face {
  font-family: "roboto";
  src: url(@/assets/fonts/Roboto-Light.ttf);
}
.p-button.p-button-success:enabled:focus,
.p-button.p-button-danger:enabled:focus,
.p-button.p-button-success:enabled:active,
.p-button.p-button-danger:enabled:active {
  box-shadow: none !important;
}
body {
  margin: 0;
  padding: 0;
  background-color: var(--surface-b);
}
html,
body {
  font-size: 1.2em;
  font-family: "roboto", sans-serif;
}
.p-component {
  font-family: "roboto", sans-serif;
}
</style>

<style>
.v-enter-active {
  transition: all 0.2s ease-in;
}

.v-enter-from {
  opacity: 0;
}
</style>
