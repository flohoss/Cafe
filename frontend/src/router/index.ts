import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "@/views/Home.vue";
import LoginView from "@/views/Login.vue";
import store from "@/store";

const routes: Array<RouteRecordRaw> = [
  { path: "/", name: "Home", component: HomeView, meta: { needsAuth: true } },
  { path: "/login", name: "Login", component: LoginView },
  { path: "/:pathMatch(.*)*", redirect: "/" },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
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
