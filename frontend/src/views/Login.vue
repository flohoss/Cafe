<template>
  <div class="flex justify-content-center align-items-center h-screen">
    <BaseCard class="col-12 md:col-6">
      <div class="flex justify-content-center w-full mb-3">
        <img class="w-10rem" src="../assets/logo.png" alt="" />
      </div>
      <InputText
        :disabled="isLoading"
        @keydown.enter="login(!v$.$invalid)"
        id="password"
        type="password"
        v-model="v$.password.$model"
        class="w-full mb-1"
        placeholder="Password"
        :class="{ 'p-invalid': v$.password.$invalid && submitted }"
      />
      <small v-if="(v$.password.$invalid && submitted) || v$.password.$pending.$response" class="p-error">
        {{ v$.password.required.$message.replace("Value", "Password") }}
      </small>
      <Button
        :loading="isLoading"
        @click="login(!v$.$invalid)"
        :label="submitLabel"
        icon="pi pi-user"
        :class="{ 'p-button-danger': (v$.$invalid && submitted) || submitLabel === 'Please try again' }"
        class="p-button-primary w-full my-1"
      />
    </BaseCard>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { AuthorizationService } from "@/services/openapi";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { required } from "@vuelidate/validators";
import { useVuelidate } from "@vuelidate/core";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import BaseCard from "@/components/BaseCard.vue";

export default defineComponent({
  name: "LoginView",
  components: { BaseCard, Button, InputText },
  setup() {
    const store = useStore();
    const router = useRouter();
    const isLoading = ref(false);
    const submitted = ref(false);
    const auth = reactive({ password: "" });
    const rules = { password: { required } };
    const v$ = useVuelidate(rules, auth);
    const submitLabel = ref("Sign In");

    function login(isFormValid: boolean) {
      submitted.value = true;
      if (!isFormValid) return;
      submitLabel.value = "Loading";
      isLoading.value = true;
      AuthorizationService.postAuthLogin(auth.password)
        .then(() => {
          store.dispatch("getTables").then(() => {
            store.commit("login");
            router.replace({ name: "Tables" });
          });
        })
        .catch(() => {
          submitLabel.value = "Please try again";
        })
        .finally(() => {
          isLoading.value = false;
        });
    }

    return { isLoading, login, submitLabel, v$, submitted };
  },
});
</script>
