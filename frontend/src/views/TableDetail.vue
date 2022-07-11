<template>
  <BaseCard>
    <BaseToolbar :isDisabled="isLoading" title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" />
    <OrderEntry
      v-for="entry in foodOrders"
      v-bind:key="entry.id"
      :order="entry"
      :isDisabled="isLoading"
      @incrementOrder="(order) => incrementOrder(order)"
      @decrementOrder="(order) => decrementOrder(order)"
    />
    <BaseToolbar :isDisabled="isLoading" title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" />
    <OrderEntry
      v-for="entry in drinkOrders"
      v-bind:key="entry.id"
      :order="entry"
      :isDisabled="isLoading"
      @incrementOrder="(order) => incrementOrder(order)"
      @decrementOrder="(order) => decrementOrder(order)"
    />
    <Dialog position="top" v-model:visible="modal" :modal="true" :showHeader="false" style="min-width: 50vw">
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
        <Button icon="pi pi-check" class="p-button-rounded p-button-success" @click="postOrder" />
      </div>
    </Dialog>

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
import Dialog from "primevue/dialog";
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
    const table = computed(() => parseInt(props.id));
    const total = ref(0);
    const orderItems = computed(() => store.getters.getOrderItems);
    const options = ref();
    const drinkOrders = ref<service_Order[]>([]);
    const foodOrders = ref<service_Order[]>([]);
    const currentItemType = ref(0);

    getData();
    async function getData(itemType?: ItemType) {
      if (itemType) {
        await store.dispatch("getOrderItems", itemType);
        if (itemType === ItemType.Drink) {
          drinkOrders.value = await OrdersService.getOrders(table.value, ItemType.Drink);
        } else {
          foodOrders.value = await OrdersService.getOrders(table.value, ItemType.Food);
        }
      } else {
        await store.dispatch("getAllOrderItems");
        drinkOrders.value = await OrdersService.getOrders(table.value, ItemType.Drink);
        foodOrders.value = await OrdersService.getOrders(table.value, ItemType.Food);
      }
      updateTotal();
    }

    function updateTotal() {
      if (drinkOrders.value.length > 0) {
        total.value = drinkOrders.value[0].table.total;
      } else if (foodOrders.value.length > 0) {
        total.value = foodOrders.value[0].table.total;
      } else {
        total.value = 0;
      }
    }

    async function addBeverage(type: ItemType) {
      modal.value = true;
      currentItemType.value = type;
      options.value = orderItems.value.get(type);
    }

    function postOrder() {
      OrdersService.postOrders(selected.value, table.value)
        .then((res) => {
          getData(res.order_item.item_type);
        })
        .finally(() => {
          modal.value = false;
          selected.value = undefined;
        });
    }

    function incrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.postOrders(order.order_item_id, order.table_id).then(async () => {
        await getData(order.order_item.item_type);
        isLoading.value = false;
      });
    }

    function decrementOrder(order: service_Order) {
      isLoading.value = true;
      OrdersService.deleteOrders(order.order_item_id, order.table_id).finally(async () => {
        await getData(order.order_item.item_type);
        isLoading.value = false;
      });
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
      foodOrders,
      drinkOrders,
      incrementOrder,
      decrementOrder,
    };
  },
});
</script>

<style>
.p-dialog-content {
  padding: 0.5rem !important;
}
</style>
