<template>
  <BaseItem>
    <div class="flex justify-content-between">
      <div>{{ orderItems.description }}</div>
      <div>{{ convertToEur(orderItems.price) }}</div>
    </div>
  </BaseItem>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_OrderItem } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur } from "@/utils";

export default defineComponent({
  name: "OrderItemCard",
  components: { BaseItem },
  props: { orderItems: { type: Object as PropType<service_OrderItem>, required: true } },
  setup() {
    return { convertToEur };
  },
});
</script>
