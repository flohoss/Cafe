<template>
  <BaseCard>
    <EmptyView v-if="orders.length === 0" message="Keine offenen Bestellungen" />
    <div v-else>
      <BaseToolbar icon="fa-box-open" title="Offen" btnIcon="check" @click="checkAllOpenOrders" />
      <div class="grid">
        <OrderCard v-for="entry in orders" v-bind:key="entry.id" :order="entry" :isDisabled="isLoading" @orderDone="(order) => orderDone(order)" />
      </div>
    </div>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onUnmounted, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrdersService, service_Order } from "@/services/openapi";
import OrderCard from "@/components/Orders/OrderCard.vue";
import { errorToast, ItemType } from "@/utils";
import { WEBSOCKET_ENDPOINT_URL } from "@/main";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import { useToast } from "primevue/usetoast";
import EmptyView from "@/views/Empty.vue";

export default defineComponent({
  name: "OrderView",
  components: { EmptyView, BaseToolbar, OrderCard, BaseCard },
  setup() {
    const toast = useToast();
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
      OrdersService.putOrders(order)
        .then(() => {
          orders.value = orders.value.filter((oldOrder) => oldOrder.id !== order.id);
        })
        .catch((err) => {
          errorToast(toast, err.body.error);
        });
    }

    function checkAllOpenOrders() {
      orders.value.forEach((order) => {
        orderDone(order);
      });
    }

    return { orders, ItemType, isLoading, orderDone, checkAllOpenOrders };
  },
});
</script>
