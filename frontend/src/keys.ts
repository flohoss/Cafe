import { InjectionKey, Ref } from "vue";

export const filter = Symbol() as InjectionKey<Ref<number[]>>;
export const loading = Symbol() as InjectionKey<Ref<boolean>>;
