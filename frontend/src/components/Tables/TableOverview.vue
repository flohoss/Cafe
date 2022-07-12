<template>
  <BaseCard>
    <BaseToolbar :isDisabled="isLoading" title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" btnIcon="plus" />
    <div class="grid">
      <TableOrderCard
        v-for="entry in orders"
        v-bind:key="entry.id"
        :itemType="ItemType.Food"
        :order="entry"
        :isDisabled="isLoading"
        @incrementOrder="(order) => incrementOrder(order)"
        @decrementOrder="(order) => decrementOrder(order)"
      />
    </div>
    <BaseToolbar :isDisabled="isLoading" title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" btnIcon="plus" />
    <div class="grid">
      <TableOrderCard
        v-for="entry in orders"
        v-bind:key="entry.id"
        :itemType="ItemType.Drink"
        :order="entry"
        :isDisabled="isLoading"
        @incrementOrder="(order) => incrementOrder(order)"
        @decrementOrder="(order) => decrementOrder(order)"
      />
    </div>
    <Sidebar v-model:visible="modal" :modal="true" :baseZIndex="10000" position="full">
      <div class="p-fluid">
        <Listbox
          v-model="selected"
          :options="options"
          :filter="true"
          optionLabel="description"
          dataKey="id"
          optionValue="id"
          listStyle="max-height:80vh"
          filterPlaceholder="Suchen"
        />
      </div>
      <div class="flex justify-content-end mt-4">
        <Button label="Speichern" icon="pi pi-check" class="p-button p-button-success mr-3" @click="postOrder" />
      </div>
    </Sidebar>

    <BottomNavigation>
      <template #left>
        <router-link to="/tables" class="no-underline">
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
        <router-link to="/bills" class="no-underline">
          <Button :disabled="isLoading" icon="pi pi-money-bill" class="p-button-danger p-button-rounded" />
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
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import Listbox from "primevue/listbox";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import Sidebar from "primevue/sidebar";

export default defineComponent({
  name: "TableOverview",
  components: { TableOrderCard, BaseToolbar, BottomNavigation, BaseCard, Button, Sidebar, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const modal = ref(false);
    const store = useStore();
    const selected = ref();
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const orders = ref<service_Order[]>([]);
    const currentItemType = ref(0);

    store.dispatch("getAllOrderItems");

    getData();
    async function getData() {
      orders.value = await OrdersService.getOrders(table.value);
      updateTotal();
    }

    function updateTotal() {
      let temp = 0;
      orders.value.forEach((order) => (temp += order.total));
      total.value = temp;
    }

    async function addBeverage(type: ItemType) {
      modal.value = true;
      currentItemType.value = type;
      options.value = orderItems.value.get(type);
    }

    async function postOrder() {
      OrdersService.postOrders(selected.value, table.value);
      await getData();
      modal.value = false;
      selected.value = undefined;
    }

    async function incrementOrder(order: service_Order) {
      isLoading.value = true;
      await OrdersService.postOrders(order.order_item_id, order.table_id);
      await getData();
      isLoading.value = false;
    }

    async function decrementOrder(order: service_Order) {
      isLoading.value = true;
      await OrdersService.deleteOrders(order.order_item_id, order.table_id);
      await getData();
      isLoading.value = false;
    }

    return {
      modal,
      selected,
      options,
      table,
      total,
      isLoading,
      convertToEur,
      addBeverage,
      ItemType,
      postOrder,
      orders,
      incrementOrder,
      decrementOrder,
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
