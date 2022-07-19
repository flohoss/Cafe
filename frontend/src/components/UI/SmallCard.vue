<template>
  <div v-if="!to" class="col-12 lg:col-6">
    <BaseItem class="relative" :bgColor="bgColor">
      <div class="flex justify-content-between overflow-hidden">
        <div class="flex flex-column align-items-start">
          <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis font-bold"><slot name="description"></slot></div>
          <div class="flex align-items-center mt-1">
            <TheBadge size="sm" color="info"><slot name="badgeOne"></slot></TheBadge>
            <TheBadge v-if="badgeTwo" size="sm" color="warning" class="ml-2"><slot name="badgeTwo"></slot></TheBadge>
          </div>
        </div>
        <slot name="right"></slot>
      </div>
    </BaseItem>
  </div>
  <router-link v-else class="col-12 lg:col-6 no-underline" :to="to">
    <BaseItem class="relative" :bgColor="bgColor">
      <div class="flex justify-content-between overflow-hidden">
        <div class="flex flex-column align-items-start">
          <div class="white-space-nowrap overflow-hidden text-overflow-ellipsis font-bold"><slot name="description"></slot></div>
          <div class="flex align-items-center mt-1">
            <TheBadge size="sm" color="success"><slot name="badge"></slot></TheBadge>
          </div>
        </div>
        <slot name="right"></slot>
      </div>
    </BaseItem>
  </router-link>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TheBadge from "@/components/UI/TheBadge.vue";
import BaseItem from "@/components/UI/BaseItem.vue";

export default defineComponent({
  name: "SmallCard",
  components: { TheBadge, BaseItem },
  props: {
    bgColor: { type: String, default: "a" },
    to: { type: String, default: "" },
    badgeTwo: { type: Boolean, default: true },
  },
});
</script>
