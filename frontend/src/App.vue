<template>
  <TheNavigation @logout="logout" />
  <router-view></router-view>
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
@import "primevue/resources/themes/bootstrap4-light-blue/theme.css";

@font-face {
  font-family: "roboto";
  src: url(@/assets/fonts/Roboto-Light.ttf);
}

.router-link-active {
}
.overflow-ellipsis {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
body {
  margin: 0;
  padding: 0;
  background-color: var(--surface-b);
}
html,
body {
  font-size: 1em;
  font-family: "roboto", sans-serif !important;
}
</style>
