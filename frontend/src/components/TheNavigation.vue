<template>
  <Menubar v-if="isLoggedIn" :model="items" class="p-1 px-3 mb-4">
    <template #start>
      <img alt="logo" src="../assets/logo.png" class="h-3rem mr-2" />
    </template>
    <template #end>
      <Button label="Tables" icon="pi pi-table" @click="routeToTable" />
    </template>
  </Menubar>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import Menubar from "primevue/menubar";
import Button from "primevue/button";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "TheNavigation",
  components: { Menubar, Button },
  emits: ["logout"],
  setup(_, { emit }) {
    const store = useStore();
    const router = useRouter();
    const isLoggedIn = computed(() => store.getters.getLoginApiStatus);

    function routeToTable() {
      router.push("/");
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

    return { items, isLoggedIn, routeToTable };
  },
});
</script>
