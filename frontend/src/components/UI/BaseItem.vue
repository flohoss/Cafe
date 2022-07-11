<template>
  <div class="col-12 p-card shadow-1 my-2 p-2 pl-3 relative" :class="'pr-' + paddingRight">
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseItem",
  props: { paddingRight: { type: String, default: "3" } },
});
</script>
