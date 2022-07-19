<template>
  <SmallCard bgColor="a" :to="'/tables/' + table.id">
    <template #description>Tisch {{ table.id }}</template>
    <template #badge>{{ since }}</template>
    <template #right>
      <div class="flex align-items-end">
        <TheBadge v-if="table.order_count" class="topRight">{{ table.order_count }}</TheBadge>
        <div v-if="table.total" class="font-bold">{{ convertToEur(table.total) }}</div>
      </div>
    </template>
  </SmallCard>
</template>

<script lang="ts">
import { defineComponent, onMounted, onUnmounted, PropType, ref } from "vue";
import { service_Table } from "@/services/openapi";
import moment from "moment";
import { convertToEur, getCurrentTimeSince } from "@/utils";
import TheBadge from "@/components/UI/TheBadge.vue";
import SmallCard from "@/components/UI/SmallCard.vue";

export default defineComponent({
  name: "TableCard",
  components: { TheBadge, SmallCard },
  props: { table: { type: Object as PropType<service_Table>, required: true } },
  setup(props) {
    moment.locale("de");
    // eslint-disable-next-line
    let ticker: any;
    const since = ref(getCurrentTimeSince(props.table.updated_at));
    onMounted(() => {
      ticker = setInterval(() => {
        since.value = getCurrentTimeSince(props.table.updated_at);
      }, 1000);
    });
    onUnmounted(() => ticker && clearInterval(ticker));
    return { since, convertToEur };
  },
});
</script>

<style scoped>
.topRight {
  position: absolute;
  top: -0.2rem;
  right: 0.3rem;
}
</style>
