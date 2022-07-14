<template>
  <BaseCard>
    <BaseItem bgColor="c">
      <div v-for="order in orders" v-bind:key="order.id">{{ order.id }}</div>

      <Sidebar v-model:visible="checkoutModal" :baseZIndex="10000" position="full"> </Sidebar>
    </BaseItem>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BaseItem from "@/components/UI/BaseItem.vue";
import Sidebar from "primevue/sidebar";
import { OrdersService } from "@/services/openapi";

export default defineComponent({
  name: "CheckoutView",
  components: { BaseItem, BaseCard, Sidebar },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const orders = ref();
    const checkoutModal = ref(false);
    const table = computed(() => parseInt(props.id));
    OrdersService.getOrders(table.value, false).then((res) => (orders.value = res));

    return { checkoutModal, orders };
  },
});
</script>

<style></style>
