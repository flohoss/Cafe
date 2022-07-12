<template>
  <div class="container">
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseCard",
});
</script>

<style scoped>
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
