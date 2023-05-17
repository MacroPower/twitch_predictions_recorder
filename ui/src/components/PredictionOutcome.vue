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
            <n-tag v-if="summary.status === 'ACTIVE'" round :bordered="false">
              {{ outcome.title }}
              <template #icon>
                <n-icon :component="PlayCircle" />
              </template>
            </n-tag>
            <n-tag
              v-else-if="summary.status === 'LOCKED'"
              round
              :bordered="false"
            >
              {{ outcome.title }}
              <template #icon>
                <n-icon :component="PauseCircle" />
              </template>
            </n-tag>
            <n-tag
              v-else-if="outcome.result_type === 'WIN'"
              round
              :bordered="false"
              :color="{ color: getColor(outcome.color) }"
            >
              {{ outcome.title }}
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
              </template>
            </n-tag>
            <n-tag
              v-else-if="outcome.result_type === 'LOSE'"
              round
              :bordered="false"
            >
              {{ outcome.title }}
              <template #icon>
                <n-icon :component="CloseCircle" />
              </template>
            </n-tag>
            <n-tag v-else round :bordered="false">
              {{ outcome.title }}
              <template #icon>
                <n-icon :component="RemoveCircle" />
              </template>
            </n-tag>
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
    <n-skeleton text :repeat="2" /> <n-skeleton text style="width: 60%" />
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
