import { InjectionKey, Ref } from "vue";

export const filter = Symbol() as InjectionKey<Ref<number[]>>;
