<template>
  <template v-if="summary">
    <n-grid cols="4" x-gap="24" y-gap="16">
      <n-grid-item :span="2">
        <n-space vertical size="small">
          <span>
            <h3>{{ summary.title }}</h3>
            <a
              :href="'https://www.twitch.tv/' + summary.channel_name"
              target="_blank"
            >
              <n-button
                secondary
                round
                type="success"
                v-if="summary.status === 'ACTIVE'"
              >
                Active ({{ summary.getRemainingTime().toFixed(0) }}s)
                <template #icon>
                  <n-icon :component="PlayCircle" />
                </template>
              </n-button>
            </a>
          </span>
          <span>{{ summary.getDate().toLocaleString() }}</span>
        </n-space>
      </n-grid-item>
      <n-grid-item
        v-for="outcome in summary.getOutcomes()"
        :key="outcome.badge_version"
      >
        <PredictionOutcome :summary="summary" :outcome="outcome" />
      </n-grid-item>
    </n-grid>
  </template>
  <template v-else>
    <n-skeleton text :repeat="2" /> <n-skeleton text style="width: 60%" />
  </template>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useThemeVars } from "naive-ui";
import { changeColor } from "seemly";
import { PlayCircle } from "@vicons/ionicons5";
import Summary from "@/models/Summary";
import PredictionOutcome from "./PredictionOutcome.vue";

export default defineComponent({
  name: "PredictionSummary",
  props: {
    summary: Summary,
  },
  setup() {
    return {
      changeColor,
      PlayCircle,
      themeVars: useThemeVars(),
    };
  },
  components: {
    PredictionOutcome,
  },
});
</script>

<style scoped></style>
