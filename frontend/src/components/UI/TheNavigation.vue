<template>
  <Menubar v-if="isLoggedIn" :model="items" class="p-1 mb-2 shadow-1 border-0 bg-color">
    <template #start>
      <router-link to="/"><img alt="logo" class="h-2-5rem mx-2" /></router-link>
    </template>
    <template #end>
      <div v-if="tablePath">
        <Button :disabled="isLoading" icon="pi pi-minus" class="p-button-danger p-button-rounded mr-2" @click="removeTable" />
        <Button :disabled="isLoading" icon="pi pi-plus" class="p-button-success p-button-rounded" @click="addTable" />
      </div>
      <router-link v-else to="/tables" class="mr-1">
        <Button label="Tische" class="p-button-secondary" icon="pi pi-table" />
      </router-link>
    </template>
  </Menubar>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import Menubar from "primevue/menubar";
import { useStore } from "vuex";
import Button from "primevue/button";
import { useRoute, useRouter } from "vue-router";
import { TablesService } from "@/services/openapi";
import tables from "@/views/Tables.vue";
import { ItemType } from "@/utils";

export default defineComponent({
  name: "TheNavigation",
  components: { Menubar, Button },
  emits: ["logout"],
  setup(_, { emit }) {
    const store = useStore();
    const route = useRoute();
    const isLoading = ref(false);
    const isLoggedIn = computed(() => store.getters.getLoginApiStatus);
    const tablePath = computed(() => route.path === "/tables");

    function removeTable() {
      isLoading.value = true;
      TablesService.deleteTables()
        .then(() => {
          store.dispatch("removeLastTable");
        })
        .finally(() => {
          isLoading.value = false;
        });
    }
    function addTable() {
      isLoading.value = true;
      TablesService.postTables()
        .then((res) => {
          store.dispatch("addTable", res);
        })
        .finally(() => {
          isLoading.value = false;
        });
    }

    const items = ref([
      { label: "Bestellungen", icon: "pi pi-fw pi-history", to: "/orders" },
      { label: "Speisen", icon: "pi pi-fw pi-shopping-cart", to: "/items/" + ItemType.Food },
      { label: "Getränke", icon: "pi pi-fw pi-shopping-cart", to: "/items/" + ItemType.Drink },
      { label: "Rechnungen", icon: "pi pi-fw pi-euro", to: "/bills" },
      { separator: true },
      {
        label: "Logout",
        icon: "pi pi-fw pi-power-off",
        command: () => {
          emit("logout");
        },
      },
    ]);

    return { items, isLoggedIn, tablePath, removeTable, addTable, isLoading };
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-b);
  color: var(--text-color);
}
.h-2-5rem {
  height: 2.5rem !important;
}
@media (prefers-color-scheme: light) {
  img {
    content: url("../../assets/logo.png");
  }
}
@media (prefers-color-scheme: dark) {
  img {
    content: url("../../assets/logo_white.png");
  }
}
</style>
