<template>
  <v-chart class="chart" :option="option" @highlight="highlight" autoresize />
</template>

<script lang="ts">
import { defineComponent, ref, provide } from "vue";
import { use } from "echarts/core";
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

const option = ref({
  dataset: {} as any,
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
  console.log(opts.batch.map((item) => item.dataIndex));
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
      option,
      highlight,
    };
  },
  updated() {
    const series: DetailsSeries[] = this.details?.getTimeSeries() || [];
    const source: Array<Array<Date | number>> = [];
    const dimensions = ["timestamp"];
    series.forEach((e) => {
      dimensions.push(e.details.title + "_points");
      dimensions.push(e.details.title + "_users");
      for (let index = 0; index < e.values.length; index++) {
        const element = e.values[index];
        if (source[index]) {
          source[index].push(element.points, element.users);
        } else {
          source[index] = [element.timestamp, element.points, element.users];
        }
      }
    });

    this.option.dataset = {
      source: source,
      dimensions: dimensions,
    };

    this.option.series = dimensions.slice(1).map((e) => ({
      name: e,
      type: "line",
      encode: {
        x: "timestamp",
        y: e,
      },
    }));
  },
});
</script>

<style scoped>
.chart {
  height: 50vh;
}
</style>
