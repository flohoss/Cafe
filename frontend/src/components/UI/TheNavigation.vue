<template>
  <Menubar v-if="isLoggedIn" :model="items" class="p-1 mb-3 shadow-1 border-0 bg-color">
    <template #start>
      <router-link class="no-underline" to="/tables"><img alt="logo" class="h-3rem mx-2" /></router-link>
    </template>
    <template #end>
      <div v-if="tablePath">
        <Button v-if="tablesCount !== 0" :disabled="isLoading" icon="pi pi-minus" class="p-button-danger p-button-rounded mr-2" @click="removeTable" />
        <Button :disabled="isLoading" icon="pi pi-plus" class="p-button-success p-button-rounded" @click="addTable" />
      </div>
      <router-link v-else :to="{ name: 'Tables' }" class="mr-1 no-underline">
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
import { useRoute } from "vue-router";
import { TablesService } from "@/services/openapi";
import { errorToast, ItemType } from "@/utils";
import { useToast } from "primevue/usetoast";

export default defineComponent({
  name: "TheNavigation",
  components: { Menubar, Button },
  emits: ["logout"],
  setup(_, { emit }) {
    const toast = useToast();
    const store = useStore();
    const route = useRoute();
    const isLoading = ref(false);
    const tablesCount = computed(() => store.getters.getTablesCount);
    const isLoggedIn = computed(() => store.getters.getLoginApiStatus);
    const tablePath = computed(() => route.path === "/tables");

    function removeTable() {
      isLoading.value = true;
      TablesService.deleteTables()
        .then(() => {
          store.dispatch("removeLastTable");
        })
        .catch((err) => {
          errorToast(toast, err.body.error);
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
        .catch((err) => {
          errorToast(toast, err.body.error);
        })
        .finally(() => {
          isLoading.value = false;
        });
    }

    const items = ref([
      { label: "Bestellungen", icon: "pi pi-fw pi-history", to: "/orders" },
      {
        label: "Artikel",
        icon: "pi pi-fw pi-shopping-cart",
        items: [
          { label: "Speisen", icon: "pi pi-fw pi-shopping-cart", to: "/items/" + ItemType.Food },
          { label: "Getränke", icon: "pi pi-fw pi-shopping-cart", to: "/items/" + ItemType.Drink },
        ],
      },
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

    return { items, isLoggedIn, tablePath, removeTable, addTable, isLoading, tablesCount };
  },
});
</script>

<style scoped>
.bg-color {
  background-color: var(--surface-a);
  color: var(--text-color);
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
