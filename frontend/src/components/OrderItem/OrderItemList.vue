<template>
  <BaseCard>
    <ConfirmDialog></ConfirmDialog>
    <div class="p-card shadow-1 md:p-3">
      <DataTable :value="orderItems" dataKey="id" :filters="filters" responsiveLayout="scroll" stripedRows class="p-datatable-sm">
        <template #header>
          <div class="grid p-fluid align-items-center">
            <div class="col-9">
              <span class="p-input-icon-left">
                <i class="pi pi-search" />
                <InputText v-model="filters['global'].value" placeholder="Suchen" @keydown.esc="filters['global'].value = null" />
                <span v-if="filters['global'].value !== null" class="leftMiddle styling" @click="filters['global'].value = null">
                  <i class="pi pi-times"></i>
                </span>
              </span>
            </div>
            <div class="col-3 text-right">
              <Button :disabled="isDisabled" icon="pi pi-plus" class="p-button-rounded" @click="modal = true" />
            </div>
          </div>
        </template>

        <Column field="description">
          <template #body="slotProps">
            <span class="white-space-nowrap">{{ slotProps.data.description }}</span>
          </template>
        </Column>
        <Column field="price" style="text-align: right">
          <template #body="slotProps">{{ convertToEur(slotProps.data.price) }}</template>
        </Column>
        <Column style="width: 3.5rem">
          <template #body="slotProps">
            <div class="flex align-items-center justify-content-end">
              <div
                class="mr-2"
                :style="{ color: isDisabled ? 'grey' : 'green', cursor: isDisabled ? 'default' : 'pointer' }"
                @click="editOrderItem(slotProps.data)"
              >
                <i class="pi pi-pencil"></i>
              </div>
              <div :style="{ color: isDisabled ? 'grey' : 'red', cursor: isDisabled ? 'default' : 'pointer' }" @click="confirmDeleteProduct(slotProps.data)">
                <i class="pi pi-trash"></i>
              </div>
            </div>
          </template>
        </Column>

        <template #empty><div class="mb-1">Keine Einträge</div></template>
      </DataTable>
    </div>

    <Dialog position="top" v-model:visible="modal" :modal="true" :showHeader="false" @hide="resetModal" style="min-width: 50vw">
      <div class="p-fluid">
        <div class="field mt-5">
          <InputText :disabled="isDisabled" id="name" v-model.trim="orderItem.description" required="true" autofocus @keydown.enter="saveOrderItem" />
        </div>
        <div class="field">
          <InputNumber
            :disabled="isDisabled"
            id="currency-germany"
            v-model="orderItem.price"
            mode="currency"
            currency="EUR"
            locale="de-DE"
            @keydown.enter="saveOrderItem"
          />
        </div>
      </div>
      <div class="flex justify-content-end">
        <Button :disabled="isDisabled" icon="pi pi-times" class="p-button-text p-button-rounded p-button-secondary mr-2" @click="resetModal" />
        <Button :loading="isDisabled" icon="pi pi-check" class="p-button-rounded p-button-success" @click="saveOrderItem" />
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
import { convertToEur, errorToast } from "@/utils";
import Dialog from "primevue/dialog";
import InputNumber from "primevue/inputnumber";
import { useConfirm } from "primevue/useconfirm";
import ConfirmDialog from "primevue/confirmdialog";
import { useToast } from "primevue/usetoast";

export default defineComponent({
  name: "OrderItemList",
  components: { BaseCard, InputText, DataTable, Column, Button, Dialog, InputNumber, ConfirmDialog },
  props: {
    orderItems: { type: Array as PropType<service_OrderItem[]>, default: () => [] },
    emptyOrderItem: { type: Object as PropType<service_OrderItem>, default: () => ({}) },
    title: { type: String, default: "" },
  },
  emits: ["orderItemChanged", "orderItemDeleted", "orderItemCreated"],
  setup(props, { emit }) {
    const isDisabled = ref(false);
    const toast = useToast();
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
      if (isDisabled.value) return;
      isDisabled.value = true;
      if (orderItem.value.id) {
        OrderItemsService.putOrdersItems(orderItem.value)
          .then((res) => emit("orderItemChanged", res))
          .catch((err) => errorToast(toast, err.body.error))
          .finally(() => resetModal());
      } else {
        OrderItemsService.postOrdersItems(orderItem.value)
          .then((res) => emit("orderItemCreated", res))
          .finally(() => resetModal());
      }
    }

    function confirmDeleteProduct(item: service_OrderItem) {
      if (isDisabled.value) return;
      confirm.require({
        message: item.description + " löschen?",
        header: "Achtung",
        position: "top",
        accept: () => {
          isDisabled.value = true;
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
      isDisabled.value = false;
    }

    return { filters, convertToEur, editOrderItem, saveOrderItem, confirmDeleteProduct, modal, orderItem, resetModal, isDisabled };
  },
});
</script>

<style>
.p-datatable .p-datatable-header,
.p-datatable .p-datatable-footer {
  background: transparent !important;
}
.p-datatable .p-datatable-header,
.p-datatable .p-datatable-footer,
.p-datatable .p-datatable-tbody > tr,
.p-datatable .p-datatable-tbody > tr > td {
  border-width: 0 !important;
}
.p-datatable-thead {
  display: none;
}
</style>

<style scoped>
.styling {
  cursor: pointer;
  color: gray;
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
