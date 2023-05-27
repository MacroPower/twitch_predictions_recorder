<template>
  <template v-if="outcome">
    <n-grid x-gap="0" :cols="2">
      <n-grid-item>
        <div class="prediction-outcome-stats">
          <n-space vertical :size="[0, 0]">
            <ListIcon
              :text="outcome.total_points.toLocaleString()"
              :icon="AnalyticsOutline"
            />
            <ListIcon
              :text="getReturn(outcome.total_points, totalPoints)"
              :icon="TrophyOutline"
            />
            <ListIcon
              :text="outcome.total_users.toLocaleString()"
              :icon="PeopleOutline"
            />
            <ListIcon :text="'250k'" :icon="PodiumOutline" />
          </n-space>
        </div>
      </n-grid-item>
      <n-grid-item>
        <div class="prediction-outcome-results">
          <n-space vertical>
            <PredictionOutcomeTitle
              :title="outcome.title"
              :status="getStatus(status, outcome.result_type)"
              :color="getColor(outcome.color)"
              :highlight="outcome.result_type === 'WIN'"
            />
            <span
              :style="{
                fontSize: '250%',
                color: getColor(outcome.color),
              }"
            >
              {{ getPercent(outcome.total_points, totalPoints) }}
            </span>
            <n-progress
              type="line"
              :color="getColor(outcome.color)"
              :rail-color="changeColor(getColor(outcome.color), { alpha: 0.2 })"
              :percentage="100 * (outcome.total_points / totalPoints)"
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
import Outcome from "@/models/Outcome";
import PredictionOutcomeTitle from "./PredictionOutcomeTitle.vue";
import ListIcon from "./ListIcon.vue";
import { getColor } from "@/utils/Color";

export default defineComponent({
  name: "PredictionOutcome",
  props: {
    status: {
      type: String,
      required: true,
    },
    totalPoints: {
      type: Number,
      required: true,
    },
    outcome: Outcome,
  },
  setup() {
    return {
      getColor,
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
    ListIcon,
  },
  methods: {
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
