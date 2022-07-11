<template>
  <BaseCard class="text-center">
    <OrderItemList :orderItems="food" @orderItemChanged="(item) => orderItemChanged(item)" @orderItemDeleted="(item) => orderItemDeleted(item)" />
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import OrderItemList from "@/components/OrderItem/OrderItemList.vue";

export default defineComponent({
  name: "ItemView",
  components: { OrderItemList, BaseCard },
  props: { id: { type: Number, default: 0 } },
  setup(props) {
    const food = ref();

    function getData() {
      OrderItemsService.getOrdersItems(props.id).then((res) => (food.value = res));
    }

    onMounted(() => getData());
    watch(props, () => getData());

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
