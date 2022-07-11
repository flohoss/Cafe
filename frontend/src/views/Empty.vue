<template>
  <BaseCard class="text-center">
    <div class="p-card w-full p-4">Bis jetzt noch nichts...</div>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";

export default defineComponent({
  name: "EmptyView",
  components: { BaseCard },
  setup() {
    return {};
  },
});
</script>

<style></style>
