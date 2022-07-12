<template>
  <router-link class="col-12 md:col-6 no-underline" :to="'/tables/' + table.id">
    <BaseItem class="relative">
      <Badge v-if="table.order_count" :value="table.order_count" class="topRight text-sm" />
      <div class="flex justify-content-between align-items-end">
        <div>
          <div class="font-bold mb-2">Tisch {{ table.id }}</div>
          <div class="text-sm">{{ since }}</div>
        </div>
        <div>
          <div v-if="table.total" class="font-bold">{{ convertToEur(table.total) }}</div>
        </div>
      </div>
    </BaseItem>
  </router-link>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { service_Table } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import moment from "moment";
import { convertToEur } from "@/utils";
import Badge from "primevue/badge";

export default defineComponent({
  name: "TableCard",
  components: { Badge, BaseItem },
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
  right: 0.3rem;
}
</style>
