<template>
  <BaseCard class="text-center">
    <OrderItemList :orderItems="food" />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import FoodList from "@/components/OrderItem/OrderItemList.vue";
import { OrderItemsService } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";

export default defineComponent({
  name: "FoodView",
  components: { OrderItemList, BaseCard },
  setup() {
    const food = ref();
    OrderItemsService.getOrdersItems().then((res) => (food.value = res));
    return { food };
  },
});
</script>

<style></style>
