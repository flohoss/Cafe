<template>
  <div
    class="p-card p-2 shadow-1"
    :style="`color: var(--text-color); background-color: var(--surface-${bgColor});border: ${border}`"
    :class="`pr-${paddingRight}`"
  >
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseItem",
  props: {
    paddingRight: { type: String, default: "2" },
    bgColor: { type: String, default: "a" },
    border: { type: String, default: "" },
  },
});
</script>
