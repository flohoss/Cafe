<template>
  <BaseCard>
    <ConfirmDialog></ConfirmDialog>
    <DataTable :value="orderItems" dataKey="id" :filters="filters" responsiveLayout="scroll" :showAddButton="true" stripedRows class="p-datatable-sm">
      <template #header>
        <div class="grid p-fluid align-items-center">
          <div class="col-10">
            <span class="p-input-icon-left">
              <i class="pi pi-search" />
              <InputText v-model="filters['global'].value" placeholder="Suchen..." @keydown.esc="filters['global'].value = null" />
              <span v-if="filters['global'].value !== null" class="leftMiddle styling" @click="filters['global'].value = null">
                <i class="pi pi-times"></i>
              </span>
            </span>
          </div>
          <div class="col-2 text-right">
            <Button icon="pi pi-plus" class="p-button-rounded" @click="modal = true" />
          </div>
        </div>
      </template>

      <Column field="description"></Column>
      <Column field="price" style="text-align: right">
        <template #body="slotProps">{{ convertToEur(slotProps.data.price) }}</template>
      </Column>
      <Column class="flex justify-content-end flex-nowrap">
        <template #body="slotProps">
          <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-1" @click="editOrderItem(slotProps.data)" />
          <Button icon="pi pi-trash" class="p-button-rounded p-button-warning" @click="confirmDeleteProduct(slotProps.data)" />
        </template>
      </Column>

      <template #empty>Keine Einträge</template>
    </DataTable>

    <Dialog v-model:visible="modal" :modal="true" :showHeader="false" @hide="resetModal">
      <div class="p-fluid">
        <div class="field mt-5">
          <InputText id="name" v-model.trim="orderItem.description" required="true" autofocus />
        </div>
        <div class="field">
          <InputNumber id="currency-germany" v-model="orderItem.price" mode="currency" currency="EUR" locale="de-DE" />
        </div>
      </div>
      <div class="flex justify-content-end">
        <Button icon="pi pi-times" class="p-button-text p-button-rounded p-button-secondary mr-2" @click="resetModal" />
        <Button icon="pi pi-check" class="p-button-rounded p-button-success" @click="saveOrderItem" />
      </div>
    </Dialog>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import InputText from "primevue/inputtext";
import { FilterMatchMode } from "primevue/api";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Button from "primevue/button";
import { convertToEur } from "@/utils";
import Dialog from "primevue/dialog";
import InputNumber from "primevue/inputnumber";
import { useConfirm } from "primevue/useconfirm";
import ConfirmDialog from "primevue/confirmdialog";

export default defineComponent({
  name: "OrderItemList",
  components: { BaseCard, InputText, DataTable, Column, Button, Dialog, InputNumber, ConfirmDialog },
  props: {
    orderItems: { type: Array as PropType<service_OrderItem[]>, default: () => [] },
    emptyOrderItem: { type: Object as PropType<service_OrderItem>, default: () => ({}) },
  },
  emits: ["orderItemChanged", "orderItemDeleted", "orderItemCreated"],
  setup(props, { emit }) {
    const confirm = useConfirm();
    const modal = ref(false);
    const filters = ref({
      global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });
    const orderItem = ref<service_OrderItem>({ ...props.emptyOrderItem });
    function editOrderItem(item: service_OrderItem) {
      orderItem.value = { ...item };
      modal.value = true;
    }
    watch(props.emptyOrderItem, () => resetModal());

    function saveOrderItem() {
      if (orderItem.value.id) {
        OrderItemsService.putOrdersItems(orderItem.value)
          .then((res) => emit("orderItemChanged", res))
          .finally(() => resetModal());
      } else {
        OrderItemsService.postOrdersItems(orderItem.value)
          .then((res) => emit("orderItemCreated", res))
          .finally(() => resetModal());
      }
    }

    function confirmDeleteProduct(item: service_OrderItem) {
      confirm.require({
        message: item.description + " löschen?",
        header: "Achtung",
        icon: "pi pi-exclamation-triangle",
        accept: () => {
          item.id &&
            OrderItemsService.deleteOrdersItems(item.id)
              .then(() => emit("orderItemDeleted", item))
              .finally(() => resetModal());
        },
      });
    }

    function resetModal() {
      modal.value = false;
      orderItem.value = { ...props.emptyOrderItem };
    }

    return { filters, convertToEur, editOrderItem, saveOrderItem, confirmDeleteProduct, modal, orderItem, resetModal };
  },
});
</script>

<style>
.p-datatable,
.p-datatable-header {
  background-color: transparent !important;
  margin-top: 0.5rem;
  padding: 0 !important;
  border: 0 !important;
}
.p-datatable-thead {
  display: none;
}
</style>

<style scoped>
.styling {
  cursor: pointer;
  color: gray;
  background-color: var(--surface-a);
  border-radius: 50%;
  padding: 0.2rem;
}
.leftMiddle {
  position: absolute;
  top: 50%;
  right: 0;
  transform: translate(-50%, -50%);
}
</style>
