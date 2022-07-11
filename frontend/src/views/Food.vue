<template>
  <BaseCard class="text-center">
    <OrderItemList :orderItems="food" @orderItemChanged="(item) => orderItemChanged(item)" @orderItemDeleted="(item) => orderItemDeleted(item)" />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";

export default defineComponent({
  name: "FoodView",
  components: { OrderItemList, BaseCard },
  setup() {
    const food = ref();
    OrderItemsService.getOrdersItems().then((res) => (food.value = res));

    function orderItemChanged(item: service_OrderItem) {
      food.value = food.value.map((origItem: service_OrderItem) => (origItem.id === item.id ? item : origItem));
    }
    function orderItemDeleted(item: service_OrderItem) {
      food.value = food.value.filter((origItem: service_OrderItem) => origItem.id !== item.id);
    }

    return { food, orderItemChanged, orderItemDeleted };
  },
});
</script>

<style></style>
