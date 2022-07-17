<template>
  <Transition>
    <div class="container">
      <div class="flex flex-column align-items-center justify-content-center">
        <img alt="logo" class="h-6rem mb-3" />
        <div class="text-center text-sm">Plätschwiesen 2, 72622 Nürtingen<br />Baden-Württemberg</div>
      </div>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <div class="flex justify-content-between my-5">
          <div>{{ date }}</div>
          <div class="mb-1">Tisch {{ bill.table_id }}</div>
          <div>{{ time }}</div>
        </div>
        <div class="text-lg">
          <div v-for="billItem in billItems" :key="billItem.id" class="flex flex-column mb-1">
            <div class="flex align-items-center justify-content-between">
              <div class="entry white-space-nowrap overflow-hidden">{{ billItem.description }}</div>
              <div>{{ convertToEur(billItem.total) }}</div>
            </div>
            <div v-if="billItem.amount !== 1" class="ml-4 font-italic text-sm">{{ billItem.amount }} x {{ convertToEur(billItem.price) }}</div>
          </div>
          <div class="flex justify-content-end font-bold mt-5">Total: {{ convertToEur(bill.total) }}</div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, PropType, ref } from "vue";
import { BillsService, service_Bill, service_BillItem } from "@/services/openapi";
import { convertToEur } from "@/utils";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import moment from "moment";

export default defineComponent({
  name: "BillModal",
  components: { WaveSpinner },
  props: { bill: { type: Object as PropType<service_Bill>, required: true } },
  setup(props) {
    const isLoading = ref(true);
    const billItems = ref<service_BillItem[]>();
    const date = computed(() => props.bill.created_at && moment.unix(props.bill.created_at).format("DD.MM.YYYY"));
    const time = computed(() => props.bill.created_at && moment.unix(props.bill.created_at).format("HH:mm") + " Uhr");
    onMounted(() => {
      props.bill.id &&
        BillsService.getBillsItems(props.bill.id)
          .then((res) => {
            billItems.value = res;
          })
          .finally(() => (isLoading.value = false));
    });

    return { convertToEur, isLoading, billItems, date, time };
  },
});
</script>

<style scoped>
.container {
  --bs-gutter-x: 0;
  --bs-gutter-y: 0;
  width: 100%;
  padding-right: 0.5rem;
  padding-left: 0.5rem;
  margin-right: auto;
  margin-left: auto;
}
@media (min-width: 576px) {
  .container {
    max-width: 540px;
  }
}
@media (min-width: 768px) {
  .container {
    max-width: 720px;
  }
}
@media (min-width: 992px) {
  .container {
    max-width: 960px;
  }
}
@media (min-width: 1200px) {
  .container {
    max-width: 1140px;
  }
}
@media (min-width: 1400px) {
  .container {
    max-width: 1320px;
  }
}
@media (prefers-color-scheme: light) {
  img {
    content: url("../../assets/logo.png");
  }
}
@media (prefers-color-scheme: dark) {
  img {
    content: url("../../assets/logo_white.png");
  }
}
.entry:first-child:after {
  content: " . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . ";
}
</style>
