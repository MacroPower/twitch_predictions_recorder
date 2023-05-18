<template>
  <template v-if="summary && outcome">
    <n-grid x-gap="0" :cols="2">
      <n-grid-item>
        <div class="prediction-outcome-stats">
          <n-space vertical :size="[0, 0]">
            <n-tag :bordered="false" :color="{ color: 'rgba(0,0,0,0)' }">
              {{ outcome.total_points.toLocaleString() }}
              <template #icon>
                <n-icon :component="AnalyticsOutline" />
              </template>
            </n-tag>
            <n-tag :bordered="false" :color="{ color: 'rgba(0,0,0,0)' }">
              {{
                getReturn(
                  outcome.total_points,
                  summary.outcomeSum().total_points
                )
              }}
              <template #icon>
                <n-icon :component="TrophyOutline" />
              </template>
            </n-tag>
            <n-tag :bordered="false" :color="{ color: 'rgba(0,0,0,0)' }">
              {{ outcome.total_users.toLocaleString() }}
              <template #icon>
                <n-icon :component="PeopleOutline" />
              </template>
            </n-tag>
            <n-tag :bordered="false" :color="{ color: 'rgba(0,0,0,0)' }">
              250k
              <template #icon>
                <n-icon :component="PodiumOutline" />
              </template>
            </n-tag>
          </n-space>
        </div>
      </n-grid-item>
      <n-grid-item>
        <div class="prediction-outcome-results">
          <n-space vertical>
            <PredictionOutcomeTitle
              :title="outcome.title"
              :status="getStatus(summary.status, outcome.result_type)"
              :color="getColor(outcome.color)"
              :highlight="outcome.result_type === 'WIN'"
            />
            <span
              :style="{
                fontSize: '250%',
                color: getColor(outcome.color),
              }"
            >
              {{
                getPercent(
                  outcome.total_points,
                  summary.outcomeSum().total_points
                )
              }}
            </span>
            <n-progress
              type="line"
              :color="getColor(outcome.color)"
              :rail-color="changeColor(getColor(outcome.color), { alpha: 0.2 })"
              :percentage="
                100 * (outcome.total_points / summary.outcomeSum().total_points)
              "
              :show-indicator="false"
            />
          </n-space>
        </div>
      </n-grid-item>
    </n-grid>
  </template>
  <template v-else>
    <n-space vertical size="small">
      <n-skeleton text />
      <n-skeleton text />
      <n-skeleton text />
      <n-skeleton text />
    </n-space>
  </template>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useThemeVars } from "naive-ui";
import { changeColor } from "seemly";
import {
  CheckmarkCircle,
  CloseCircle,
  RemoveCircle,
  PlayCircle,
  PauseCircle,
  AnalyticsOutline,
  TrophyOutline,
  PeopleOutline,
  PodiumOutline,
} from "@vicons/ionicons5";
import Summary from "@/models/Summary";
import Outcome from "@/models/Outcome";
import PredictionOutcomeTitle from "./PredictionOutcomeTitle.vue";

export default defineComponent({
  name: "PredictionOutcome",
  props: {
    summary: Summary,
    outcome: Outcome,
  },
  setup() {
    return {
      changeColor,
      CheckmarkCircle,
      CloseCircle,
      RemoveCircle,
      PlayCircle,
      PauseCircle,
      AnalyticsOutline,
      TrophyOutline,
      PeopleOutline,
      PodiumOutline,
      themeVars: useThemeVars(),
    };
  },
  components: {
    PredictionOutcomeTitle,
  },
  methods: {
    getColor(colorName: string): string {
      if (colorName === "BLUE") {
        return "rgb(56, 122, 255)";
      }
      if (colorName === "PINK") {
        return "rgb(245, 0, 155)";
      }
      return "white";
    },
    getPercent(value: number, total: number): string {
      return ((value / total) * 100).toFixed(0) + "%";
    },
    getReturn(value: number, total: number): string {
      return "1:" + (total / value).toFixed(2);
    },
    getStatus(summaryStatus: string, outcomeStatus: string) {
      if (summaryStatus != "RESOLVED") {
        return summaryStatus;
      }
      return outcomeStatus;
    },
  },
});
</script>

<style scoped>
.prediction-outcome-stats {
  text-align: left;
}
.prediction-outcome-results {
  text-align: right;
}
</style>
