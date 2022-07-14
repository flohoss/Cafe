import { service_Order } from "@/services/openapi";

export function convertToEur(value: number | undefined) {
  const temp: number = value ? value : 0;
  return temp.toLocaleString("de-DE", { style: "currency", currency: "EUR" });
}

export enum ItemType {
  Food,
  Drink,
}

export interface WebSocketMsg {
  type: NotifierType;
  payload: service_Order;
}

export enum NotifierType {
  Create,
  Delete,
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
