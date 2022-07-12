<template>
  <BaseCard>
    <div class="p-card w-full p-4 text-center">{{ message }}</div>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";

export default defineComponent({
  name: "EmptyView",
  components: { BaseCard },
  props: { message: { type: String, default: "Bis jetzt noch nichts..." } },
});
</script>

<style></style>
