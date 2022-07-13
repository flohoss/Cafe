<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <BaseToolbar :isDisabled="isDisabled" title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" btnIcon="plus" />
        <div class="grid">
          <TableOrderCard
            v-for="entry in food"
            v-bind:key="entry.id"
            :order="entry"
            :isDisabled="isDisabled"
            @incrementOrder="(order) => incrementOrder(order)"
            @decrementOrder="(order) => decrementOrder(order)"
          />
        </div>
        <BaseToolbar :isDisabled="isDisabled" title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" btnIcon="plus" />
        <div class="grid">
          <TableOrderCard
            v-for="entry in drinks"
            v-bind:key="entry.id"
            :order="entry"
            :isDisabled="isDisabled"
            @incrementOrder="(order) => incrementOrder(order)"
            @decrementOrder="(order) => decrementOrder(order)"
          />
        </div>
        <div class="h-4rem"></div>
      </div>
    </Transition>

    <Sidebar v-model:visible="modal" :modal="true" :baseZIndex="10000" position="full">
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
        <Button :loading="isDisabled" label="Speichern" icon="pi pi-check" class="p-button p-button-success mr-3" @click="postOrder" />
      </div>
    </Sidebar>

    <BottomNavigation>
      <template #left>
        <router-link to="/tables" class="no-underline">
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
        <router-link to="/bills" class="no-underline">
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
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import Listbox from "primevue/listbox";
import TableOrderCard from "@/components/Tables/TableOrderCard.vue";
import Sidebar from "primevue/sidebar";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";

export default defineComponent({
  name: "TableOverview",
  components: { WaveSpinner, TableOrderCard, BaseToolbar, BottomNavigation, BaseCard, Button, Sidebar, Listbox },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const isDisabled = ref(false);
    const modal = ref(false);
    const store = useStore();
    const selected = ref();
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const orders = ref<service_Order[]>([]);
    const drinks = computed(() => orders.value.filter((order) => order.order_item.item_type === ItemType.Drink));
    const food = computed(() => orders.value.filter((order) => order.order_item.item_type === ItemType.Food));

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
      modal.value = false;
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
      modal.value = true;
      options.value = orderItems.value.get(type);
    }

    function postOrder() {
      isDisabled.value = true;
      OrdersService.postOrders(selected.value, table.value).finally(() => getData());
    }

    function incrementOrder(order: service_Order) {
      isDisabled.value = true;
      OrdersService.postOrders(order.order_item_id, order.table_id).finally(() => getData());
    }

    function decrementOrder(order: service_Order) {
      isDisabled.value = true;
      OrdersService.deleteOrders(order.order_item_id, order.table_id).finally(() => getData());
    }

    return {
      modal,
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
      drinks,
      food,
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
