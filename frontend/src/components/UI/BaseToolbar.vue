<template>
  <Toolbar class="border-0 shadow-1 my-2 p-2 pr-3 bg-color">
    <template #start>
      <font-awesome-icon :icon="icon" size="1x" style="width: 2.5rem" />
      <div class="font-bold">{{ title }}</div>
    </template>
    <template #end>
      <Button v-if="btnIcon" :disabled="isDisabled" :icon="'pi pi-' + btnIcon" class="p-button-success p-button-rounded" @click="$emit('click')" />
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
    isDisabled: { type: Boolean, default: false },
    btnIcon: { type: String, default: "" },
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-a);
  color: var(--text-color);
}
</style>
