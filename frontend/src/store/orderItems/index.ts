import { OrderItemsService, service_OrderItem } from "@/services/openapi";
import { ItemType } from "@/utils";

interface AppStateModel {
  orderItems: Map<number, service_OrderItem[]>;
}

interface mutationOrderItems {
  orderItems: service_OrderItem[];
  orderType: ItemType;
}

const orderItemStore = {
  state: () => ({
    orderItems: new Map<number, service_OrderItem[]>(),
  }),
  getters: {
    getOrderItems(state: AppStateModel) {
      return state.orderItems;
    },
  },
  mutations: {
    setOrderItems(state: AppStateModel, payload: mutationOrderItems) {
      state.orderItems.set(payload.orderType, payload.orderItems);
    },
    pushOrderItem(state: AppStateModel, orderItem: service_OrderItem) {
      const tempOrderItems = state.orderItems.get(orderItem.item_type);
      tempOrderItems && tempOrderItems.push(orderItem);
    },
    filterOrderItem(state: AppStateModel, orderItem: service_OrderItem) {
      const tempOrderItems = state.orderItems.get(orderItem.item_type);
      tempOrderItems &&
        state.orderItems.set(
          orderItem.item_type,
          tempOrderItems.filter((origItem: service_OrderItem) => origItem.id !== orderItem.id)
        );
    },
    putOrderItem(state: AppStateModel, orderItem: service_OrderItem) {
      const tempOrderItems = state.orderItems.get(orderItem.item_type);
      tempOrderItems &&
        state.orderItems.set(
          orderItem.item_type,
          tempOrderItems.map((origItem: service_OrderItem) => (origItem.id === orderItem.id ? orderItem : origItem))
        );
    },
  },
  actions: {
    async getOrderItems(context: any, orderType: ItemType) {
      const orderTypeArray = context.getters.getOrderItems;
      if (!orderTypeArray.get(orderType)) {
        const orderItems: service_OrderItem[] | null = await OrderItemsService.getOrdersItems(orderType);
        context.commit("setOrderItems", { orderItems, orderType });
      }
    },
    addOrderItem(context: any, orderItem: service_OrderItem) {
      context.commit("pushOrderItem", orderItem);
    },
    deleteOrderItem(context: any, orderItem: service_OrderItem) {
      context.commit("filterOrderItem", orderItem);
    },
    updateOrderItem(context: any, orderItem: service_OrderItem) {
      context.commit("putOrderItem", orderItem);
    },
  },
};

export default orderItemStore;
