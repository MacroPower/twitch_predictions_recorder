<template>
  <n-grid cols="2" x-gap="24" y-gap="16">
    <n-grid-item>
      <n-space vertical size="small">
        <span v-if="summary">
          <h3>{{ summary.title }}</h3>
          <n-space vertical size="small">
            <a
              :href="'https://www.twitch.tv/' + summary.channel_name"
              target="_blank"
              v-if="summary.status === 'ACTIVE'"
            >
              <n-button secondary round type="success">
                Active (closes {{ sUseTimeAgo(summary.getEndDate()).value }})
                <template #icon>
                  <n-icon :component="PlayCircle" />
                </template>
              </n-button>
            </a>
            <a
              :href="'https://www.twitch.tv/' + summary.channel_name"
              target="_blank"
              v-else-if="summary.status === 'LOCKED'"
            >
              <n-button secondary round type="info">
                Locked (started {{ sUseTimeAgo(summary.getDate()).value }})
                <template #icon>
                  <n-icon :component="PauseCircle" />
                </template>
              </n-button>
            </a>
            <router-link v-else :to="'details/' + summary.id">
              <n-button secondary round>
                Details
                <template #icon>
                  <n-icon :component="AddCircle" />
                </template>
              </n-button>
            </router-link>
            <n-text>{{ summary.getDate().toLocaleString() }}</n-text>
          </n-space>
        </span>
        <span v-else>
          <n-space vertical size="small">
            <n-skeleton text style="width: 100%" />
            <n-skeleton text style="width: 20%" />
            <br />
            <n-skeleton text style="width: 30%" />
          </n-space>
        </span>
      </n-space>
    </n-grid-item>
    <n-grid-item>
      <n-grid cols="2" x-gap="24" y-gap="16">
        <n-grid-item
          v-for="outcome in summary?.getOutcomes() || [undefined, undefined]"
          :key="outcome?.badge_version"
        >
          <PredictionOutcome
            :total-points="summary?.outcomeSum().total_points || 0"
            :outcome="outcome"
            :status="summary?.status || ''"
          />
        </n-grid-item>
      </n-grid>
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useThemeVars } from "naive-ui";
import { changeColor } from "seemly";
import { PlayCircle, AddCircle, PauseCircle } from "@vicons/ionicons5";
import Summary from "@/models/Summary";
import PredictionOutcome from "./PredictionOutcome.vue";
import { useTimeAgo } from "@vueuse/core";

const sUseTimeAgo = (date: Date) => useTimeAgo(date, { showSecond: true });

export default defineComponent({
  name: "PredictionSummary",
  props: {
    summary: Summary,
  },
  setup() {
    return {
      changeColor,
      sUseTimeAgo,
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
