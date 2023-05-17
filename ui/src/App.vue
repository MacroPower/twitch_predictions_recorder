<template>
  <n-config-provider :theme="darkTheme">
    <nav>
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </nav>
    <router-view />
    <n-global-style />
  </n-config-provider>
</template>

<script>
import { defineComponent } from "vue";
import { darkTheme } from "naive-ui";

export default defineComponent({
  setup() {
    return {
      darkTheme,
    };
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

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #efeff1;
}

nav a.router-link-exact-active {
  color: #bf94ff;
}
</style>
