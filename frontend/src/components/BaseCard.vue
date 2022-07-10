<template>
  <div class="grid mx-3">
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseCard",
});
</script>

<style scoped></style>
