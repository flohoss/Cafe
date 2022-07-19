<template>
  <SmallCard bgColor="d" :badgeTwo="order.total !== order.order_item.price">
    <template #description>{{ order.order_item.description }}</template>
    <template #badgeOne>{{ convertToEur(order.order_item.price) }}</template>
    <template #badgeTwo>{{ convertToEur(order.total) }}</template>
    <template #right>
      <slot></slot>
    </template>
  </SmallCard>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import { convertToEur, ItemType } from "@/utils";
import SmallCard from "@/components/UI/SmallCard.vue";

export default defineComponent({
  name: "TableOrderCard",
  components: { SmallCard },
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
