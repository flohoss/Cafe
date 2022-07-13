<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <EmptyView v-else-if="tables && tables.length === 0" message="Keine Tische" />
      <div v-else class="grid">
        <TableCard v-for="table in tables" v-bind:key="table.id" :table="table" />
      </div>
    </Transition>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { useStore } from "vuex";
import TableCard from "@/components/Tables/TableCard.vue";
import EmptyView from "@/views/Empty.vue";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import { TablesService } from "@/services/openapi";

export default defineComponent({
  name: "TablesView",
  components: { WaveSpinner, EmptyView, TableCard, BaseCard },
  setup() {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);

    function getData() {
      isLoading.value = true;
      store.dispatch("fetchTables").then(() => (isLoading.value = false));
    }
    onMounted(() => getData());

    return { tables, isLoading };
  },
});
</script>
