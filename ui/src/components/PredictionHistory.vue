<template>
  <div class="title">
    <h1>History</h1>
    <n-button @click="getSummary">Refresh</n-button>
  </div>

  <n-layout embedded content-style="padding: 24px;">
    <n-space vertical>
      <n-card v-for="summary in summaries" :key="summary?.timestamp" hoverable>
        <PredictionSummary :summary="summary" />
      </n-card>
    </n-space>
  </n-layout>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";
import { useThemeVars } from "naive-ui";
import PredictionSummary from "./PredictionSummary.vue";
import Summary from "@/models/Summary";

export default defineComponent({
  name: "PredictionHistory",
  props: {},
  setup() {
    const state = reactive({
      summaries: [
        undefined,
        undefined,
        undefined,
        undefined,
        undefined,
        undefined,
        undefined,
        undefined,
      ] as Summary[] | undefined[],
    });

    return {
      ...toRefs(state),
      themeVars: useThemeVars(),
    };
  },
  components: {
    PredictionSummary,
  },
  methods: {
    async getSummary() {
      const summaries = (await Summary.get()) as Summary[];
      this.summaries = summaries.sort((a, b) => {
        return b.getDate().getTime() - a.getDate().getTime();
      });
      console.log(summaries);
    },
  },
});
</script>

<style scoped>
.title {
  text-align: center;
}
</style>
