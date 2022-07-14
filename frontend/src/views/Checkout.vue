<template>
  <BaseCard>
    <BaseToolbar title="Aufteilen" icon="fa-arrows-split-up-and-left" />
    <BaseItem bgColor="c">
      <div class="field-checkbox mt-2">
        <Checkbox id="binary" v-model="checkAll" :binary="true" @click="checkAllClicked" />
        <label for="binary">Alle Auswählen</label>
      </div>
      <Divider />
      <div v-for="order in orders" :key="order.id" class="field-checkbox">
        <Checkbox :id="order.id" name="order" :value="order.id" v-model="selected" />
        <label :for="order.id">{{ order.order_item.description }}</label>
      </div>

      <Sidebar v-model:visible="checkoutModal" :baseZIndex="10000" position="full"> </Sidebar>
    </BaseItem>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref, watch } from "vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BaseItem from "@/components/UI/BaseItem.vue";
import Sidebar from "primevue/sidebar";
import { OrdersService, service_Order } from "@/services/openapi";
import InputSwitch from "primevue/inputswitch";
import Checkbox from "primevue/checkbox";
import BaseToolbar from "@/components/UI/BaseToolbar.vue";
import Divider from "primevue/divider";

export default defineComponent({
  name: "CheckoutView",
  components: { BaseToolbar, BaseItem, BaseCard, Sidebar, Checkbox, Divider },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const orders = ref<service_Order[]>([]);
    const selected = ref<number[]>([]);
    const checkoutModal = ref(false);
    const checkAll = ref(selected.value.length === orders.value.length);
    const table = computed(() => parseInt(props.id));
    watch(selected, (newValue) => {
      checkAll.value = newValue.length === orders.value.length;
    });
    onMounted(() => {
      OrdersService.getOrders(table.value, false).then((res) => {
        orders.value = res;
        initSelectedArray();
      });
    });

    function initSelectedArray() {
      orders.value.forEach((order) => order.id && selected.value.push(order.id));
    }

    function checkAllClicked() {
      checkAll.value ? (selected.value = []) : initSelectedArray();
    }

    return { checkoutModal, orders, selected, checkAll, checkAllClicked };
  },
});
</script>

<style></style>
