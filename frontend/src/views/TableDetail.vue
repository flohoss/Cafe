<template>
  <BaseCard class="p-card p-2">
    {{ table }}
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/BaseCard.vue";
import { useStore } from "vuex";
import WaveSpinner from "@/components/WaveSpinner.vue";
import { service_Table } from "@/services/openapi";

export default defineComponent({
  name: "TableDetail",
  components: { BaseCard },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);
    const table = tables.value.find((table: service_Table) => table.id === parseInt(props.id));

    return { table, isLoading };
  },
});
</script>

<style></style>
