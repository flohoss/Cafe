<template>
  <BaseCard>
    <BaseItem bgColor="c" class="flex justify-content-center mb-2">
      <Calendar id="basic" v-model="today" autocomplete="off" :showIcon="true" inputStyle="text-align:center" :manualInput="false" />
    </BaseItem>
    <EmptyView v-if="bills.length === 0" message="Keine Rechnungen" />
    <BillCard v-for="bill in bills" v-bind:key="bill.id" :bill="bill" />
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

export default defineComponent({
  name: "BillView",
  components: { EmptyView, BillCard, BaseItem, BaseCard, Calendar },
  setup() {
    const today = ref(new Date());
    const bills = ref<service_Bill[]>([]);

    onMounted(() => getData());
    watch(today, () => getData());

    function getData() {
      BillsService.getBills(today.value.getFullYear(), today.value.getUTCMonth() + 1, today.value.getDate()).then((res) => (bills.value = res));
    }
    return { today, bills };
  },
});
</script>

<style></style>
