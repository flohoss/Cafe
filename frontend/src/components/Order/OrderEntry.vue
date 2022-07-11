<template>
  <BaseItem>
    <div class="flex justify-content-between">
      <div class="flex flex-column justify-content-between">
        <div>{{ order.order_item.description }}</div>
        <div class="text-sm">Einzelpreis: {{ convertToEur(order.order_item.price) }}</div>
      </div>
      <div class="flex align-items-center">
        <div><Button icon="pi pi-minus" class="p-button-rounded p-button-text p-button-danger" /></div>
        <div><Button icon="pi pi-plus" class="p-button-rounded p-button-text p-button-success" /></div>
      </div>
    </div>
  </BaseItem>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { service_Order } from "@/services/openapi";
import BaseItem from "@/components/UI/BaseItem.vue";
import { convertToEur } from "@/utils";
import Button from "primevue/button";

export default defineComponent({
  name: "OrderEntry",
  components: { BaseItem, Button },
  props: { order: { type: Object as PropType<service_Order>, required: true } },
  setup() {
    return { convertToEur };
  },
});
</script>
