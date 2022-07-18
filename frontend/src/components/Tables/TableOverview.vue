<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="initialLoading" />
      <div v-else>
        <OverviewPerType :type="[ItemType.Food]" :orders="orders" @getData="getData" @openModal="(t) => addBeverage(t)" />
        <OverviewPerType :type="[ItemType.ColdDrink, ItemType.HotDrink]" :orders="orders" @getData="getData" @openModal="(t) => addBeverage(t)" />
        <div class="h-4rem"></div>

        <BottomNavigation>
          <template #left>
            <router-link :to="{ name: 'Tables' }" class="no-underline">
              <Button :disabled="isLoading" icon="pi pi-arrow-left" class="p-button-rounded" />
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
              <Button :disabled="isLoading || orders.length === 0" icon="pi pi-money-bill" class="p-button-danger p-button-rounded"
            /></router-link>
          </template>
        </BottomNavigation>
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
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, provide, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { useStore } from "vuex";
import { OrdersService, service_Order, service_OrderItem } from "@/services/openapi";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import Button from "primevue/button";
import { convertToEur, ItemType } from "@/utils";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import Sidebar from "primevue/sidebar";
import Listbox from "primevue/listbox";
import OverviewPerType from "@/components/Tables/OverviewPerType.vue";
import { loading } from "@/keys";

export default defineComponent({
  name: "TableOverview",
  components: { OverviewPerType, WaveSpinner, BottomNavigation, BaseCard, Button, Sidebar, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const initialLoading = ref(false);
    const isLoading = ref(false);
    provide(loading, isLoading);
    const newOrderModal = ref(false);
    const store = useStore();
    const selectedOrder = ref();
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const orders = ref<service_Order[]>([]);

    store.dispatch("getAllOrderItems");

    getData(true);

    function getData(initial = false) {
      initial && (initialLoading.value = true);
      OrdersService.getOrders(table.value, true)
        .then((res) => (orders.value = res))
        .finally(() => {
          updateTotal();
          resetValues();
        });
    }

    function resetValues() {
      newOrderModal.value = false;
      selectedOrder.value = undefined;
      isLoading.value = false;
      initialLoading.value = false;
    }

    function updateTotal() {
      let temp = 0;
      orders.value.forEach((order) => (temp += order.total));
      total.value = temp;
    }

    function addBeverage(itemType: ItemType[]) {
      newOrderModal.value = true;
      options.value = [];
      itemType.forEach((type) => {
        options.value = options.value.concat(orderItems.value.get(type));
      });
      options.value.sort((a: service_OrderItem, b: service_OrderItem) => {
        let x = a.description.toLowerCase();
        let y = b.description.toLowerCase();
        if (x < y) return -1;
        if (x > y) return 1;
        return 0;
      });
    }

    function postOrder() {
      isLoading.value = true;
      if (selectedOrder.value) {
        OrdersService.postOrders(selectedOrder.value, table.value).finally(() => getData());
      } else isLoading.value = false;
    }

    return {
      initialLoading,
      isLoading,
      newOrderModal,
      selected: selectedOrder,
      options,
      table,
      total,
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
