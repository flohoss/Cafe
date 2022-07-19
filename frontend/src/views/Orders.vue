<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <EmptyView v-else-if="orders.length === 0" message="Keine offenen Bestellungen" />
      <div v-else>
        <OrderSection :orders="foodOrders" :itemType="ItemType.Food" @filterOrders="(id) => filterOrder(id)" />
        <OrderSection :orders="hotOrders" :itemType="ItemType.HotDrink" @filterOrders="(id) => filterOrder(id)" />
        <OrderSection :orders="coldOrders" :itemType="ItemType.ColdDrink" @filterOrders="(id) => filterOrder(id)" :edit="false" />
      </div>
    </Transition>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, onUnmounted, provide, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrdersService, service_Order } from "@/services/openapi";
import { detailedItemTypeIcon, detailedItemTypeString, ItemType, NotifierType, WebSocketMsg } from "@/utils";
import { WEBSOCKET_ENDPOINT_URL } from "@/main";
import EmptyView from "@/views/Empty.vue";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import { disabled, loading } from "@/keys";
import OrderSection from "@/components/Orders/OrderSection.vue";

export default defineComponent({
  name: "OrderView",
  components: { OrderSection, EmptyView, BaseCard, WaveSpinner },
  setup() {
    const isLoading = ref(true);
    const isDisabled = ref(false);
    provide(disabled, isDisabled);
    provide(loading, isDisabled);
    const orders = ref<service_Order[]>([]);
    const foodOrders = computed(() => orders.value.filter((order) => order.order_item.item_type === ItemType.Food));
    const hotOrders = computed(() => orders.value.filter((order) => order.order_item.item_type === ItemType.HotDrink));
    const coldOrders = computed(() => orders.value.filter((order) => order.order_item.item_type === ItemType.ColdDrink));
    let ws = ref<WebSocket | null>(null);

    getData();
    function getData() {
      isLoading.value = true;
      OrdersService.getOrders()
        .then((res) => (orders.value = res))
        .finally(() => {
          isLoading.value = false;
          startWebsocket();
        });
    }
    onUnmounted(() => stopWebsocket());

    function filterOrder(id: number) {
      orders.value = orders.value.filter((old) => old.id !== id);
    }

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

    function handleWebsocketError() {
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

    return {
      orders,
      hotOrders,
      coldOrders,
      foodOrders,
      filterOrder,
      ItemType,
      isLoading,
      isDisabled,
      detailedItemTypeString,
      detailedItemTypeIcon,
    };
  },
});
</script>
