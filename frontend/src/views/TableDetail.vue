<template>
  <BaseCard>
    <BaseToolbar title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" />
    <BaseItem>{{ foodOrders }}</BaseItem>
    <BaseToolbar title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" />
    <BaseItem>{{ drinkOrders }}</BaseItem>
    <Dialog v-model:visible="modal" :modal="true" :showHeader="false">
      <div class="p-fluid">
        <Listbox
          class="mt-1"
          v-model="selected"
          :options="options"
          :filter="true"
          optionLabel="description"
          dataKey="id"
          optionValue="id"
          listStyle="max-height:70vh"
          filterPlaceholder="Auswählen"
        />
      </div>
      <div class="flex justify-content-end mt-4">
        <Button icon="pi pi-times" class="p-button-text p-button-rounded p-button-secondary mr-2" @click="modal = false" />
        <Button icon="pi pi-check" class="p-button-rounded p-button-success" @click="addOrder" />
      </div>
    </Dialog>

    <BottomNavigation>
      <template #left>
        <router-link to="/tables">
          <Button icon="pi pi-arrow-left" class="p-button-rounded" />
        </router-link>
      </template>
      <template #middle>
        <div class="flex flex-column align-items-center">
          <div class="text-sm">Tisch {{ table.id }}</div>
          <div class="font-bold">{{ convertToEur(table.total) }}</div>
        </div>
      </template>
      <template #right>
        <router-link to="/bills">
          <Button icon="pi pi-money-bill" class="p-button-danger p-button-rounded" />
        </router-link>
      </template>
    </BottomNavigation>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { useStore } from "vuex";
import { OrdersService, service_Order, service_Table } from "@/services/openapi";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import Button from "primevue/button";
import { convertToEur, ItemType } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import Listbox from "primevue/listbox";
import Dialog from "primevue/dialog";
import BaseItem from "@/components/UI/BaseItem.vue";

export default defineComponent({
  name: "TableDetail",
  components: { BaseItem, BaseToolbar, BottomNavigation, BaseCard, Button, Dialog, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const modal = ref(false);
    const store = useStore();
    const selected = ref();
    const tables = computed(() => store.getters.getTables);
    const table = tables.value.find((table: service_Table) => table.id === parseInt(props.id));
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const drinkOrders = ref();
    const foodOrders = ref();

    getData();
    async function getData() {
      await store.dispatch("getAllOrderItems");
      drinkOrders.value = await OrdersService.getOrders(table.id, ItemType.Drink);
      foodOrders.value = await OrdersService.getOrders(table.id, ItemType.Food);
    }

    async function addBeverage(type: ItemType) {
      modal.value = true;
      options.value = orderItems.value.get(type);
    }

    function addOrder() {
      OrdersService.postOrders(selected.value, table.id).finally(() => {
        modal.value = false;
        selected.value = undefined;
      });
    }

    return { modal, selected, options, table, isLoading, convertToEur, addBeverage, ItemType, addOrder, foodOrders, drinkOrders };
  },
});
</script>

<style>
.p-dialog-content {
  padding: 0.5rem !important;
}
</style>
