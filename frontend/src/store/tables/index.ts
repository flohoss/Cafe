import { service_Table, TablesService } from "@/services/openapi";

interface AppStateModel {
  tables: service_Table[] | null;
}

const tableStore = {
  state: () => ({
    tables: null,
  }),
  getters: {
    getTables(state: AppStateModel) {
      return state.tables;
    },
  },
  mutations: {
    setTables(state: AppStateModel, tables: service_Table[]) {
      state.tables = tables;
    },
    popTables(state: AppStateModel) {
      state.tables?.pop();
    },
    pushTable(state: AppStateModel, table: service_Table) {
      state.tables?.push(table);
    },
  },
  actions: {
    // eslint-disable-next-line
    async getTables(context: any) {
      const tables: service_Table[] | null = await TablesService.getTables();
      context.commit("setTables", tables);
    },
    // eslint-disable-next-line
    removeLastTable(context: any) {
      context.commit("popTables");
    },
    // eslint-disable-next-line
    addTable(context: any, table: service_Table) {
      context.commit("pushTable", table);
    },
  },
};

export default tableStore;
