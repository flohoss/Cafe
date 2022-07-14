<template>
  <div class="flex align-items-center">
    <div @click="!isDisabled && $emit('decrementOrder')" :style="{ color: isDisabled ? 'grey' : 'red' }" style="cursor: pointer">
      <i class="pi pi-minus"></i>
    </div>
    <div class="mx-2 font-bold">{{ order.order_count }}</div>
    <div @click="!isDisabled && $emit('incrementOrder')" :style="{ color: isDisabled ? 'grey' : 'green' }" style="cursor: pointer">
      <i class="pi pi-plus"></i>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";

export default defineComponent({
  name: "OrderAmountChange",
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
    isDisabled: { type: Boolean, default: false },
  },
  emits: ["incrementOrder", "decrementOrder"],
});
</script>
