<template>
  <Toolbar class="border-0 shadow-1 my-2 p-2 bg-color">
    <template #start>
      <div class="flex align-items-center">
        <div class="font-bold text-2xl">{{ title }}</div>
        <div class="ml-3"><i :class="icon" style="font-size: 1.5rem"></i></div>
      </div>
    </template>
    <template #end>
      <Button v-if="btnIcon" :disabled="isLoading" :icon="'pi pi-' + btnIcon" class="p-button-success p-button-rounded" @click="$emit('click')" />
    </template>
  </Toolbar>
</template>

<script lang="ts">
import { defineComponent, inject, ref } from "vue";
import Toolbar from "primevue/toolbar";
import Button from "primevue/button";
import { loading } from "@/keys";

export default defineComponent({
  name: "BaseToolbar",
  components: { Toolbar, Button },
  emits: ["click"],
  props: {
    title: { type: String, default: "" },
    icon: { type: String, default: "" },
    btnIcon: { type: String, default: "" },
  },
  setup() {
    const isLoading = inject(loading, ref(false));
    return { isLoading };
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-a);
  color: var(--text-color);
}
</style>
