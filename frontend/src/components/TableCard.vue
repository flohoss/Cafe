<template>
  <div class="col-12 md:col-6">
    <router-link class="no-underline" :to="'/tables/' + table.id">
      <div class="p-card shadow-1 p-2 px-3 relative">
        <Badge v-if="table.order_count" class="topRight">{{ table.order_count }}</Badge>
        <div class="flex justify-content-between align-items-end">
          <div>
            <div class="font-bold mb-2">Tisch {{ table.id }}</div>
            <div>{{ since }}</div>
          </div>
          <div>
            <div v-if="table.total" class="font-bold">{{ convertToEur(table.total) }}</div>
          </div>
        </div>
      </div>
    </router-link>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Table } from "@/services/openapi";
import moment from "moment";
import { convertToEur } from "@/utils";
import Badge from "primevue/badge";

export default defineComponent({
  name: "TableCard",
  components: { Badge },
  props: { table: { type: Object as PropType<service_Table>, required: true } },
  setup(props) {
    moment.locale("de");
    const since = computed(() => props.table.updated_at && moment.unix(props.table.updated_at).fromNow());
    return { since, convertToEur };
  },
});
</script>

<style scoped>
.topRight {
  position: absolute;
  top: -0.2rem;
  right: -0.3rem;
  color: var(--primary-color-text);
  background-color: var(--primary-color);
  border-radius: 50%;
}
</style>
