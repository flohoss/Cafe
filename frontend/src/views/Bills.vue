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
          <BillCard v-for="bill in bills" v-bind:key="bill.id" :bill="bill" />
        </div>
      </div>
    </Transition>
  </BaseCard>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BaseItem from "@/components/UI/BaseItem.vue";
import Calendar from "primevue/calendar";
import BillCard from "@/components/Bills/BillCard.vue";
import { BillsService, service_Bill } from "@/services/openapi";
import EmptyView from "@/views/Empty.vue";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";

export default defineComponent({
  name: "BillView",
  components: { WaveSpinner, EmptyView, BillCard, BaseItem, BaseCard, Calendar },
  setup() {
    const today = ref(new Date());
    const bills = ref<service_Bill[]>([]);
    const isLoading = ref(false);

    onMounted(() => getData());
    watch(today, () => getData());

    function getData() {
      isLoading.value = true;
      BillsService.getBills(today.value.getFullYear(), today.value.getUTCMonth() + 1, today.value.getDate())
        .then((res) => (bills.value = res))
        .finally(() => (isLoading.value = false));
    }
    return { today, bills, isLoading };
  },
});
</script>

<style></style>
