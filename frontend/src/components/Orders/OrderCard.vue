<template>
  <div class="col-12 lg:col-6">
    <BaseItem>
      <div class="flex justify-content-between overflow-hidden">
        <div class="flex flex-column align-items-start">
          <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis font-bold">{{ order.order_item.description }}</div>
          <div class="flex align-items-center mt-1">
            <TheBadge size="sm" color="danger">Tisch {{ order.table_id }}</TheBadge>
            <TheBadge size="sm" color="info" class="ml-2">{{ time }} Uhr</TheBadge>
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
import TheBadge from "@/components/UI/TheBadge.vue";

export default defineComponent({
  name: "OrderCard",
  components: { TheBadge, BaseItem, Button },
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
