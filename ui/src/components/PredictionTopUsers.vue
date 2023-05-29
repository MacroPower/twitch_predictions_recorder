<template>
  <n-grid x-gap="0" y-gap="12" cols="2 1200:4">
    <n-grid-item v-for="outcome in outcomes" :key="outcome.id">
      <n-h3>{{ outcome.title }}</n-h3>
      <n-ol>
        <n-li
          v-for="predictor in outcome.top_predictors"
          :key="predictor.user.user_id"
        >
          {{ predictor.user.user_display_name }}:
          {{ predictor.points.toLocaleString() }} ->
          {{
            Math.floor(
              predictor.points * (getTotalPoints() / outcome.total_points)
            ).toLocaleString()
          }}
        </n-li>
      </n-ol>
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts">
import { defineComponent, PropType, toRefs } from "vue";
import Outcome, { getOutcomeSum } from "@/models/Outcome";

export default defineComponent({
  name: "PredictionTopUsers",
  props: {
    outcomes: {
      type: Array as PropType<Outcome[]>,
      required: true,
    },
  },
  setup(props) {
    const { outcomes: outcomesRef } = toRefs(props);

    const getTotalPoints = () => {
      return getOutcomeSum(outcomesRef.value).total_points;
    };

    return {
      outcomesRef,
      getTotalPoints,
    };
  },
});
</script>
