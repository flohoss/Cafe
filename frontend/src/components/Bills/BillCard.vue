<template>
  <div class="col-12 lg:col-6">
    <ConfirmDialog></ConfirmDialog>
    <BaseItem class="flex flex-column" bgColor="d">
      <div class="flex justify-content-between">
        <div class="mb-1">Tisch {{ bill.table_id }}</div>
      </div>
      <div class="flex justify-content-between">
        <TheBadge>{{ date }}</TheBadge>
        <div class="font-bold">{{ convertToEur(bill.total) }}</div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import { BillsService, service_Bill } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, errorToast } from "@/utils";
import moment from "moment";
import TheBadge from "@/components/UI/TheBadge.vue";
import ConfirmDialog from "primevue/confirmdialog";
import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";

export default defineComponent({
  name: "BillCard",
  components: { TheBadge, BaseItem, ConfirmDialog },
  props: {
    bill: { type: Object as PropType<service_Bill>, required: true },
  },
  emits: ["deleteBill"],
  setup(props, { emit }) {
    const confirm = useConfirm();
    const toast = useToast();
    const isLoading = ref(false);
    const date = computed(() => props.bill.created_at && moment.unix(props.bill.created_at).format("HH:mm") + " Uhr");
    function deleteBill() {
      confirm.require({
        message: "Angezeigte Bestellungen abrechnen?",
        header: "Abrechnung",
        icon: "pi pi-info-circle",
        acceptClass: "p-button-danger",
        accept: () => {
          isLoading.value = true;
          props.bill.id &&
            BillsService.deleteBills(props.bill.id)
              .then(() => emit("deleteBill"))
              .catch((err) => errorToast(toast, err.body.error))
              .finally(() => (isLoading.value = false));
        },
      });
    }
    return { convertToEur, date, deleteBill, isLoading };
  },
});
</script>
