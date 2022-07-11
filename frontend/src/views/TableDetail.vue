<template>
  <BaseCard>
    <BaseToolbar title="Speisen" icon="fa-cheese" @click="addBeverage(ItemType.Food)" />
    <BaseToolbar title="Getränke" icon="fa-champagne-glasses" @click="addBeverage(ItemType.Drink)" />
    <BaseItem>{{ table }}</BaseItem>
    <BottomNavigation>
      <template #left>
        <router-link to="/tables">
          <Button icon="pi pi-arrow-left" />
        </router-link>
      </template>
      <template #middle>
        <div class="flex flex-column align-items-center">
          <div class="font-bold">Tisch {{ table.id }}</div>
          <div>{{ convertToEur(table.total) }}</div>
        </div>
      </template>
      <template #right>
        <router-link to="/bills">
          <Button icon="pi pi-money-bill" class="p-button-danger" />
        </router-link>
      </template>
    </BottomNavigation>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import BaseCard from "@/components/BaseCard.vue";
import BaseItem from "@/components/BaseItem.vue";
import { useStore } from "vuex";
import { service_Table } from "@/services/openapi";
import BottomNavigation from "@/components/BottomNavigation.vue";
import Button from "primevue/button";
import { convertToEur, ItemType } from "@/utils";
import BaseToolbar from "@/components/BaseToolbar.vue";

export default defineComponent({
  name: "TableDetail",
  components: { BaseToolbar, BottomNavigation, BaseCard, BaseItem, Button },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const isLoading = ref(false);
    const store = useStore();
    const tables = computed(() => store.getters.getTables);
    const table = tables.value.find((table: service_Table) => table.id === parseInt(props.id));

    function addBeverage(type: ItemType) {
      console.log("add", type);
    }

    return { table, isLoading, convertToEur, addBeverage, ItemType };
  },
});
</script>
