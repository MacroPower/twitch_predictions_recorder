<template>
  <div class="summary">
    <h1>History</h1>
    <button @click="getSummary">Refresh</button>
    <div class="predictionList">
      <div
        class="prediction"
        v-for="summary in summaries"
        :key="summary.timestamp"
      >
        <div class="predictionTitle">
          <h3>{{ summary.title }}</h3>
          <h5>{{ summary.status }}</h5>
        </div>
        <div
          class="predictionOutcome"
          v-for="outcome in summary.getOutcomes()"
          :key="outcome.badge_version"
        >
          <div class="predictionOutcomeDetails">
            <p>{{ outcome.total_points.toLocaleString() }}</p>
            <p>
              {{
                getReturn(
                  outcome.total_points,
                  summary.outcomeSum().total_points
                )
              }}
            </p>
            <p>{{ outcome.total_users.toLocaleString() }}</p>
            <p>250k</p>
          </div>
          <div class="predictionOutcomeSummary">
            <p :style="{ fontSize: '125%', color: getColor(outcome.color) }">
              {{ outcome.result_type || "~" }} / {{ outcome.title }}
            </p>
            <p :style="{ fontSize: '225%', color: getColor(outcome.color) }">
              {{
                getPercent(
                  outcome.total_points,
                  summary.outcomeSum().total_points
                )
              }}
            </p>
            <p>
              {{
                getPercentProgress(
                  outcome.total_points,
                  summary.outcomeSum().total_points
                )
              }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";
import Summary from "@/models/Summary";

export default defineComponent({
  name: "EventSummary",
  props: {},
  setup() {
    const state = reactive({
      summaries: [] as Summary[],
    });

    return { ...toRefs(state) };
  },
  methods: {
    async getSummary() {
      const summaries = (await Summary.get()) as Summary[];
      this.summaries = summaries.sort((a, b) => {
        return b.getDate().getTime() - a.getDate().getTime();
      });
      console.log(summaries);
    },
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
    getPercentProgress(value: number, total: number): string {
      // Returns between 1 and 10 "#" characters depending on the percentage
      // of the total value that the value represents.
      const percent = (value / total) * 100;
      const numHashes = Math.round(percent / 10);
      return "#".repeat(numHashes);
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.predictionList {
  padding: 0;
}
.predictionList > .prediction {
  margin-top: 10px;
  margin-left: auto;
  margin-right: auto;
  border-radius: 0.6rem !important;
  border-width: 1px;
  border-color: rgba(83, 83, 95, 0.48);
  border-style: solid;
  box-shadow: rgba(0, 0, 0, 0.4) 0px 4px 8px 0px;
  width: 80%;
  max-width: 1000px;
  padding: 0;
  background-color: rgb(24, 24, 27);
  display: grid;
  grid-template-columns: 2fr auto;
  grid-auto-flow: column;
}
.predictionOutcome {
  margin: 0;
  display: grid;
  grid-template-columns: 1fr 1fr;
  font-size: 0.8rem;
  line-height: 1;
  list-style-type: none;
  padding-left: 10px;
  padding-right: 10px;
  border-width: 1px;
  border-color: rgba(83, 83, 95, 0.48);
  border-left-style: solid;
  min-width: 160px;
}
.predictionOutcome p {
  line-height: 0;
  margin-top: 10px;
  margin-bottom: 10px;
}
.predictionOutcomeDetails {
  text-align: left;
  display: flex;
  justify-content: center;
  align-content: center;
  flex-direction: column;
}
.predictionOutcomeSummary {
  text-align: right;
  display: flex;
  justify-content: center;
  align-content: center;
  flex-direction: column;
}
.predictionTitle {
  margin: 0;
  line-height: 0;
  display: flex;
  justify-content: center;
  align-content: center;
  flex-direction: column;
}
</style>
