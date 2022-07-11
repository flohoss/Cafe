<template>
  <BaseItem paddingRight="1">
    <div class="flex justify-content-between">
      <div class="flex flex-column justify-content-between">
        <div>{{ order.order_item.description }}</div>
        <div class="text-sm">Einzelpreis: {{ convertToEur(order.order_item.price) }}</div>
      </div>
      <div class="flex align-items-center">
        <div>
          <Button :disabled="isDisabled" icon="pi pi-minus" class="p-button-rounded p-button-text p-button-danger" @click="$emit('decrementOrder', order)" />
        </div>
        <div>{{ order.total }}</div>
        <div>
          <Button :disabled="isDisabled" icon="pi pi-plus" class="p-button-rounded p-button-text p-button-success" @click="$emit('incrementOrder', order)" />
        </div>
      </div>
    </div>
  </BaseItem>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur } from "@/utils";
import Button from "primevue/button";

export default defineComponent({
  name: "OrderEntry",
  components: { BaseItem, Button },
  props: { order: { type: Object as PropType<service_Order>, required: true }, isDisabled: { type: Boolean, default: false } },
  emits: ["decrementOrder", "incrementOrder"],
  setup() {
    return { convertToEur };
  },
});
</script>
