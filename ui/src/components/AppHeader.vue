<template>
  <n-space>
    <n-menu :options="menuOptions" :value="menuValue" mode="horizontal" />
  </n-space>
</template>

<script lang="ts">
import { defineComponent, h, computed } from "vue";
import type { MenuOption } from "naive-ui";
import { RouterLink } from "vue-router";
import { useRoute } from "vue-router";

const menuOptions: MenuOption[] = [
  {
    label: () => h(RouterLink, { to: "/" }, "Home"),
    key: "home",
  },
  {
    label: () => h(RouterLink, { to: "/about" }, "About"),
    key: "about",
  },
];

export default defineComponent({
  name: "AppHeader",
  setup() {
    const route = useRoute();

    const menuValueRef = computed(() => {
      if (/\/about/.test(route.path)) return "about";
      else if (route.name === "home") return "home";
      return null;
    });

    return {
      menuOptions,
      menuValue: menuValueRef,
    };
  },
});
</script>

<style scoped></style>
