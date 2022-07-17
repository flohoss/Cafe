<template>
  <div class="col-12 lg:col-6">
    <BaseItem bgColor="d">
      <div class="flex flex-column overflow-hidden">
        <div class="font-bold white-space-nowrap overflow-hidden text-overflow-ellipsis mb-1">{{ order.order_item.description }}</div>
        <div class="flex align-items-center justify-content-between">
          <div>
            <TheBadge> {{ convertToEur(order.order_item.price) }} </TheBadge>
            <TheBadge v-if="showTotal" class="ml-2" color="warning"> {{ convertToEur(order.total) }} </TheBadge>
          </div>
          <slot></slot>
        </div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, ItemType } from "@/utils";
import TheBadge from "@/components/UI/TheBadge.vue";

export default defineComponent({
  name: "TableOrderCard",
  components: { TheBadge, BaseItem },
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
  },
  emits: ["decrementOrder", "incrementOrder"],
  setup(props) {
    const showTotal = computed(() => props.order.order_item.price !== props.order.total);
    return { convertToEur, ItemType, showTotal };
  },
});
</script>
