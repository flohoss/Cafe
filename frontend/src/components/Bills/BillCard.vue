<template>
  <div class="col-12 lg:col-6">
    <BaseItem class="relative" bgColor="d">
      <div class="flex justify-content-between">
        <div>{{ date }}</div>
        <div>{{ convertToEur(bill.total) }}</div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Bill } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur } from "@/utils";
import moment from "moment";

export default defineComponent({
  name: "BillCard",
  components: { BaseItem },
  props: {
    bill: { type: Object as PropType<service_Bill>, required: true },
  },
  setup(props) {
    const date = computed(() => props.bill.created_at && moment.unix(props.bill.created_at).format("DD.MM.YY HH:mm") + " Uhr");
    return { convertToEur, date };
  },
});
</script>
