<template>
  <BaseCard class="text-center">
    <OrderItemList
      :orderItems="food"
      :emptyOrderItem="emptyOrderItem"
      @orderItemChanged="(item) => orderItemChanged(item)"
      @orderItemDeleted="(item) => orderItemDeleted(item)"
      @orderItemCreated="(item) => orderItemCreated(item)"
    />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";
import { stringify } from "ts-jest";

export default defineComponent({
  name: "ItemView",
  components: { OrderItemList, BaseCard },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const food = ref();
    const emptyOrderItem = reactive<service_OrderItem>({ description: "", item_type: 0, price: 0 });

    function getData() {
      const intId = parseInt(props.id);
      OrderItemsService.getOrdersItems(intId).then((res) => (food.value = res));
      emptyOrderItem.item_type = intId;
    }

    onMounted(() => getData());
    watch(props, () => getData());

    function orderItemChanged(item: service_OrderItem) {
      food.value = food.value.map((origItem: service_OrderItem) => (origItem.id === item.id ? item : origItem));
    }
    function orderItemDeleted(item: service_OrderItem) {
      food.value = food.value.filter((origItem: service_OrderItem) => origItem.id !== item.id);
    }
    function orderItemCreated(item: service_OrderItem) {
      food.value.push(item);
    }

    return { food, orderItemChanged, orderItemDeleted, orderItemCreated, emptyOrderItem };
  },
});
</script>

<style></style>
