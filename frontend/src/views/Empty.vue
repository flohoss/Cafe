<template>
  <BaseCard>
    <BaseItem bgColor="c">
      <div class="p-3 text-center">{{ message }}</div>
    </BaseItem>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BaseItem from "@/components/UI/BaseItem.vue";

export default defineComponent({
  name: "EmptyView",
  components: { BaseItem, BaseCard },
  props: { message: { type: String, default: "Bis jetzt noch nichts..." } },
});
</script>

<style></style>
