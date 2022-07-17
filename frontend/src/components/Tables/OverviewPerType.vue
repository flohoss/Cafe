<template>
  <div>
    <BaseToolbar :title="title" :icon="icon" @click="$emit('openModal')" btnIcon="plus" />
    <div class="grid">
      <TableOrderCard v-for="order in OrdersForType" v-bind:key="order.id" :order="order">
        <OrderAmountChange :order="order" :isDisabled="isDisabled" @incrementOrder="incrementOrder(order)" @decrementOrder="decrementOrder(order)" />
      </TableOrderCard>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { convertToEur, ItemType, ItemTypeIcon, ItemTypeString } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import OrderAmountChange from "@/components/Tables/OrderAmountChange.vue";

export default defineComponent({
  name: "OverviewPerType",
  components: { TableOrderCard, BaseToolbar, OrderAmountChange },
  props: {
    orders: { type: Array as PropType<service_Order[]>, default: () => [] },
    type: { type: Number, default: 0 },
  },
  emits: ["openModal", "getData"],
  setup(props, { emit }) {
    const OrdersForType = computed(() => props.orders.filter((order) => order.order_item.item_type === props.type));
    const isDisabled = ref(false);
    const icon = computed(() => ItemTypeIcon(props.type));
    const title = computed(() => ItemTypeString(props.type));

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

    return { OrdersForType, isDisabled, convertToEur, ItemType, incrementOrder, decrementOrder, icon, title };
  },
});
</script>
