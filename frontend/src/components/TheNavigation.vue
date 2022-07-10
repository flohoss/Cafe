<template>
  <Menubar v-if="isLoggedIn" :model="items" class="p-1 px-3 mb-4 shadow-1 border-0">
    <template #start>
      <img alt="logo" class="h-2-5rem mr-2" />
    </template>
    <template #end>
      <ul class="p-menubar-root-list">
        <li class="p-menuitem">
          <router-link to="/tables" class="p-menuitem-link"><i class="pi pi-table mr-2"></i>Tische</router-link>
        </li>
      </ul>
    </template>
  </Menubar>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from "vue";
import Menubar from "primevue/menubar";
import { useStore } from "vuex";

export default defineComponent({
  name: "TheNavigation",
  components: { Menubar },
  emits: ["logout"],
  setup(_, { emit }) {
    const store = useStore();
    const isLoggedIn = computed(() => store.getters.getLoginApiStatus);

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

    return { items, isLoggedIn };
  },
});
</script>

<style scoped>
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
