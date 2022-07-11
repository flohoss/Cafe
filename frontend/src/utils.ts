export function convertToEur(value: number) {
  return value.toLocaleString("de-DE", { style: "currency", currency: "EUR" });
}

export enum ItemType {
  Food,
  Drink,
}
