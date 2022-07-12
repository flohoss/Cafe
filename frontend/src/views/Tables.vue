<template>
  <BaseCard>
    <div v-if="tables.length === 0" class="p-card w-full p-4 text-center">Keine Tische</div>
    <TableCard v-else v-for="table in tables" v-bind:key="table.id" :table="table" />
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { useStore } from "vuex";
import TableCard from "@/components/Table/TableEntry.vue";

export default defineComponent({
  name: "TablesView",
  components: { TableCard, BaseCard },
  setup() {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);

    function getData() {
      store.dispatch("getTables");
    }
    onMounted(() => getData());

    return { tables, isLoading };
  },
});
</script>
