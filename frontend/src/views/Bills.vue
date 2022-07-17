<template>
  <BaseCard>
    <BaseItem bgColor="c" class="flex justify-content-center mb-2">
      <Calendar id="basic" v-model="today" autocomplete="off" inputStyle="text-align:center" :manualInput="false" />
    </BaseItem>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <EmptyView v-if="bills.length === 0" message="Keine Rechnungen" />
        <div v-else class="grid">
          <BillCard v-for="bill in bills" v-bind:key="bill.id" :bill="bill" style="cursor: pointer" @click="openBill(bill.id)" />
        </div>
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
import BaseItem from "@/components/UI/BaseItem.vue";
import Calendar from "primevue/calendar";
import BillCard from "@/components/Bills/BillCard.vue";
import { BillsService, service_Bill } from "@/services/openapi";
import EmptyView from "@/views/Empty.vue";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import Sidebar from "primevue/sidebar";
import BillModal from "@/components/Bills/BillModal.vue";
import { emptyBill } from "@/utils";

export default defineComponent({
  name: "BillView",
  components: { BillModal, WaveSpinner, EmptyView, BillCard, BaseItem, BaseCard, Calendar, Sidebar },
  props: { id: { type: Number, required: true } },
  setup(props) {
    const today = ref(new Date());
    const bills = ref<service_Bill[]>([]);
    const isLoading = ref(false);
    const billModal = ref(false);
    const bill = ref<service_Bill>({ ...emptyBill });

    getData();
    watch(today, () => getData());

    function getData() {
      isLoading.value = true;
      BillsService.getBills(today.value.getFullYear(), today.value.getUTCMonth() + 1, today.value.getDate())
        .then((res) => (bills.value = res))
        .finally(() => {
          if (props.id) {
            openBill(props.id);
          }
          isLoading.value = false;
        });
    }

    function openBill(billId: number) {
      bill.value = bills.value[billId - 1];
      billModal.value = true;
    }
    return { today, bills, isLoading, openBill, billModal, bill };
  },
});
</script>

<style></style>
