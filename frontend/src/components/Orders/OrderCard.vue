<template>
  <SmallCard bgColor="d">
    <template #description>{{ order.order_item.description }}</template>
    <template #badgeOne>Tisch {{ order.table_id }}</template>
    <template #badgeTwo>{{ since }}</template>
    <template #right>
      <div class="flex align-items-center">
        <Button v-if="!newOrder" :disabled="isDisabled" icon="pi pi-check" class="p-button-rounded p-button-success" @click="$emit('orderDone', order)" />
        <TheBadge v-else color="danger">NEU</TheBadge>
      </div>
    </template>
  </SmallCard>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, onUnmounted, PropType, ref } from "vue";
import { service_Order } from "@/services/openapi";
import { convertToEur, getCurrentTimeSince, ItemType } from "@/utils";
import Button from "primevue/button";
import moment from "moment";
import TheBadge from "@/components/UI/TheBadge.vue";
import SmallCard from "@/components/UI/SmallCard.vue";

export default defineComponent({
  name: "OrderCard",
  components: { SmallCard, TheBadge, Button },
  props: {
    order: { type: Object as PropType<service_Order>, required: true },
    isDisabled: { type: Boolean, default: false },
  },
  emits: ["orderDone"],
  setup(props) {
    moment.locale("de");
    // eslint-disable-next-line
    let ticker: any;
    const since = ref(getCurrentTimeSince(props.order.updated_at));
    onMounted(() => {
      ticker = setInterval(() => {
        since.value = getCurrentTimeSince(props.order.updated_at);
      }, 1000);
    });
    onUnmounted(() => ticker && clearInterval(ticker));
    const newOrder = computed(() => since.value === "vor ein paar Sekunden");
    return { convertToEur, ItemType, since, newOrder };
  },
});
</script>
