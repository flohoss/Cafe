import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import store from "@/store";
import TableView from "@/views/Tables.vue";
import LoginView from "@/views/Login.vue";
import EmptyView from "@/views/Empty.vue";
import FoodView from "@/views/Food.vue";
import DrinksView from "@/views/Drinks.vue";
import TableDetail from "@/views/TableDetail.vue";

const routes: Array<RouteRecordRaw> = [
  { path: "/tables", name: "Tables", component: TableView, meta: { needsAuth: true } },
  { path: "/tables/:id", name: "TableDetail", props: true, component: TableDetail, meta: { needsAuth: true } },
  { path: "/orders", name: "Orders", component: EmptyView, meta: { needsAuth: true } },
  { path: "/food", name: "Food", component: FoodView, meta: { needsAuth: true } },
  { path: "/drinks", name: "Drinks", component: DrinksView, meta: { needsAuth: true } },
  { path: "/bills", name: "Bills", component: EmptyView, meta: { needsAuth: true } },
  { path: "/login", name: "Login", component: LoginView },
  { path: "/:pathMatch(.*)*", redirect: { name: "Tables" } },
];

const router = createRouter({
  routes,
  history: createWebHistory(process.env.BASE_URL),
});

router.beforeEach((to, from, next) => {
  if (to.meta.needsAuth && !store.getters.getLoginApiStatus) {
    next({ name: "Login" });
  } else if (to.name === "Login" && store.getters.getLoginApiStatus) {
    next({ name: "Tables" });
  } else {
    next();
  }
});

export default router;
