<template>
  <div class="grid m-1 mx-2 md:m-3 lg:m-4 p-card border-round-lg relative p-0 md:p-2">
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
.p-card {
  background: transparent !important;
  box-shadow: none !important;
  border-radius: 0 !important;
}
@media (min-width: 768px) {
  .p-card {
    background: var(--surface-a) !important;
    box-shadow: 0 1px 10px rgba(0, 0, 0, 0.12), 0 4px 5px rgba(0, 0, 0, 0.14), 0 2px 4px -1px rgba(0, 0, 0, 0.2) !important;
    border-radius: 0.5rem !important;
  }
}
</style>
