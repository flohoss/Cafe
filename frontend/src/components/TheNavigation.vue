<template>
  <Menubar v-if="isLoggedIn" :model="items" class="p-1 mb-2 shadow-1 border-0 bg-color">
    <template #start>
      <router-link to="/"><img alt="logo" class="h-2-5rem mx-2" /></router-link>
    </template>
    <template #end>
      <div v-if="tablePath">
        <Button :disabled="isLoading" icon="pi pi-minus" class="p-button-danger mr-2" @click="removeTable" />
        <Button :disabled="isLoading" icon="pi pi-plus" class="p-button-success" @click="addTable" />
      </div>
      <ul v-else>
        <li class="p-menuitem">
          <router-link to="/tables" class="p-menuitem-link py-1"><i class="pi pi-table mr-2"></i>Tische</router-link>
        </li>
      </ul>
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
          store.commit("popTables");
        })
        .finally(() => {
          isLoading.value = false;
        });
    }
    function addTable() {
      isLoading.value = true;
      TablesService.postTables()
        .then((res) => {
          store.commit("pushTable", res);
        })
        .finally(() => {
          isLoading.value = false;
        });
    }

    const items = ref([
      {
        label: "Bestellungen",
        icon: "pi pi-fw pi-history",
        to: "/orders",
      },
      {
        label: "Artikel",
        icon: "pi pi-fw pi-shopping-cart",
        items: [
          { label: "Speisen", to: "/foods" },
          { label: "Getränke", to: "/drinks" },
        ],
      },
      {
        label: "Rechnungen",
        icon: "pi pi-fw pi-euro",
        to: "/bills",
      },
      {
        separator: true,
      },
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
  background-color: var(--surface-a);
  color: var(--text-color);
}
.h-2-5rem {
  height: 2.5rem !important;
}
@media (prefers-color-scheme: light) {
  img {
    content: url("../assets/logo.png");
  }
}
@media (prefers-color-scheme: dark) {
  img {
    content: url("../assets/logo_white.png");
  }
}
</style>
