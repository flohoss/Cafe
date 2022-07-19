<template>
  <div>
    <BaseToolbar :title="generalItemTypeString(type)" :icon="generalItemTypeIcon(type)" @click="$emit('openModal', type)" btnIcon="plus" />
    <div class="grid">
      <TableOrderCard v-for="order in OrdersForType" v-bind:key="order.id" :order="order">
        <div class="flex align-items-end">
          <OrderAmountChange :order="order" :isDisabled="isLoading" @incrementOrder="incrementOrder(order)" @decrementOrder="decrementOrder(order)" />
        </div>
      </TableOrderCard>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, PropType, ref } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { convertToEur, ItemType, generalItemTypeString, generalItemTypeIcon } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import OrderAmountChange from "@/components/Tables/OrderAmountChange.vue";
import { loading } from "@/keys";

export default defineComponent({
  name: "OverviewPerType",
  components: { TableOrderCard, BaseToolbar, OrderAmountChange },
  props: {
    orders: { type: Array as PropType<service_Order[]>, default: () => [] },
    type: { type: Array as PropType<number[]>, required: true },
  },
  emits: ["openModal", "getData"],
  setup(props, { emit }) {
    const OrdersForType = computed(() => props.orders.filter((order) => props.type.includes(order.order_item.item_type)));
    const isLoading = inject(loading, ref(false));

    function incrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.postOrders(order.order_item_id, order.table_id).finally(() => emit("getData"));
    }

    function decrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.deleteOrders(order.order_item_id, order.table_id).finally(() => emit("getData"));
    }

    return { OrdersForType, isLoading, convertToEur, ItemType, incrementOrder, decrementOrder, generalItemTypeIcon, generalItemTypeString };
  },
});
</script>
