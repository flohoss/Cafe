<template>
  <div class="badge p-badge" :class="`p-badge-${color}`">
    <span :class="`text-${size}`"><slot></slot></span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "TheBadge",
  props: { color: { type: String, default: "primary" }, size: { type: String, default: "xs" } },
});
</script>

<style scoped>
.badge {
  --bs-badge-padding-x: 0.5em;
  --bs-badge-padding-y: 0.3em;
  --bs-badge-font-size: 0.75em;
  --bs-badge-font-weight: 700;
  --bs-badge-border-radius: 0.375rem;
  display: inline-block;
  padding: var(--bs-badge-padding-y) var(--bs-badge-padding-x);
  font-size: var(--bs-badge-font-size);
  font-weight: var(--bs-badge-font-weight);
  line-height: 1;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;
  border-radius: var(--bs-badge-border-radius, 0);
}
.badge:empty {
  display: none;
}
</style>
