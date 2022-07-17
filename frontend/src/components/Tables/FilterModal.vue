<template>
  <div class="container">
    <div class="field-checkbox mt-2">
      <Checkbox id="binary" v-model="checkAll" :binary="true" @click="checkAllClicked" />
      <label for="binary">Alle Auswählen</label>
    </div>
    <Divider />
    <div v-for="order in orders" :key="order.id" class="field-checkbox">
      <Checkbox :id="order.id" name="order" :value="order.id" v-model="orderFilter" />
      <label :for="order.id">{{ order.order_item.description }}</label>
    </div>
    <div class="flex justify-content-end">
      <div class="text-right">
        <Button label="Löschen" icon="pi pi-times" class="p-button-danger mr-2 p-button-text" @click="deleteFilter" />
        <Button icon="pi pi-check" class="p-button-success p-button-rounded" @click="$emit('newFilter')" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, ref, watch } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import Checkbox from "primevue/checkbox";
import Divider from "primevue/divider";
import { convertToEur } from "@/utils";
import Button from "primevue/button";
import { filter } from "@/keys";

export default defineComponent({
  name: "FilterModal",
  components: { Checkbox, Divider, Button },
  props: { tableId: { type: Number, required: true } },
  emits: ["newFilter"],
  setup(props, { emit }) {
    const orders = ref<service_Order[]>([]);
    const orderFilter = inject(filter, ref());
    const checkAll = ref(false);

    function checkAllCheck() {
      if (orderFilter.value) checkAll.value = orderFilter.value.length === orders.value.length;
    }

    watch(orderFilter, () => checkAllCheck());

    OrdersService.getOrders(props.tableId, false).then((res) => {
      orders.value = res;
      checkAllCheck();
    });

    function setAllOrdersSelected() {
      const temp: number[] = [];
      orders.value.forEach((order) => order.id && temp.push(order.id));
      orderFilter.value = temp;
    }

    function checkAllClicked() {
      if (!orderFilter.value) {
        setAllOrdersSelected();
      } else if (orderFilter.value) {
        if (orderFilter.value.length === orders.value.length) {
          orderFilter.value = undefined;
        } else {
          setAllOrdersSelected();
        }
      }
    }

    function deleteFilter() {
      orderFilter.value = undefined;
      emit("newFilter");
    }

    return { orders, orderFilter, checkAll, checkAllClicked, convertToEur, deleteFilter };
  },
});
</script>

<style scoped>
.container {
  --bs-gutter-x: 0;
  --bs-gutter-y: 0;
  width: 100%;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  margin-right: auto;
  margin-left: auto;
}
@media (min-width: 576px) {
  .container {
    max-width: 540px;
  }
}
@media (min-width: 768px) {
  .container {
    max-width: 720px;
  }
}
@media (min-width: 992px) {
  .container {
    max-width: 960px;
  }
}
@media (min-width: 1200px) {
  .container {
    max-width: 1140px;
  }
}
@media (min-width: 1400px) {
  .container {
    max-width: 1320px;
  }
}
</style>
