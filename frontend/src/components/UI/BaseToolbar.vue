<template>
  <Toolbar class="border-0 shadow-1 my-2 p-2 bg-color">
    <template #start>
      <font-awesome-icon :icon="icon" size="1x" style="width: 2.5rem" />
      <div class="font-bold">{{ title }}</div>
    </template>
    <template #end>
      <Button :disabled="isLoading" icon="pi pi-plus" class="p-button-success p-button-rounded" @click="$emit('click')" />
    </template>
  </Toolbar>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Toolbar from "primevue/toolbar";
import Button from "primevue/button";

export default defineComponent({
  name: "BaseToolbar",
  components: { Toolbar, Button },
  emits: ["click"],
  props: {
    title: { type: String, default: "" },
    icon: { type: String, default: "" },
    isLoading: { type: Boolean, default: false },
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-c);
  color: var(--text-color);
}
</style>
