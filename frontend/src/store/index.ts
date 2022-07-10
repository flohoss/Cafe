import { createStore } from "vuex";
import { service_Table, TablesService } from "@/services/openapi";

interface AppStateModel {
  isLoggedIn: boolean;
  tables: service_Table[] | null;
}
export default createStore({
  state: {
    isLoggedIn: false,
    tables: null,
  },
  getters: {
    getLoginApiStatus(state: AppStateModel) {
      return state.isLoggedIn;
    },
    getTables(state: AppStateModel) {
      return state.tables;
    },
  },
  mutations: {
    login(state: AppStateModel) {
      state.isLoggedIn = true;
    },
    logout(state: AppStateModel) {
      state.isLoggedIn = false;
    },
    setTables(state: AppStateModel, tables: service_Table[]) {
      state.tables = tables;
    },
  },
  actions: {
    // eslint-disable-next-line
    async getTables(context: any) {
      const tables: service_Table[] | null = await TablesService.getTables();
      context.commit("setTables", tables);
    },
  },
  modules: {},
});
