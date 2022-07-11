<template>
  <BaseCard class="text-center">
    <OrderItemList
      :orderItems="currentOrderItems"
      :emptyOrderItem="emptyOrderItem"
      @orderItemChanged="(item) => orderItemChanged(item)"
      @orderItemDeleted="(item) => orderItemDeleted(item)"
      @orderItemCreated="(item) => orderItemCreated(item)"
    />
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";
import { stringify } from "ts-jest";
import { useStore } from "vuex";
import { ItemType } from "@/utils";

export default defineComponent({
  name: "ItemView",
  components: { OrderItemList, BaseCard },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const store = useStore();
    const orderItems = computed(() => store.getters.getOrderItems);
    const currentOrderItems = ref();
    const emptyOrderItem = reactive<service_OrderItem>({ description: "", item_type: 0, price: 0 });
    const intId = ref<ItemType>(parseInt(props.id));

    async function getData() {
      intId.value = parseInt(props.id);
      await store.dispatch("getOrderItems", intId.value);
      emptyOrderItem.item_type = intId.value;
      refreshMap();
    }

    function refreshMap() {
      currentOrderItems.value = orderItems.value.get(intId.value);
    }

    onMounted(() => getData());
    watch(props, () => getData());

    function orderItemChanged(item: service_OrderItem) {
      store.dispatch("updateOrderItem", item);
      refreshMap();
    }
    function orderItemDeleted(item: service_OrderItem) {
      store.dispatch("deleteOrderItem", item);
      refreshMap();
    }
    function orderItemCreated(item: service_OrderItem) {
      store.dispatch("addOrderItem", item);
      refreshMap();
    }

    return { currentOrderItems, orderItemChanged, orderItemDeleted, orderItemCreated, emptyOrderItem };
  },
});
</script>

<style></style>
