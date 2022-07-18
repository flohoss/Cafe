<template>
  <BaseCard>
    <Transition>
      <WaveSpinner v-if="isLoading" />
      <div v-else>
        <BaseItem>
          <div class="field-checkbox mt-1">
            <Checkbox id="binary" v-model="checkAll" :binary="true" @click="checkAllClicked" />
            <label for="binary">Alle Auswählen</label>
          </div>
          <hr style="color: var(--text-color)" class="my-3" />
          <div v-for="order in orders" :key="order.id" class="field-checkbox">
            <Checkbox :id="order.id" name="order" :value="order.id" v-model="orderFilter" />
            <label :for="order.id" class="w-full">
              <span class="flex justify-content-between">
                <span class="overflow-hidden text-overflow-ellipsis white-space-nowrap">{{ order.order_item.description }}</span>
                <span>{{ convertToEur(order.order_item.price) }}</span>
              </span>
            </label>
          </div>
        </BaseItem>
        <div class="h-4rem"></div>

        <BottomNavigation>
          <template #left>
            <router-link :to="{ name: 'TableDetail' }" class="no-underline">
              <Button :disabled="applyFilterLoading" icon="pi pi-arrow-left" class="p-button-rounded" />
            </router-link>
          </template>
          <template #middle>
            <div class="flex flex-column align-items-center"></div>
          </template>
          <template #right>
            <Button
              :disabled="deleteFilterLoading"
              :loading="applyFilterLoading"
              icon="pi pi-check"
              class="p-button-success p-button-rounded"
              @click="generateBill"
            />
          </template>
        </BottomNavigation>
      </div>
    </Transition>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from "vue";
import { OrdersService, service_Order } from "@/services/openapi";
import Checkbox from "primevue/checkbox";
import { convertToEur } from "@/utils";
import Button from "primevue/button";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import BaseItem from "@/components/UI/BaseItem.vue";

export default defineComponent({
  name: "CheckoutView",
  components: { BaseItem, BaseCard, WaveSpinner, Checkbox, Button, BottomNavigation },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const table = computed(() => parseInt(props.id));
    const orders = ref<service_Order[]>([]);
    const orderFilter = ref<number[]>([]);
    const isLoading = ref(false);
    const applyFilterLoading = ref(false);
    const deleteFilterLoading = ref(false);
    const checkAll = ref(false);

    function checkAllCheck() {
      if (orderFilter.value) checkAll.value = orderFilter.value.length === orders.value.length;
    }

    watch(orderFilter, () => checkAllCheck());
    getData();

    function getData() {
      isLoading.value = true;
      OrdersService.getOrders(table.value, false)
        .then((res) => {
          orders.value = res;
          setAllOrdersSelected();
          checkAllCheck();
        })
        .finally(() => (isLoading.value = false));
    }

    function setAllOrdersSelected() {
      const temp: number[] = [];
      orders.value.forEach((order) => order.id && temp.push(order.id));
      orderFilter.value = temp;
    }

    function checkAllClicked() {
      if (orderFilter.value.length === orders.value.length) {
        orderFilter.value = [];
      } else {
        setAllOrdersSelected();
      }
    }

    function generateBill() {
      applyFilterLoading.value = true;
    }

    return { orders, orderFilter, checkAll, checkAllClicked, convertToEur, generateBill, isLoading, applyFilterLoading, deleteFilterLoading };
  },
});
</script>

<style scoped>
.field-checkbox:last-child {
  margin-bottom: 0.25rem;
}
</style>
