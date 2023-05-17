<template>
  <template v-if="summary">
    <n-grid cols="2" x-gap="24" y-gap="16">
      <n-grid-item>
        <n-space vertical size="small">
          <span>
            <h3>{{ summary.title }}</h3>
            <a
              :href="'https://www.twitch.tv/' + summary.channel_name"
              target="_blank"
              v-if="summary.status === 'ACTIVE'"
            >
              <n-button secondary round type="success">
                Active ({{ summary.getRemainingTime().toFixed(0) }}s)
                <template #icon>
                  <n-icon :component="PlayCircle" />
                </template>
              </n-button>
            </a>
            <a
              :href="'https://www.twitch.tv/' + summary.channel_name"
              target="_blank"
              v-if="summary.status === 'LOCKED'"
            >
              <n-button secondary round type="info">
                Locked
                <template #icon>
                  <n-icon :component="PauseCircle" />
                </template>
              </n-button>
            </a>
            <n-button secondary round v-else>
              Details
              <template #icon>
                <n-icon :component="AddCircle" />
              </template>
            </n-button>
          </span>
          <span>{{ summary.getDate().toLocaleString() }}</span>
        </n-space>
      </n-grid-item>
      <n-grid-item>
        <n-grid cols="2" x-gap="24" y-gap="16">
          <n-grid-item
            v-for="outcome in summary.getOutcomes()"
            :key="outcome.badge_version"
          >
            <PredictionOutcome :summary="summary" :outcome="outcome" />
          </n-grid-item>
        </n-grid>
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
import { PlayCircle, AddCircle, PauseCircle } from "@vicons/ionicons5";
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
      AddCircle,
      PauseCircle,
      themeVars: useThemeVars(),
    };
  },
  components: {
    PredictionOutcome,
  },
});
</script>

<style scoped></style>
