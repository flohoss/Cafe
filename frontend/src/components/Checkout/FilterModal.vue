<template>
  <div class="m-3">
    <div class="field-checkbox mt-2">
      <Checkbox id="binary" v-model="checkAll" :binary="true" @click="checkAllClicked" />
      <label for="binary">Alle Auswählen</label>
    </div>
    <Divider />
    <div v-for="order in orders" :key="order.id" class="field-checkbox">
      <Checkbox :id="order.id" name="order" :value="order.id" v-model="selected" />
      <label :for="order.id">{{ order.order_item.description }}</label>
    </div>
    <div class="flex justify-content-end">
      <Button label="Anwenden" icon="pi pi-check" @click="$emit('newFilter', selected)" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, PropType, ref, watch } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import Checkbox from "primevue/checkbox";
import Divider from "primevue/divider";
import { convertToEur } from "@/utils";
import Button from "primevue/button";

export default defineComponent({
  name: "FilterModal",
  components: { Checkbox, Divider, Button },
  props: { tableId: { type: Number, default: 0 }, filter: { type: Array as PropType<number[]>, required: true } },
  emits: ["newFilter"],
  setup(props) {
    const orders = ref<service_Order[]>([]);
    const selected = ref<number[]>(props.filter);
    const checkAll = ref(setCheckAll());
    const total = ref(0);

    watch(selected, (newValue) => {
      checkAll.value = newValue.length === orders.value.length;
    });

    onMounted(() => {
      OrdersService.getOrders(props.tableId, false)
        .then((res) => {
          orders.value = res;
        })
        .finally(() => {
          props.filter.length === 0 && initSelectedArray();
          checkAll.value = setCheckAll();
        });
    });

    function setCheckAll() {
      return selected.value.length === orders.value.length;
    }

    function initSelectedArray() {
      orders.value.forEach((order) => order.id && selected.value.push(order.id));
    }

    function checkAllClicked() {
      checkAll.value ? (selected.value = []) : initSelectedArray();
    }

    return { orders, selected, checkAll, checkAllClicked, convertToEur, total };
  },
});
</script>

<style></style>
