import { createStore } from "vuex";
import tableStore from "@/store/tables";
import orderItemStore from "@/store/orderItems";

interface AppStateModel {
  isLoggedIn: boolean;
}
export default createStore({
  state: {
    isLoggedIn: false,
  },
  getters: {
    getLoginApiStatus(state: AppStateModel) {
      return state.isLoggedIn;
    },
  },
  mutations: {
    login(state: AppStateModel) {
      state.isLoggedIn = true;
    },
    logout(state: AppStateModel) {
      state.isLoggedIn = false;
    },
  },
  actions: {},
  modules: {
    tableStore,
    orderItemStore,
  },
});
