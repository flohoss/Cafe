<template>
  <div>
    <BaseToolbar v-if="checkout" :title="title" icon="fa-cheese" />
    <BaseToolbar v-else :title="title" icon="fa-cheese" @click="$emit('openModal', type)" btnIcon="plus" />
    <div class="grid">
      <TableOrderCard v-for="order in OrdersForType" v-bind:key="order.id" :order="order">
        <span v-if="checkout" class="font-bold">{{ convertToEur(order.total) }}</span>
        <OrderAmountChange v-else :order="order" :isDisabled="isDisabled" @incrementOrder="incrementOrder(order)" @decrementOrder="decrementOrder(order)" />
      </TableOrderCard>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { convertToEur, ItemType } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import OrderAmountChange from "@/components/Tables/OrderAmountChange.vue";

export default defineComponent({
  name: "OverviewPerType",
  components: { TableOrderCard, BaseToolbar, OrderAmountChange },
  props: {
    orders: { type: Array as PropType<service_Order[]>, default: () => [] },
    type: { type: Number, default: 0 },
    title: { type: String, default: "" },
    checkout: { type: Boolean, default: false },
  },
  emits: ["openModal", "getData"],
  setup(props, { emit }) {
    const OrdersForType = computed(() => props.orders.filter((order) => order.order_item.item_type === props.type));
    const isDisabled = ref(false);

    function getData() {
      emit("getData");
      isDisabled.value = false;
    }

    function incrementOrder(order: service_Order) {
      isDisabled.value = true;
      OrdersService.postOrders(order.order_item_id, order.table_id).finally(() => getData());
    }

    function decrementOrder(order: service_Order) {
      isDisabled.value = true;
      OrdersService.deleteOrders(order.order_item_id, order.table_id).finally(() => getData());
    }

    return { OrdersForType, isDisabled, convertToEur, ItemType, incrementOrder, decrementOrder };
  },
});
</script>
