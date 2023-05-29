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
              <n-space justify="end" align="center" class="right-nav">
                <n-button text @click="toggleTheme">
                  <template #icon><n-icon :component="themeIcon" /></template
                ></n-button>
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
import { LightModeFilled, DarkModeFilled } from "@vicons/material";

export default defineComponent({
  setup() {
    const osThemeRef = useOsTheme();

    const theme = ref<GlobalTheme | null>(
      osThemeRef.value === "dark" ? darkTheme : null
    );

    const themeIcon = ref(LightModeFilled);

    function toggleTheme() {
      if (theme.value?.name === "dark") {
        theme.value = null;
        themeIcon.value = DarkModeFilled;
      } else {
        theme.value = darkTheme;
        themeIcon.value = LightModeFilled;
      }
    }

    return {
      toggleTheme,
      themeIcon,
      theme,
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

.nav {
  padding: 4px;
}

.right-nav {
  height: 100%;
  margin-right: 18px;
}
</style>
