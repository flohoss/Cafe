<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <EmptyView v-else-if="orders.length === 0" message="Keine offenen Bestellungen" />
      <div v-else>
        <BaseToolbar icon="fa-solid fa-box-open" title="Offen" btnIcon="check" @click="checkAllOpenOrders" />
        <div class="grid">
          <OrderCard v-for="entry in orders" v-bind:key="entry.id" :order="entry" :isDisabled="isDisabled" @orderDone="(order) => orderDone(order)" />
        </div>
      </div>
    </Transition>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onMounted, onUnmounted, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrdersService, service_Order } from "@/services/openapi";
import OrderCard from "@/components/Orders/OrderCard.vue";
import { errorToast, ItemType, NotifierType, WebSocketMsg } from "@/utils";
import { WEBSOCKET_ENDPOINT_URL } from "@/main";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import { useToast } from "primevue/usetoast";
import EmptyView from "@/views/Empty.vue";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";

export default defineComponent({
  name: "OrderView",
  components: { EmptyView, BaseToolbar, OrderCard, BaseCard, WaveSpinner },
  setup() {
    const toast = useToast();
    const isLoading = ref(true);
    const isDisabled = ref(false);
    const orders = ref<service_Order[]>([]);
    let ws = ref<WebSocket | null>(null);

    onMounted(() => {
      isLoading.value = true;
      OrdersService.getOrders()
        .then((res) => (orders.value = res))
        .finally(() => {
          isLoading.value = false;
          startWebsocket();
        });
    });
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
      isDisabled.value = true;
      const messageEvent = evt as MessageEvent;
      const webSocketMsg: WebSocketMsg = JSON.parse(messageEvent.data);
      if (webSocketMsg.type === NotifierType.Create) {
        orders.value.push(webSocketMsg.payload[0]);
      } else if (webSocketMsg.type === NotifierType.Delete) {
        orders.value = orders.value.filter((o) => o.id !== webSocketMsg.payload[0].id);
      } else if (webSocketMsg.type === NotifierType.DeleteAll) {
        webSocketMsg.payload.forEach((obj) => {
          orders.value = orders.value.filter((o) => o.id !== obj.id);
        });
      }
      sortOrders();
      isDisabled.value = false;
    }

    function sortOrders() {
      orders.value.sort((a, b) => (a.updated_at && b.updated_at ? a.updated_at - b.updated_at : 0));
    }

    function orderDone(order: service_Order) {
      isDisabled.value = true;
      order.is_served = true;
      OrdersService.putOrders(order)
        .then(() => (orders.value = orders.value.filter((oldOrder) => oldOrder.id !== order.id)))
        .catch((err) => errorToast(toast, err.body.error))
        .finally(() => (isDisabled.value = false));
    }

    function checkAllOpenOrders() {
      orders.value.forEach((order) => {
        orderDone(order);
      });
    }

    return { orders, ItemType, isLoading, isDisabled, orderDone, checkAllOpenOrders };
  },
});
</script>
