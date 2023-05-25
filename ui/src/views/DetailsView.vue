<template>
  <div class="title">
    <h1>Details</h1>
  </div>

  <n-layout embedded content-style="padding: 24px;">
    <n-space vertical>
      <n-card>
        <PredictionSummary :summary="summary" />
      </n-card>
      <n-card>
        <PredictionDetails :details="details" />
      </n-card>
    </n-space>
  </n-layout>
</template>

<script lang="ts">
import { defineComponent, computed, reactive, toRefs } from "vue";
import PredictionSummary from "@/components/PredictionSummary.vue";
import PredictionDetails from "@/components/PredictionDetails.vue";
import { useRoute } from "vue-router";
import Summary from "@/models/Summary";
import Details from "@/models/Details";

export default defineComponent({
  name: "DetailsView",
  components: {
    PredictionSummary,
    PredictionDetails,
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

    return {
      ...toRefs(state),
      eventID: eventIDRef,
    };
  },
  methods: {
    async getSummary() {
      const summaries = (await Summary.params({
        id: this.eventID || "",
      }).get()) as Summary[];
      this.summary = summaries[0];
    },
    async getDetails() {
      const details = (await Details.params({
        id: this.eventID || "",
      }).get()) as Details[];
      details[0].sort();
      this.details = details[0];
    },
  },
  mounted() {
    this.getSummary();
    this.getDetails();
  },
});
</script>
