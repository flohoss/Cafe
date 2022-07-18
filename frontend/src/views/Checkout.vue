<template>
  <BaseCard>
    <ConfirmDialog></ConfirmDialog>
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
            <div class="flex flex-column align-items-center">
              <div class="text-sm">Tisch {{ table }}</div>
              <div class="font-bold">{{ convertToEur(total) }}</div>
            </div>
          </template>
          <template #right>
            <Button
              :disabled="total === 0"
              :loading="applyFilterLoading"
              icon="pi pi-money-bill"
              class="p-button-danger p-button-rounded"
              @click="generateBill"
            />
          </template>
        </BottomNavigation>
      </div>
    </Transition>

    <Sidebar v-model:visible="billModal" :baseZIndex="10000" position="full" @hide="billModalClosed">
      <BillModal :bill="bill" />
    </Sidebar>
  </BaseCard>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from "vue";
import { BillsService, OrdersService, service_Bill, service_Order } from "@/services/openapi";
import Checkbox from "primevue/checkbox";
import { convertToEur, emptyBill, errorToast } from "@/utils";
import Button from "primevue/button";
import WaveSpinner from "@/components/UI/WaveSpinner.vue";
import BaseCard from "@/components/UI/BaseCard.vue";
import BottomNavigation from "@/components/UI/BottomNavigation.vue";
import BaseItem from "@/components/UI/BaseItem.vue";
import ConfirmDialog from "primevue/confirmdialog";
import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import Sidebar from "primevue/sidebar";
import BillModal from "@/components/Bills/BillModal.vue";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "CheckoutView",
  components: { BaseItem, BaseCard, WaveSpinner, Checkbox, Button, BottomNavigation, ConfirmDialog, Sidebar, BillModal },
  props: { id: { type: String, default: "0" } },
  setup(props) {
    const confirm = useConfirm();
    const toast = useToast();
    const router = useRouter();
    const table = computed(() => parseInt(props.id));
    const orders = ref<service_Order[]>([]);
    const orderFilter = ref<number[]>([]);
    const isLoading = ref(false);
    const applyFilterLoading = ref(false);
    const checkAll = ref(false);
    const total = ref(0);
    const bill = ref<service_Bill>({ ...emptyBill });
    const billModal = ref(false);

    function checkAllCheck() {
      if (orderFilter.value) checkAll.value = orderFilter.value.length === orders.value.length;
    }

    function calculateTotal() {
      let temp = 0;
      orders.value.forEach((order) => {
        if (order.id && orderFilter.value.includes(order.id)) temp += order.order_item.price;
      });
      total.value = temp;
    }

    watch(orderFilter, () => {
      checkAllCheck();
      calculateTotal();
    });

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
      confirm.require({
        message: "Alle ausgewählte Bestellungen abrechnen?",
        header: "Abrechnen",
        icon: "pi pi-exclamation-triangle",
        acceptClass: "p-button-danger",
        rejectClass: "p-button-secondary",
        accept: () => {
          BillsService.postBills(table.value, orderFilter.value.toString())
            .then((res) => {
              bill.value = res;
              billModal.value = true;
              getData();
            })
            .catch((err) => errorToast(toast, err.body.error))
            .finally(() => (applyFilterLoading.value = false));
        },
        reject: () => {
          applyFilterLoading.value = false;
        },
      });
    }

    function billModalClosed() {
      if (orderFilter.value.length === 0) {
        router.push({ name: "Bills" });
      }
    }

    return {
      orders,
      orderFilter,
      checkAll,
      checkAllClicked,
      convertToEur,
      generateBill,
      isLoading,
      applyFilterLoading,
      total,
      table,
      bill,
      billModal,
      billModalClosed,
    };
  },
});
</script>

<style scoped>
.field-checkbox:last-child {
  margin-bottom: 0.25rem;
}
</style>
