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
      <div class="text-right">
        <Button label="Anwenden" icon="pi pi-check" @click="applyFilter" />
        <div v-if="error" class="text-red-500 mt-2 text-sm">{{ error }}</div>
      </div>
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
  props: { tableId: { type: Number, default: 0 } },
  emits: ["newFilter"],
  setup(props, { emit }) {
    const orders = ref<service_Order[]>([]);
    const selected = ref<number[]>([]);
    const checkAll = ref(false);
    const error = ref("");

    watch(selected, () => {
      checkAll.value = sameAmountSelected();
    });

    onMounted(() => {
      OrdersService.getOrders(props.tableId, false)
        .then((res) => {
          orders.value = res;
        })
        .finally(() => {
          initSelectedArray();
          checkAll.value = sameAmountSelected();
        });
    });

    function sameAmountSelected() {
      return selected.value.length === orders.value.length;
    }

    function initSelectedArray() {
      orders.value.forEach((order) => order.id && selected.value.push(order.id));
    }

    function checkAllClicked() {
      checkAll.value ? (selected.value = []) : initSelectedArray();
    }

    function applyFilter() {
      if (sameAmountSelected()) {
        emit("newFilter", []);
        error.value = "";
      } else if (selected.value.length !== 0) {
        emit("newFilter", selected.value);
        error.value = "";
      } else error.value = "Bitte mindestens 1 Artikel wählen";
    }

    return { orders, selected, checkAll, checkAllClicked, convertToEur, applyFilter, error };
  },
});
</script>

<style></style>
