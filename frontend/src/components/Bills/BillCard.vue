<template>
  <div class="col-12 lg:col-6">
    <BaseItem class="flex flex-column" bgColor="d">
      <div class="mb-1">Tisch {{ bill.table_id }}</div>
      <div class="flex justify-content-between">
        <TheBadge>{{ date }}</TheBadge>
        <div class="font-bold">{{ convertToEur(bill.total) }}</div>
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
import TheBadge from "@/components/UI/TheBadge.vue";

export default defineComponent({
  name: "BillCard",
  components: { TheBadge, BaseItem },
  props: {
    bill: { type: Object as PropType<service_Bill>, required: true },
  },
  setup(props) {
    const date = computed(() => props.bill.created_at && moment.unix(props.bill.created_at).format("HH:mm") + " Uhr");
    return { convertToEur, date };
  },
});
</script>
