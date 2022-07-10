<template>
  <WaveSpinner v-show="isLoading" />
  <Transition>
    <BaseCard v-if="tables" style="min-height: 3rem">
      <div class="grid">
        <TableCard v-for="table in tables" v-bind:key="table.id" :table="table" />
      </div>
    </BaseCard>
  </Transition>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/BaseCard.vue";
import { useStore } from "vuex";
import WaveSpinner from "@/components/WaveSpinner.vue";
import TableCard from "@/components/TableCard.vue";

export default defineComponent({
  name: "TablesView",
  components: { TableCard, BaseCard, WaveSpinner },
  setup() {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);

    return { tables, isLoading };
  },
});
</script>

<style></style>
