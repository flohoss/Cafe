<template>
  <BaseCard>
    <div class="grid p-fluid">
      <div class="col-12 pb-0">
        <span class="p-input-icon-left">
          <i class="pi pi-search" />
          <InputText type="search" v-model="filter" class="p-inputtext-sm" placeholder="Suchen" @keydown.esc="filter = ''" />
          <span v-if="filter !== ''" class="leftMiddle styling"><i class="pi pi-times" @click="filter = ''"></i></span>
        </span>
      </div>
    </div>
    <OrderItemCard v-for="item in filteredFood" v-bind:key="item.id" :orderItems="item">Speisen</OrderItemCard>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import { service_OrderItem } from "@/services/openapi";
import OrderItemCard from "@/components/OrderItem/OrderItemCard.vue";
import InputText from "primevue/inputtext";

export default defineComponent({
  name: "OrderItemList",
  components: { OrderItemCard, BaseCard, InputText },
  props: { orderItems: { type: Array as PropType<service_OrderItem[]>, default: () => [] } },
  setup(props) {
    const filter = ref("");
    const filteredFood = computed(() => props.orderItems.filter((item) => item.description.includes(filter.value)));
    return { filteredFood, filter };
  },
});
</script>

<style scoped>
.styling {
  cursor: pointer;
  color: gray;
  background-color: var(--surface-a);
  border-radius: 50%;
  padding: 0.2rem;
}
.leftMiddle {
  position: absolute;
  top: 50%;
  right: 0;
  transform: translate(-50%, -50%);
}
</style>
