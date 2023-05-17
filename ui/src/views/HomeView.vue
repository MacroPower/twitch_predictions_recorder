<template>
  <div class="home">
    <PredictionHistory />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import PredictionHistory from "@/components/PredictionHistory.vue"; // @ is an alias to /src

export default defineComponent({
  name: "HomeView",
  components: {
    PredictionHistory,
  },
});
</script>
