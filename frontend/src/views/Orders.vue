<template>
  <BaseCard>
    <div v-if="orders.length === 0" class="p-card w-full p-4 text-center">Keine offenen Bestellungen</div>
    <OrderEntry v-else v-for="entry in orders" v-bind:key="entry.id" :order="entry" :isDisabled="isLoading" @orderDone="(order) => orderDone(order)" />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onUnmounted, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrdersService, service_Order } from "@/services/openapi";
import OrderEntry from "@/components/Order/OrderEntry.vue";
import { ItemType } from "@/utils";
import { WEBSOCKET_ENDPOINT_URL } from "@/main";

export default defineComponent({
  name: "OrderView",
  components: { OrderEntry, BaseCard },
  setup() {
    const isLoading = ref(true);
    const orders = ref<service_Order[]>([]);

    OrdersService.getOrders()
      .then((res) => (orders.value = res))
      .finally(() => (isLoading.value = false));
    let ws = ref<WebSocket | null>(null);
    startWebsocket();
    onUnmounted(() => stopWebsocket());

    function startWebsocket() {
      ws.value = new WebSocket(WEBSOCKET_ENDPOINT_URL);
      ws.value.addEventListener("message", parseWebsocket);
      ws.value.addEventListener("error", handleWebsocketError);
    }

    function stopWebsocket() {
      if (ws.value) {
        ws.value.removeEventListener("message", parseWebsocket);
        ws.value.removeEventListener("error", handleWebsocketError);
        ws.value.close();
      }
    }

    async function handleWebsocketError() {
      stopWebsocket();
      setTimeout(() => {
        startWebsocket();
      }, 1000);
    }

    function parseWebsocket(evt: Event) {
      isLoading.value = true;
      const messageEvent = evt as MessageEvent;
      const order = JSON.parse(messageEvent.data);
      orders.value.push(order);
      sortOrders();
      isLoading.value = false;
    }

    function sortOrders() {
      orders.value.sort((a, b) => (a.updated_at && b.updated_at ? a.updated_at - b.updated_at : 0));
    }

    function orderDone(order: service_Order) {
      order.is_served = true;
      OrdersService.putOrders(order).then(() => {
        orders.value = orders.value.filter((oldOrder) => oldOrder.id !== order.id);
      });
    }
    return { orders, ItemType, isLoading, orderDone };
  },
});
</script>

<style></style>
