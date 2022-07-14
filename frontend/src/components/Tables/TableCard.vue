<template>
  <router-link class="col-12 lg:col-6 no-underline" :to="'/tables/' + table.id">
    <BaseItem class="relative">
      <TheBadge v-if="table.order_count" class="topRight">{{ table.order_count }}</TheBadge>
      <div class="flex justify-content-between align-items-end">
        <div>
          <div class="font-bold mb-1">Tisch {{ table.id }}</div>
          <TheBadge color="success">{{ since }}</TheBadge>
        </div>
        <div>
          <div v-if="table.total" class="font-bold">{{ convertToEur(table.total) }}</div>
        </div>
      </div>
    </BaseItem>
  </router-link>
</template>

<script lang="ts">
import { defineComponent, onMounted, onUnmounted, PropType, ref } from "vue";
import { service_Table } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import moment from "moment";
import { convertToEur, getCurrentTimeSince } from "@/utils";
import TheBadge from "@/components/UI/TheBadge.vue";

export default defineComponent({
  name: "TableCard",
  components: { TheBadge, BaseItem },
  props: { table: { type: Object as PropType<service_Table>, required: true } },
  setup(props) {
    moment.locale("de");
    let ticker: null | number = null;
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
