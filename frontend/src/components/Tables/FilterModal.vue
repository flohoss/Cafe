<template>
  <div class="m-3">
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
        <Button label="Anwenden" icon="pi pi-check" @click="$emit('newFilter')" />
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

    return { orders, orderFilter, checkAll, checkAllClicked, convertToEur };
  },
});
</script>

<style></style>
