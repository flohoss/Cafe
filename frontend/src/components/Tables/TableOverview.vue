<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <OverviewPerType :type="ItemType.Food" :orders="orders" @getData="getData" @openModal="(type) => addBeverage(type)" />
        <OverviewPerType :type="ItemType.Drink" :orders="orders" @getData="getData" @openModal="(type) => addBeverage(type)" />
        <div class="h-4rem"></div>
      </div>
    </Transition>

    <Sidebar v-model:visible="newOrderModal" :baseZIndex="10000" position="full">
      <div class="p-fluid">
        <Listbox
          v-model="selected"
          :options="options"
          :filter="true"
          optionLabel="description"
          dataKey="id"
          optionValue="id"
          listStyle="max-height:65vh"
          filterPlaceholder="Suchen"
        />
      </div>
      <div class="flex justify-content-end mt-4">
        <Button :loading="isLoading" label="Speichern" icon="pi pi-check" class="p-button p-button-success mr-3" @click="postOrder" />
      </div>
    </Sidebar>

    <BottomNavigation>
      <template #left>
        <router-link :to="{ name: 'Tables' }" class="no-underline">
          <Button :disabled="isDisabled" icon="pi pi-arrow-left" class="p-button-rounded" />
        </router-link>
      </template>
      <template #middle>
        <div class="flex flex-column align-items-center">
          <div class="text-sm">Tisch {{ table }}</div>
          <div class="font-bold">{{ convertToEur(total) }}</div>
        </div>
      </template>
      <template #right>
        <router-link :to="{ name: 'Checkout' }" class="no-underline">
          <Button :disabled="isDisabled" icon="pi pi-money-bill" class="p-button-danger p-button-rounded" />
        </router-link>
      </template>
    </BottomNavigation>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { useStore } from "vuex";
import { OrdersService, service_Order } from "@/services/openapi";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import Button from "primevue/button";
import { convertToEur, ItemType } from "@/utils";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import Sidebar from "primevue/sidebar";
import Listbox from "primevue/listbox";
import OverviewPerType from "@/components/Tables/OverviewPerType.vue";

export default defineComponent({
  name: "TableOverview",
  components: { OverviewPerType, WaveSpinner, BottomNavigation, BaseCard, Button, Sidebar, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const isDisabled = ref(false);
    const newOrderModal = ref(false);
    const store = useStore();
    const selected = ref();
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const orders = ref<service_Order[]>([]);

    store.dispatch("getAllOrderItems");

    getData(true);

    function getData(initial = false) {
      if (initial) isLoading.value = true;
      OrdersService.getOrders(table.value)
        .then((res) => (orders.value = res))
        .finally(() => {
          updateTotal();
          resetValues();
        });
    }

    function resetValues() {
      newOrderModal.value = false;
      selected.value = undefined;
      isLoading.value = false;
      isDisabled.value = false;
    }

    function updateTotal() {
      let temp = 0;
      orders.value.forEach((order) => (temp += order.total));
      total.value = temp;
    }

    async function addBeverage(type: ItemType) {
      newOrderModal.value = true;
      options.value = orderItems.value.get(type);
    }

    function postOrder() {
      isDisabled.value = true;
      if (selected.value) {
        OrdersService.postOrders(selected.value, table.value).finally(() => getData());
      } else isLoading.value = false;
    }

    return {
      newOrderModal,
      selected,
      options,
      table,
      total,
      isDisabled,
      isLoading,
      convertToEur,
      addBeverage,
      ItemType,
      postOrder,
      orders,
      getData,
    };
  },
});
</script>

<style>
.p-sidebar-content {
  margin: 0 !important;
  padding: 0 !important;
}

.p-listbox {
  border: 0 !important;
}
</style>
