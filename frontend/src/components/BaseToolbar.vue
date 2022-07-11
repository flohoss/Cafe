<template>
  <Toolbar class="border-0 shadow-1 my-2 p-2 bg-color">
    <template #start>
      <slot name="start"></slot>
    </template>
    <template #end>
      <slot name="end"></slot>
    </template>
  </Toolbar>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Toolbar from "primevue/toolbar";

export default defineComponent({
  name: "BaseToolbar",
  components: { Toolbar },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-a);
  color: var(--text-color);
}
</style>
