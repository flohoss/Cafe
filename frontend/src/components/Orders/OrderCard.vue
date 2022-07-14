<template>
  <div class="col-12 lg:col-6">
    <BaseItem class="relative" bgColor="d" :border="border">
      <div class="flex justify-content-between overflow-hidden">
        <div class="flex flex-column align-items-start">
          <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis font-bold">{{ order.order_item.description }}</div>
          <div class="flex align-items-center mt-1">
            <TheBadge size="sm" color="info">Tisch {{ order.table_id }}</TheBadge>
            <TheBadge size="sm" color="warning" class="ml-2">{{ since }}</TheBadge>
          </div>
        </div>
        <div class="flex align-items-center">
          <Button v-if="!newOrder" :disabled="isDisabled" icon="pi pi-check" class="p-button-rounded p-button-success" @click="$emit('orderDone', order)" />
          <TheBadge v-else color="danger">NEU</TheBadge>
        </div>
      </div>
    </BaseItem>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, onUnmounted, PropType, ref } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur, getCurrentTimeSince, ItemType } from "@/utils";
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
    moment.locale("de");
    let ticker: null | number = null;
    const since = ref(getCurrentTimeSince(props.order.updated_at));
    onMounted(() => {
      ticker = setInterval(() => {
        since.value = getCurrentTimeSince(props.order.updated_at);
      }, 1000);
    });
    onUnmounted(() => ticker && clearInterval(ticker));
    const newOrder = computed(() => since.value === "vor ein paar Sekunden");
    const border = computed(() => (props.order.order_item.item_type === ItemType.Food ? "2px dashed red" : "2px dashed var(--primary-color)"));
    return { convertToEur, ItemType, since, newOrder, border };
  },
});
</script>
