import { service_Bill, service_Order } from "@/services/openapi";

export function convertToEur(value: number | undefined) {
  const temp: number = value ? value : 0;
  return temp.toLocaleString("de-DE", { style: "currency", currency: "EUR" });
}

export enum ItemType {
  Food,
  ColdDrink,
  HotDrink,
}

export function detailedItemTypeString(type: ItemType) {
  switch (type) {
    case ItemType.Food:
      return "Speisen";
    case ItemType.ColdDrink:
      return "Kaltgetränke";
    default:
      return "Heißgetränke";
  }
}

export function generalItemTypeString(type: ItemType[]) {
  if (type.includes(ItemType.Food)) {
    return "Speisen";
  } else {
    return "Getränke";
  }
}

export function detailedItemTypeIcon(type: ItemType) {
  switch (type) {
    case ItemType.Food:
      return "fa-solid fa-cheese";
    case ItemType.ColdDrink:
      return "fa-solid fa-champagne-glasses";
    default:
      return "fa-solid fa-mug-hot";
  }
}

export function generalItemTypeIcon(type: ItemType[]) {
  if (type.includes(ItemType.Food)) {
    return "fa-solid fa-cheese";
  } else {
    return "fa-solid fa-champagne-glasses";
  }
}

export interface WebSocketMsg {
  type: NotifierType;
  payload: service_Order[];
}

export enum NotifierType {
  Create,
  Delete,
  DeleteAll,
}

import { ToastServiceMethods } from "primevue/toastservice";
import moment from "moment";

const timeToLife = 3600;

export function errorToast(toast: ToastServiceMethods, message: string) {
  toast.removeAllGroups();
  toast.add({ severity: "error", summary: "Fehler", detail: message, group: "br", life: timeToLife });
}

export function getCurrentTimeSince(updated_at: number | undefined) {
  return updated_at ? moment.unix(updated_at).fromNow() : "";
}

export const emptyBill: service_Bill = { table_id: 0, total: 0 };
