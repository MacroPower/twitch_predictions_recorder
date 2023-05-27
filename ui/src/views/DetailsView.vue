<template>
  <div class="title">
    <h1>Details</h1>
  </div>

  <n-layout embedded content-style="padding: 24px;">
    <n-space vertical>
      <n-card>
        <PredictionSummary
          :summary="getSummarySample(details, currentSample.index)"
        />
      </n-card>
      <n-card>
        <PredictionGraph :timeseries="timeseries" :highlight="highlight" />
      </n-card>
      <n-card>
        <span v-if="details">
          <span v-for="outcome in timeseries" :key="outcome.details.timestamp">
            <p>
              {{ getTopUsersSample(details, currentSample.index) }}
            </p>
          </span>
        </span>
      </n-card>
    </n-space>
  </n-layout>
</template>

<script lang="ts">
import {
  defineComponent,
  computed,
  reactive,
  ref,
  toRefs,
  onMounted,
} from "vue";
import PredictionSummary from "@/components/PredictionSummary.vue";
import PredictionGraph from "@/components/PredictionGraph.vue";
import { useRoute } from "vue-router";
import Summary from "@/models/Summary";
import Details from "@/models/Details";
import Outcome from "@/models/Outcome";

const currentSample = ref({
  index: 0,
});

const getTopUsersSample = (d: Details, i: number) => {
  if (!d || d?.event_series.length == 0) {
    return undefined;
  }
  let currentEvent = d.event_series[d.event_series.length - 1];
  if (i >= 0 && i < d.event_series.length) {
    currentEvent = d.event_series[i];
  }
  return currentEvent.outcomes.map((e) => ({
    top: e.top_predictors.map((p) => ({
      user: p.user.user_display_name,
      points: p.points,
    })),
  }));
};

const getSummarySample = (d: Details, i: number) => {
  if (!d || d?.event_series.length == 0) {
    return undefined;
  }
  let currentEvent = d.event_series[d.event_series.length - 1];
  if (i >= 0 && i < d.event_series.length) {
    currentEvent = d.event_series[i];
  }
  const outcomes = currentEvent.outcomes.map(function (e) {
    const outcome = new Outcome();
    outcome.color = e.color;
    outcome.title = e.title;
    outcome.badge_version = e.badge_version;
    outcome.badge_set_id = e.badge_set_id;
    outcome.total_points = e.total_points;
    outcome.total_users = e.total_users;
    outcome.result_type = e.result_type;
    return outcome;
  });
  const summary = new Summary();
  summary.id = d.id;
  summary.timestamp = currentEvent.timestamp;
  summary.channel_name = d.channel_name;
  summary.created_at = d.created_at;
  summary.prediction_window_seconds = d.prediction_window_seconds;
  summary.title = d.title;
  summary.status = currentEvent.status;
  summary.outcomes = outcomes;
  return summary;
};

export default defineComponent({
  name: "DetailsView",
  components: {
    PredictionSummary,
    PredictionGraph,
  },
  setup() {
    const route = useRoute();
    const state = reactive({
      summary: undefined as Summary | undefined,
      details: undefined as Details | undefined,
    });

    const eventIDRef = computed(() => {
      const matches = route.path.match(/^\/details\/([A-Za-z0-9-]+)/);
      if (matches?.length == 2) {
        return matches[1];
      }
      return undefined;
    });

    const highlight = (opts: any) => {
      const batch = opts.batch || [{ dataIndex: -1 }];
      currentSample.value.index = batch[0].dataIndex;
      console.log(batch.map((item: any) => item.dataIndex));
    };

    async function getSummary() {
      const summaries = (await Summary.params({
        id: eventIDRef.value || "",
      }).get()) as Summary[];
      state.summary = summaries[0];
    }
    async function getDetails() {
      const details = (await Details.params({
        id: eventIDRef.value || "",
      }).get()) as Details[];
      details[0].sort();
      state.details = details[0];
    }

    onMounted(() => {
      getSummary();
      getDetails();
    });

    const timeseries = computed(() => {
      const ts = state.details?.getTimeSeries();
      return ts || [];
    });

    return {
      ...toRefs(state),
      highlight,
      getTopUsersSample,
      getSummarySample,
      getSummary,
      getDetails,
      timeseries,
      currentSample,
      eventID: eventIDRef,
    };
  },
});
</script>
