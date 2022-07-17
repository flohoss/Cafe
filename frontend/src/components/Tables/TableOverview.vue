<template>
  <BaseCard>
    <ConfirmDialog></ConfirmDialog>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <OverviewPerType :filter="!!orderFilter" :type="ItemType.Food" :orders="orders" @getData="getData" @openModal="addBeverage(ItemType.Food)" />
        <OverviewPerType :filter="!!orderFilter" :type="ItemType.Drink" :orders="orders" @getData="getData" @openModal="addBeverage(ItemType.Drink)" />
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

    <Sidebar v-model:visible="filterModal" :baseZIndex="10000" position="full">
      <CheckoutView :tableId="table" @newFilter="getData" />
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
        <Button :disabled="isDisabled" icon="pi pi-filter" class="p-button-rounded mr-1" @click="filterModal = true" />
        <Button :disabled="isDisabled" icon="pi pi-money-bill" class="p-button-danger p-button-rounded" @click="checkoutOrders" />
      </template>
    </BottomNavigation>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, provide, ref } from "vue";
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
import CheckoutView from "@/components/Tables/FilterModal.vue";
import ConfirmDialog from "primevue/confirmdialog";
import { useConfirm } from "primevue/useconfirm";
import { filter } from "@/keys";

export default defineComponent({
  name: "TableOverview",
  components: { CheckoutView, OverviewPerType, WaveSpinner, BottomNavigation, BaseCard, Button, Sidebar, Listbox, ConfirmDialog },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const confirm = useConfirm();
    const isLoading = ref(false);
    const isDisabled = ref(false);
    const newOrderModal = ref(false);
    const filterModal = ref(false);
    const store = useStore();
    const selectedOrder = ref();
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const orders = ref<service_Order[]>([]);
    const orderFilter = ref<number[]>();
    provide(filter, orderFilter);

    store.dispatch("getAllOrderItems");

    getData(true);

    function getData(initial = false) {
      initial && (isLoading.value = true);
      OrdersService.getOrders(table.value, true, orderFilter.value && orderFilter.value.toString())
        .then((res) => (orders.value = res))
        .finally(() => {
          updateTotal();
          resetValues();
        });
    }

    function resetValues() {
      newOrderModal.value = false;
      filterModal.value = false;
      selectedOrder.value = undefined;
      isLoading.value = false;
      isDisabled.value = false;
    }

    function updateTotal() {
      let temp = 0;
      orders.value.forEach((order) => (temp += order.total));
      total.value = temp;
    }

    function addBeverage(type: ItemType) {
      newOrderModal.value = true;
      options.value = orderItems.value.get(type);
    }

    function postOrder() {
      isDisabled.value = true;
      if (selectedOrder.value) {
        OrdersService.postOrders(selectedOrder.value, table.value).finally(() => getData());
      } else isLoading.value = false;
    }

    function checkoutOrders() {
      confirm.require({
        message: "Angezeigte Bestellungen abrechnen?",
        header: "Abrechnung",
        icon: "pi pi-info-circle",
        acceptClass: "p-button-danger",
        accept: () => {
          console.log("checkout");
        },
      });
    }

    return {
      newOrderModal,
      filterModal,
      orderFilter,
      selected: selectedOrder,
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
      checkoutOrders,
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
