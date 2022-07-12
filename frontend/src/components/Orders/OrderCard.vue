<template>
  <div class="col-12 lg:col-6">
    <BaseItem>
      <div class="flex justify-content-between overflow-hidden">
        <div class="flex flex-column align-items-start">
          <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis font-bold">{{ order.order_item.description }}</div>
          <div class="flex align-items-center mt-1">
            <Badge severity="danger" class="text-sm">Tisch {{ order.table_id }}</Badge>
            <Badge severity="info" class="text-sm ml-2">{{ time }} Uhr</Badge>
          </div>
        </div>
        <div class="flex align-items-center">
          <Button :disabled="isDisabled" icon="pi pi-check" class="p-button-rounded p-button-success" @click="$emit('orderDone', order)" />
        </div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, ItemType } from "@/utils";
import Button from "primevue/button";
import moment from "moment";
import Badge from "primevue/badge";

export default defineComponent({
  name: "OrderCard",
  components: { BaseItem, Button, Badge },
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
    isDisabled: { type: Boolean, default: false },
  },
  emits: ["orderDone"],
  setup(props) {
    const time = computed(() => props.order.updated_at && moment.unix(props.order.updated_at).format("DD.MM HH:mm"));
    return { convertToEur, ItemType, time };
  },
});
</script>
