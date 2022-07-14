<template>
  <div class="col-12 lg:col-6">
    <BaseItem class="relative" bgColor="d">{{ bill }}</BaseItem>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_Bill } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, ItemType } from "@/utils";

export default defineComponent({
  name: "BillCard",
  components: { BaseItem },
  props: {
    bill: { type: Object as PropType<service_Bill>, required: true },
  },
  setup() {
    return { convertToEur, ItemType };
  },
});
</script>
