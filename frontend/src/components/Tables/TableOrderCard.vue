<template>
  <div class="col-12 lg:col-6">
    <BaseItem bgColor="d">
      <div class="flex flex-column overflow-hidden">
        <div class="font-bold white-space-nowrap overflow-hidden text-overflow-ellipsis mb-1">{{ order.order_item.description }}</div>
        <div class="flex align-items-center justify-content-between">
          <div>
            <TheBadge> {{ convertToEur(order.order_item.price) }} </TheBadge>
            <TheBadge color="warning" v-if="order.order_item.price !== order.total" class="ml-2"> {{ convertToEur(order.total) }} </TheBadge>
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
import TheBadge from "@/components/UI/TheBadge.vue";

export default defineComponent({
  name: "TableOrderCard",
  components: { TheBadge, BaseItem },
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
    isDisabled: { type: Boolean, default: false },
  },
  emits: ["decrementOrder", "incrementOrder"],
  setup() {
    return { convertToEur, ItemType };
  },
});
</script>
