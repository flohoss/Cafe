import { InjectionKey, Ref } from "vue";

export const loading = Symbol() as InjectionKey<Ref<boolean>>;
