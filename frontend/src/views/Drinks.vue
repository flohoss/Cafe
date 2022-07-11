<template>
  <BaseCard class="text-center">
    <OrderItemList :orderItems="drinks" />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import FoodList from "@/components/OrderItem/OrderItemList.vue";
import { OrderItemsService } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";

export default defineComponent({
  name: "DrinksView",
  components: { OrderItemList, BaseCard },
  setup() {
    const drinks = ref();
    OrderItemsService.getOrdersItems().then((res) => (drinks.value = res));
    return { drinks };
  },
});
</script>

<style></style>
