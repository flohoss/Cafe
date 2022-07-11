<template>
  <BaseCard v-if="tables.length !== 0">
    <TableCard v-for="table in tables" v-bind:key="table.id" :table="table" />
  </BaseCard>
  <BaseCard v-else class="text-center">
    <div class="p-card w-full p-3">Keine Tische angelegt</div>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/BaseCard.vue";
import { useStore } from "vuex";
import TableCard from "@/components/TableCard.vue";

export default defineComponent({
  name: "TablesView",
  components: { TableCard, BaseCard },
  setup() {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);

    return { tables, isLoading };
  },
});
</script>
