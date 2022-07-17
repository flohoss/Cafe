<template>
  <div>
    <BaseToolbar :title="title" :icon="icon" @click="$emit('openModal')" :btnIcon="!filter ? 'plus' : ''" />
    <div class="grid">
      <TableOrderCard v-for="order in OrdersForType" v-bind:key="order.id" :order="order">
        <TheBadge v-if="filter" size="md" color="warning"> {{ order.order_count }}x </TheBadge>
        <OrderAmountChange v-else :order="order" :isDisabled="isLoading" @incrementOrder="incrementOrder(order)" @decrementOrder="decrementOrder(order)" />
      </TableOrderCard>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, PropType, ref } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { convertToEur, ItemType, ItemTypeIcon, ItemTypeString } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import OrderAmountChange from "@/components/Tables/OrderAmountChange.vue";
import TheBadge from "@/components/UI/TheBadge.vue";
import { loading } from "@/keys";

export default defineComponent({
  name: "OverviewPerType",
  components: { TableOrderCard, BaseToolbar, OrderAmountChange, TheBadge },
  props: {
    orders: { type: Array as PropType<service_Order[]>, default: () => [] },
    type: { type: Number, required: true },
    filter: { type: Boolean, required: true },
  },
  emits: ["openModal", "getData"],
  setup(props, { emit }) {
    const OrdersForType = computed(() => props.orders.filter((order) => order.order_item.item_type === props.type));
    const isLoading = inject(loading, ref(false));
    const icon = computed(() => ItemTypeIcon(props.type));
    const title = computed(() => ItemTypeString(props.type));

    function incrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.postOrders(order.order_item_id, order.table_id).finally(() => emit("getData"));
    }

    function decrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.deleteOrders(order.order_item_id, order.table_id).finally(() => emit("getData"));
    }

    return { OrdersForType, isLoading, convertToEur, ItemType, incrementOrder, decrementOrder, icon, title };
  },
});
</script>
