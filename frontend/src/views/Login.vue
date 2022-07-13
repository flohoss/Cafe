<template>
  <div class="col-12 sm:w-30rem center">
    <div class="flex justify-content-center w-full mb-3">
      <img alt="logo" class="h-10rem" />
    </div>
    <InputText
      :disabled="isLoading"
      @keydown.enter="login(!v$.$invalid)"
      id="password"
      type="password"
      v-model="v$.password.$model"
      class="w-full mb-1"
      placeholder="Passwort"
      :class="{ 'p-invalid': v$.password.$invalid && submitted }"
    />
    <small v-if="(v$.password.$invalid && submitted) || v$.password.$pending.$response" class="p-error">
      {{ v$.password.required.$message.replace("Value", "Passwort") }}
    </small>
    <Button
      :loading="isLoading"
      @click="login(!v$.$invalid)"
      :label="submitLabel"
      icon="pi pi-user"
      :class="{ 'p-button-danger': (v$.$invalid && submitted) || submitLabel === 'Bitte versuchen Sie es erneut' }"
      class="p-button-primary w-full my-1"
    />
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
import { ItemType } from "@/utils";

export default defineComponent({
  name: "LoginView",
  components: { Button, InputText },
  setup() {
    const store = useStore();
    const router = useRouter();
    const isLoading = ref(false);
    const submitted = ref(false);
    const auth = reactive({ password: "" });
    const rules = { password: { required } };
    const v$ = useVuelidate(rules, auth);
    const submitLabel = ref("Einloggen");

    function login(isFormValid: boolean) {
      submitted.value = true;
      if (!isFormValid) return;
      submitLabel.value = "Laden";
      isLoading.value = true;
      AuthorizationService.postAuthLogin(auth.password)
        .then(() => {
          store.commit("login");
          router.replace({ name: "Tables" });
        })
        .catch(() => {
          submitLabel.value = "Bitte versuchen Sie es erneut";
        })
        .finally(() => {
          isLoading.value = false;
        });
    }

    return { isLoading, login, submitLabel, v$, submitted };
  },
});
</script>

<style scoped>
.center {
  position: absolute;
  top: 50%;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, -50%);
}

@media (prefers-color-scheme: light) {
  img {
    content: url("../assets/logo.png");
  }
}

@media (prefers-color-scheme: dark) {
  img {
    content: url("../assets/logo_white.png");
  }
}
</style>
