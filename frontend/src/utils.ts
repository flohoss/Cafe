export function convertToEur(value: number) {
  return value.toLocaleString("de-DE", { style: "currency", currency: "EUR" });
}
