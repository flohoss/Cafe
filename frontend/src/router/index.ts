import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import TableView from "@/views/Tables.vue";
import LoginView from "@/views/Login.vue";
import EmptyView from "@/views/Empty.vue";
import store from "@/store";

const routes: Array<RouteRecordRaw> = [
  { path: "/", name: "Tables", component: TableView, meta: { needsAuth: true } },
  { path: "/orders", name: "Orders", component: EmptyView, meta: { needsAuth: true } },
  { path: "/foods", name: "Foods", component: EmptyView, meta: { needsAuth: true } },
  { path: "/drinks", name: "Drinks", component: EmptyView, meta: { needsAuth: true } },
  { path: "/bills", name: "Bills", component: EmptyView, meta: { needsAuth: true } },
  { path: "/login", name: "Login", component: LoginView },
  { path: "/:pathMatch(.*)*", redirect: "/" },
];

const router = createRouter({
  routes,
  history: createWebHistory(process.env.BASE_URL),
  linkActiveClass: "p-menuitem-active",
  linkExactActiveClass: "p-menuitem-active",
});

router.beforeEach((to, from, next) => {
  if (to.meta.needsAuth && !store.getters.getLoginApiStatus) {
    next("/login");
  } else if (to.name === "Login" && store.getters.getLoginApiStatus) {
    next("/");
  } else {
    next();
  }
});

export default router;
