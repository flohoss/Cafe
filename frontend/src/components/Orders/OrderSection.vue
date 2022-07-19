<template>
  <div v-if="orders.length !== 0">
    <BaseToolbar :icon="detailedItemTypeIcon(itemType)" :title="detailedItemTypeString(itemType)" btnIcon="check" @click="checkAllOpenOrders" />
    <div class="grid">
      <OrderCard v-for="order in orders" v-bind:key="order.id" :order="order" :isDisabled="isDisabled" @orderDone="(o) => orderDone(o)" />
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, PropType, ref } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { detailedItemTypeIcon, detailedItemTypeString, errorToast, ItemType, moreThanOneMinuteAgo } from "@/utils";
import OrderCard from "@/components/Orders/OrderCard.vue";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import { useToast } from "primevue/usetoast";
import { disabled } from "@/keys";

export default defineComponent({
  name: "OrderSection",
  components: { OrderCard, BaseToolbar },
  props: {
    orders: { type: Object as PropType<service_Order[]>, required: true },
    itemType: { type: Number as PropType<ItemType>, required: true },
  },
  emits: ["filterOrders"],
  setup(props, { emit }) {
    const toast = useToast();
    const isDisabled = inject(disabled, ref(false));
    const collapseOrders = ref(true);
    const collapseIcon = computed(() => (collapseOrders.value ? "chevron-down" : "chevron-up"));

    function checkAllOpenOrders() {
      props.orders.forEach((order) => {
        if (moreThanOneMinuteAgo(order.updated_at)) orderDone(order);
      });
    }

    function orderDone(order: service_Order) {
      isDisabled.value = true;
      order.is_served = true;
      OrdersService.putOrders(order)
        .then(() => emit("filterOrders", order.id))
        .catch((err) => errorToast(toast, err.body.error))
        .finally(() => (isDisabled.value = false));
    }
    return { detailedItemTypeIcon, detailedItemTypeString, checkAllOpenOrders, orderDone, isDisabled, collapseIcon, collapseOrders };
  },
});
</script>
