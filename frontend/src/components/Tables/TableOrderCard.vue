<template>
  <div v-if="order.order_item.item_type === itemType" class="col-12 lg:col-6">
    <BaseItem bgColor="c">
      <div class="flex flex-column overflow-hidden">
        <div class="font-bold white-space-nowrap overflow-hidden text-overflow-ellipsis">{{ order.order_item.description }}</div>
        <div class="flex align-items-baseline justify-content-between">
          <div>
            <Badge severity="info" class="text-sm" :value="convertToEur(order.order_item.price)" />
            <Badge severity="warning" class="text-sm ml-2" v-if="order.order_item.price !== order.total" :value="convertToEur(order.total)" />
          </div>
          <div class="flex align-items-center">
            <div @click="!isDisabled && $emit('decrementOrder', order)" :style="{ color: isDisabled ? 'grey' : 'red' }" style="cursor: pointer">
              <i class="pi pi-minus"></i>
            </div>
            <div class="mx-2 font-bold">{{ order.order_count }}</div>
            <div @click="!isDisabled && $emit('incrementOrder', order)" :style="{ color: isDisabled ? 'grey' : 'green' }" style="cursor: pointer">
              <i class="pi pi-plus"></i>
            </div>
          </div>
        </div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, ItemType } from "@/utils";
import Badge from "primevue/badge";

export default defineComponent({
  name: "TableOrderCard",
  components: { BaseItem, Badge },
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
