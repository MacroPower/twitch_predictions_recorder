<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <EventSummary />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import EventSummary from "@/components/EventSummary.vue"; // @ is an alias to /src

export default defineComponent({
  name: "HomeView",
  components: {
    EventSummary,
  },
});
</script>
