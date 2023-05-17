<template>
  <n-tag v-if="highlight" round :bordered="false" :color="{ color: color }">
    <n-ellipsis style="max-width: 80px">{{ title }}</n-ellipsis>
    <template #icon>
      <n-icon :component="CheckmarkCircle" />
    </template>
  </n-tag>
  <n-tag v-else round :bordered="false">
    <n-ellipsis style="max-width: 80px">{{ title }}</n-ellipsis>
    <template #icon>
      <n-icon :component="getIcon(status)" />
    </template>
  </n-tag>
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
} from "@vicons/ionicons5";

export default defineComponent({
  name: "PredictionOutcomeTitle",
  props: {
    title: String,
    status: String,
    color: String,
    highlight: Boolean,
  },
  setup() {
    return {
      changeColor,
      CheckmarkCircle,
      CloseCircle,
      RemoveCircle,
      PlayCircle,
      PauseCircle,
      themeVars: useThemeVars(),
    };
  },
  methods: {
    getIcon(status: string | undefined) {
      if (status === "ACTIVE") {
        return PlayCircle;
      }
      if (status === "LOCKED") {
        return PauseCircle;
      }
      if (status === "WIN") {
        return CheckmarkCircle;
      }
      if (status === "LOSE") {
        return CloseCircle;
      }
      return RemoveCircle;
    },
  },
});
</script>

<style scoped></style>
