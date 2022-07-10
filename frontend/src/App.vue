<template>
  <router-view></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>

<style lang="less">
@import "primevue/resources/themes/bootstrap4-light-blue/theme.css";

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
  margin: 0;
  padding: 0;
  background-color: var(--surface-a);
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
@media (min-width: 768px) {
  body {
    background-color: var(--surface-d);
  }
  .p-card {
    background: var(--surface-a) !important;
    box-shadow: 0 1px 10px rgba(0, 0, 0, 0.12), 0 4px 5px rgba(0, 0, 0, 0.14), 0 2px 4px -1px rgba(0, 0, 0, 0.2) !important;
    border-radius: 0.5rem !important;
  }
  .mobileShadow {
    padding-bottom: 0;
    margin-bottom: 0;
    box-shadow: none;
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
.bgHover:hover {
  background-color: var(--surface-c);
  border-radius: 0.5rem;
}
</style>
