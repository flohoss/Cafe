<template>
  <div class="col-12 p-card shadow-1 my-2 p-2 px-3 relative">
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseItem",
});
</script>
