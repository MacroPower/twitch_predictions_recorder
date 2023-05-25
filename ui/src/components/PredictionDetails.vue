<template>
  <n-grid x-gap="0" :cols="2">
    <n-grid-item>
      <h3 class="title">Points</h3>
      <v-chart
        class="chart"
        :option="optionPoints"
        group="c"
        @highlight="highlight"
        autoresize
      />
    </n-grid-item>
    <n-grid-item>
      <h3 class="title">Users</h3>
      <v-chart
        class="chart"
        :option="optionUsers"
        group="c"
        @highlight="highlight"
        autoresize
      />
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts">
import { defineComponent, ref, provide } from "vue";
import { use, connect } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { LineChart } from "echarts/charts";
import {
  GridComponent,
  PolarComponent,
  GeoComponent,
  SingleAxisComponent,
  ParallelComponent,
  CalendarComponent,
  GraphicComponent,
  ToolboxComponent,
  TooltipComponent,
  AxisPointerComponent,
  BrushComponent,
  TitleComponent,
  TimelineComponent,
  MarkPointComponent,
  MarkLineComponent,
  MarkAreaComponent,
  LegendComponent,
  DataZoomComponent,
  DataZoomInsideComponent,
  DataZoomSliderComponent,
  VisualMapComponent,
  VisualMapContinuousComponent,
  VisualMapPiecewiseComponent,
  AriaComponent,
  DatasetComponent,
  TransformComponent,
} from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";
import Details, { DetailsSeries } from "@/models/Details";
import { getColor } from "@/utils/Color";

use([
  CanvasRenderer,
  LineChart,
  GridComponent,
  PolarComponent,
  GeoComponent,
  SingleAxisComponent,
  ParallelComponent,
  CalendarComponent,
  GraphicComponent,
  ToolboxComponent,
  TooltipComponent,
  AxisPointerComponent,
  BrushComponent,
  TitleComponent,
  TimelineComponent,
  MarkPointComponent,
  MarkLineComponent,
  MarkAreaComponent,
  LegendComponent,
  DataZoomComponent,
  DataZoomInsideComponent,
  DataZoomSliderComponent,
  VisualMapComponent,
  VisualMapContinuousComponent,
  VisualMapPiecewiseComponent,
  AriaComponent,
  DatasetComponent,
  TransformComponent,
]);

provide(THEME_KEY, "dark");

const optionPoints = ref({
  color: [] as string[],
  legend: {
    data: [] as string[],
  },
  dataset: {} as { source: (number | Date)[][]; dimensions: string[] },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      animation: false,
    },
  },
  xAxis: { type: "time" },
  yAxis: {},
  series: [] as any[],
});

const optionUsers = ref({
  color: [] as string[],
  legend: {
    data: [] as string[],
  },
  dataset: {} as { source: (number | Date)[][]; dimensions: string[] },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      animation: false,
    },
  },
  xAxis: { type: "time" },
  yAxis: {},
  series: [] as any[],
});

const highlight = (opts: HighlightEvent) => {
  console.log(opts.batch?.map((item) => item.dataIndex));
};

interface HighlightEvent {
  batch: HighlightBatch[];
}

interface HighlightBatch {
  dataIndex: number;
  dataIndexInside: number;
  seriesIndex: number;
}

export default defineComponent({
  name: "PredictionDetails",
  props: {
    details: Details,
  },
  components: {
    VChart,
  },
  setup() {
    return {
      optionPoints,
      optionUsers,
      highlight,
    };
  },
  mounted() {
    connect("c");
  },
  updated() {
    const series: DetailsSeries[] = this.details?.getTimeSeries() || [];
    const dimensions: string[] = [];
    const colors: string[] = [];
    const sourcePoints: Array<Array<Date | number>> = [];
    const sourceUsers: Array<Array<Date | number>> = [];
    series.forEach((e) => {
      dimensions.push(e.details.title);
      colors.push(getColor(e.details.color));
      for (let index = 0; index < e.values.length; index++) {
        const element = e.values[index];
        if (element.status != "ACTIVE") {
          continue;
        }
        if (sourcePoints[index]) {
          sourcePoints[index].push(element.points);
          sourceUsers[index].push(element.users);
        } else {
          sourcePoints[index] = [element.timestamp, element.points];
          sourceUsers[index] = [element.timestamp, element.users];
        }
      }
    });

    this.optionPoints.dataset = {
      source: sourcePoints,
      dimensions: ["timestamp", ...dimensions],
    };
    this.optionUsers.dataset = {
      source: sourceUsers,
      dimensions: ["timestamp", ...dimensions],
    };

    const outcomeDimensions = dimensions.map((e) => ({
      name: e,
      type: "line",
      encode: {
        x: "timestamp",
        y: e,
      },
      smooth: true,
      symbol: "none",
    }));

    this.optionPoints.series = outcomeDimensions;
    this.optionUsers.series = outcomeDimensions;

    this.optionPoints.color = colors;
    this.optionUsers.color = colors;

    this.optionPoints.legend.data = dimensions;
    this.optionUsers.legend.data = dimensions;
  },
});
</script>

<style scoped>
.chart {
  height: 30vh;
}

.title {
  text-align: center;
}
</style>
