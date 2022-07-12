<template>
  <BaseItem v-if="order.order_item.item_type === itemType" paddingRight="1">
    <div class="flex flex-column overflow-hidden">
      <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis">{{ order.order_item.description }}</div>
      <div class="flex align-items-baseline justify-content-between">
        <div>
          <Badge severity="info" class="text-sm" :value="convertToEur(order.order_item.price)" />
          <Badge severity="warning" class="text-sm ml-2" v-if="order.order_item.price !== order.total" :value="convertToEur(order.total)" />
        </div>
        <div class="flex align-items-center">
          <Button :disabled="isDisabled" icon="pi pi-minus" class="p-button-rounded p-button-text p-button-danger" @click="$emit('decrementOrder', order)" />
          <div class="mx-1 font-bold">{{ order.order_count }}</div>
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
import { convertToEur, ItemType } from "@/utils";
import Button from "primevue/button";
import Badge from "primevue/badge";

export default defineComponent({
  name: "OrderEntry",
  components: { BaseItem, Button, Badge },
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
    isDisabled: { type: Boolean, default: false },
    itemType: { type: Number, required: true },
  },
  emits: ["decrementOrder", "incrementOrder"],
  setup() {
    return { convertToEur, ItemType };
  },
});
</script>
