<template>
  <BaseCard>
    <DataTable :value="orderItems" dataKey="id" :filters="filters" responsiveLayout="scroll" :showAddButton="true" stripedRows class="p-datatable-sm">
      <template #header>
        <div class="grid p-fluid">
          <div class="col-12">
            <span class="p-input-icon-left">
              <i class="pi pi-search" />
              <InputText v-model="filters['global'].value" placeholder="Suchen..." @keydown.esc="filters['global'].value = null" />
              <span v-if="filters['global'].value !== null" class="leftMiddle styling" @click="filters['global'].value = null">
                <i class="pi pi-times"></i>
              </span>
            </span>
          </div>
        </div>
      </template>

      <Column field="description"></Column>
      <Column field="price" style="text-align: right">
        <template #body="slotProps">{{ convertToEur(slotProps.data.price) }}</template>
      </Column>
      <Column style="text-align: right">
        <template #body="slotProps">
          <Button icon="pi pi-pencil" class="p-button-rounded p-button-success mr-2" @click="editProduct(slotProps.data)" />
          <Button icon="pi pi-trash" class="p-button-rounded p-button-warning" @click="confirmDeleteProduct(slotProps.data)" />
        </template>
      </Column>

      <template #empty>Keine Einträge</template>
    </DataTable>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { service_OrderItem } from "@/services/openapi";
import InputText from "primevue/inputtext";
import { FilterMatchMode } from "primevue/api";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Button from "primevue/button";
import { convertToEur } from "@/utils";

export default defineComponent({
  name: "OrderItemList",
  components: { BaseCard, InputText, DataTable, Column, Button },
  props: { orderItems: { type: Array as PropType<service_OrderItem[]>, default: () => [] } },
  setup() {
    const filters = ref({
      global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });
    return { filters, convertToEur };
  },
});
</script>

<style>
.p-datatable,
.p-datatable-header {
  background-color: transparent !important;
  margin-top: 1rem;
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
