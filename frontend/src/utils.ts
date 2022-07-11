export function convertToEur(value: number | undefined) {
  const temp: number = value ? value : 0;
  return temp.toLocaleString("de-DE", { style: "currency", currency: "EUR" });
}

export enum ItemType {
  Food,
  Drink,
}
