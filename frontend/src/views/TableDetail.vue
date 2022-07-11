<template>
  <BaseCard>
    <BaseToolbar title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" />
    <BaseToolbar title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" />
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
import { service_Table } from "@/services/openapi";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import Button from "primevue/button";
import { convertToEur, ItemType } from "@/utils";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import Listbox from "primevue/listbox";
import Dialog from "primevue/dialog";

export default defineComponent({
  name: "TableDetail",
  components: { BaseToolbar, BottomNavigation, BaseCard, Button, Dialog, Listbox },
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

    async function addBeverage(type: ItemType) {
      modal.value = true;
      await store.dispatch("getOrderItems", type);
      options.value = orderItems.value.get(type);
    }

    function addOrder() {
      modal.value = false;
      selected.value = undefined;
    }

    return { modal, selected, options, table, isLoading, convertToEur, addBeverage, ItemType, addOrder };
  },
});
</script>

<style>
.p-dialog-content {
  padding: 0.5rem !important;
}
</style>
