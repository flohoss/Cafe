<template>
  <BaseCard>
    <ConfirmDialog></ConfirmDialog>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else class="p-card shadow-1 md:p-3">
        <DataTable :value="bills" dataKey="id" :filters="filters" responsiveLayout="scroll" stripedRows class="p-datatable-sm">
          <template #header>
            <div class="grid p-fluid align-items-center">
              <div class="col-12 md:col-4">
                <Calendar id="basic" v-model="today" autocomplete="off" inputStyle="text-align:center" :manualInput="false" />
              </div>
              <div class="col-12 md:col-8">
                <span class="p-input-icon-left">
                  <i class="pi pi-search" />
                  <InputText v-model="filters['global'].value" placeholder="Suchen" @keydown.esc="filters['global'].value = null" />
                  <span v-if="filters['global'].value !== null" class="leftMiddle styling" @click="filters['global'].value = null">
                    <i class="pi pi-times"></i>
                  </span>
                </span>
              </div>
            </div>
          </template>

          <Column field="table_id">
            <template #body="slotProps">
              <span class="white-space-nowrap">
                Tisch {{ slotProps.data.table_id }} <span class="text-sm">({{ time(slotProps.data.created_at) }})</span>
              </span>
            </template>
          </Column>
          <Column field="total" style="text-align: right">
            <template #body="slotProps">{{ convertToEur(slotProps.data.total) }}</template>
          </Column>
          <Column style="width: 3.5rem">
            <template #body="slotProps">
              <div class="flex align-items-center justify-content-end">
                <div class="btn success mr-2" @click="openBill(slotProps.data.id)">
                  <i class="pi pi-eye"></i>
                </div>
                <div class="btn danger" @click="deleteBill(slotProps.data.id)">
                  <i class="pi pi-trash"></i>
                </div>
              </div>
            </template>
          </Column>

          <template #empty><div class="mb-1">Keine Rechnungen</div></template>
        </DataTable>
      </div>
    </Transition>
    <Sidebar v-model:visible="billModal" :baseZIndex="10000" position="full">
      <BillModal :bill="bill" />
    </Sidebar>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import Calendar from "primevue/calendar";
import { BillsService, service_Bill } from "@/services/openapi";
import Sidebar from "primevue/sidebar";
import BillModal from "@/components/Bills/BillModal.vue";
import { convertToEur, emptyBill, errorToast } from "@/utils";
import { FilterMatchMode } from "primevue/api";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import moment from "moment";
import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import ConfirmDialog from "primevue/confirmdialog";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";

export default defineComponent({
  name: "BillView",
  components: { BillModal, BaseCard, Calendar, Sidebar, DataTable, Column, InputText, ConfirmDialog, WaveSpinner },
  setup() {
    const confirm = useConfirm();
    const toast = useToast();
    const today = ref(new Date());
    const bills = ref<service_Bill[]>([]);
    const isLoading = ref(false);
    const billModal = ref(false);
    const bill = ref<service_Bill>({ ...emptyBill });
    const filters = ref({
      global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });

    getData();
    watch(today, () => getData());

    function getData() {
      isLoading.value = true;
      BillsService.getBills(today.value.getFullYear(), today.value.getUTCMonth() + 1, today.value.getDate())
        .then((res) => (bills.value = res))
        .finally(() => {
          isLoading.value = false;
        });
    }

    function openBill(billId: number) {
      const temp: service_Bill | undefined = bills.value.find((bill) => bill.id === billId);
      temp && (bill.value = temp);
      billModal.value = true;
    }

    function time(unixDate: number) {
      return moment.unix(unixDate).format("HH:mm") + " Uhr";
    }

    function deleteBill(billId: number) {
      confirm.require({
        message: "Rechnung löschen?",
        header: "Rechnung",
        icon: "pi pi-info-circle",
        acceptClass: "p-button-danger",
        accept: () => {
          isLoading.value = true;
          BillsService.deleteBills(billId)
            .then(() => (bills.value = bills.value.filter((bill) => bill.id !== billId)))
            .catch((err) => errorToast(toast, err.body.error))
            .finally(() => (isLoading.value = false));
        },
      });
    }

    return { convertToEur, openBill, deleteBill, today, bills, isLoading, filters, billModal, bill, time };
  },
});
</script>

<style scoped>
.danger {
  cursor: pointer;
  color: red;
}
.success {
  cursor: pointer;
  color: green;
}
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
