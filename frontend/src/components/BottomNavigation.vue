<template>
  <div class="absolute bottom-0 left-0 w-full">
    <div class="flex justify-content-between align-items-center border-round-xs p-2 m-2 bg-color shadow-1">
      <div><slot name="left"></slot></div>
      <div><slot name="middle"></slot></div>
      <div><slot name="right"></slot></div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BottomNavigation",
  setup() {
    return {};
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-a);
  color: var(--text-color);
}
</style>
