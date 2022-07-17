<template>
  <Transition>
    <WaveSpinner v-if="isLoading" />
    <div v-else class="container">
      <div class="field-checkbox mt-2">
        <Checkbox id="binary" v-model="checkAll" :binary="true" @click="checkAllClicked" />
        <label for="binary">Alle Auswählen</label>
      </div>
      <hr style="color: var(--text-color)" class="my-3" />
      <div v-for="order in orders" :key="order.id" class="field-checkbox">
        <Checkbox :id="order.id" name="order" :value="order.id" v-model="orderFilter" />
        <label :for="order.id">{{ order.order_item.description }}</label>
      </div>
      <div class="flex justify-content-end">
        <div class="text-right">
          <Button
            :disabled="applyFilterLoading"
            :loading="deleteFilterLoading"
            label="Löschen"
            icon="pi pi-times"
            class="p-button-danger mr-2 p-button-text"
            @click="deleteFilter"
          />
          <Button
            :disabled="deleteFilterLoading"
            :loading="applyFilterLoading"
            icon="pi pi-check"
            class="p-button-success p-button-rounded"
            @click="applyFilter"
          />
        </div>
      </div>
    </div>
  </Transition>
</template>

<script lang="ts">
import { defineComponent, inject, ref, watch } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import Checkbox from "primevue/checkbox";
import { convertToEur } from "@/utils";
import Button from "primevue/button";
import { filter } from "@/keys";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";

export default defineComponent({
  name: "FilterModal",
  components: { WaveSpinner, Checkbox, Button },
  props: { tableId: { type: Number, required: true } },
  emits: ["newFilter"],
  setup(props, { emit }) {
    const orders = ref<service_Order[]>([]);
    const orderFilter = inject(filter, ref());
    const isLoading = ref(false);
    const applyFilterLoading = ref(false);
    const deleteFilterLoading = ref(false);
    const checkAll = ref(false);

    function checkAllCheck() {
      if (orderFilter.value) checkAll.value = orderFilter.value.length === orders.value.length;
    }

    watch(orderFilter, () => checkAllCheck());
    getData();

    function getData() {
      isLoading.value = true;
      OrdersService.getOrders(props.tableId, false)
        .then((res) => {
          orders.value = res;
          checkAllCheck();
        })
        .finally(() => (isLoading.value = false));
    }

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

    function applyFilter() {
      applyFilterLoading.value = true;
      emit("newFilter");
    }

    function deleteFilter() {
      deleteFilterLoading.value = true;
      orderFilter.value = undefined;
      emit("newFilter");
    }

    return { orders, orderFilter, checkAll, checkAllClicked, convertToEur, deleteFilter, applyFilter, isLoading, applyFilterLoading, deleteFilterLoading };
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
