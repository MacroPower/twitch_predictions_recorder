<template>
  <div class="summary">
    <h1>History</h1>
    <button @click="getSummary">Refresh</button>

    <n-layout embedded content-style="padding: 24px;">
      <n-space vertical>
        <n-card v-for="summary in summaries" :key="summary.timestamp">
          <PredictionSummary :summary="summary" />
        </n-card>
      </n-space>
    </n-layout>
  </div>
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
      summaries: [] as Summary[],
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

<style scoped></style>
