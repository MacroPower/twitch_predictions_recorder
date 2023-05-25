<template>
  <n-message-provider>
    <n-config-provider :theme="theme">
      <n-layout>
        <n-layout-header bordered class="nav">
          <n-grid x-gap="12" :cols="2">
            <n-grid-item>
              <AppHeader />
            </n-grid-item>
            <n-grid-item>
              <n-space justify="end">
                <n-button @click="theme = darkTheme"> Dark </n-button>
                <n-button @click="theme = null"> Light </n-button>
              </n-space>
            </n-grid-item>
          </n-grid>
        </n-layout-header>
        <n-layout-content>
          <router-view />
        </n-layout-content>
        <n-layout-footer bordered> Footer Footer Footer </n-layout-footer>
      </n-layout>
      <n-global-style />
    </n-config-provider>
  </n-message-provider>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useOsTheme, darkTheme } from "naive-ui";
import type { GlobalTheme } from "naive-ui";
import AppHeader from "@/components/AppHeader.vue";

export default defineComponent({
  setup() {
    const osThemeRef = useOsTheme();
    return {
      darkTheme,
      theme: ref<GlobalTheme | null>(
        osThemeRef.value === "dark" ? darkTheme : null
      ),
    };
  },
  components: {
    AppHeader,
  },
});
</script>

<style>
#app {
  font-family: v-sans, v-mono, other-fallbacks;
  font-weight: 500;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.title {
  text-align: center;
}
</style>
