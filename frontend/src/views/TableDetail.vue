<template>
  <BaseCard>
    <BaseToolbar title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" />
    <OrderEntry v-for="entry in foodOrders" v-bind:key="entry.id" :order="entry" />
    <BaseToolbar title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" />
    <OrderEntry v-for="entry in drinkOrders" v-bind:key="entry.id" :order="entry" />
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
          filterPlaceholder="Suchen"
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
import OrderEntry from "@/components/Order/OrderEntry.vue";

export default defineComponent({
  name: "TableDetail",
  components: { OrderEntry, BaseToolbar, BottomNavigation, BaseCard, Button, Dialog, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const modal = ref(false);
    const store = useStore();
    const selected = ref();
    const tables = computed(() => store.getters.getTables);
    const table = ref(tables.value.find((table: service_Table) => table.id === parseInt(props.id)));
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const drinkOrders = ref<service_Order[]>([]);
    const foodOrders = ref<service_Order[]>([]);
    const currentItemType = ref(0);

    function sortOrders(orders: service_Order[]) {
      orders.sort((a, b) => {
        let fa = a.order_item.description.toLowerCase(),
          fb = b.order_item.description.toLowerCase();

        if (fa < fb) {
          return -1;
        }
        if (fa > fb) {
          return 1;
        }
        return 0;
      });
    }

    getData();
    async function getData() {
      await store.dispatch("getAllOrderItems");
      drinkOrders.value = await OrdersService.getOrders(table.value.id, ItemType.Drink);
      foodOrders.value = await OrdersService.getOrders(table.value.id, ItemType.Food);
    }

    async function addBeverage(type: ItemType) {
      modal.value = true;
      currentItemType.value = type;
      options.value = orderItems.value.get(type);
    }

    function addOrder() {
      OrdersService.postOrders(selected.value, table.value.id)
        .then((res) => {
          table.value = res.table;
          if (res.order_item.item_type === ItemType.Drink) {
            drinkOrders.value.push(res);
            sortOrders(drinkOrders.value);
          } else {
            foodOrders.value.push(res);
            sortOrders(foodOrders.value);
          }
        })
        .finally(() => {
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
