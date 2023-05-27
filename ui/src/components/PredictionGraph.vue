<template>
  <n-grid x-gap="0" :cols="2">
    <n-grid-item>
      <h3 class="title">Points</h3>
      <v-chart
        class="chart"
        :option="getOptions(timeseriesRef, 'points')"
        group="c"
        @highlight="highlight"
        autoresize
      />
    </n-grid-item>
    <n-grid-item>
      <h3 class="title">Users</h3>
      <v-chart
        class="chart"
        :option="getOptions(timeseriesRef, 'users')"
        group="c"
        @highlight="highlight"
        autoresize
      />
    </n-grid-item>
  </n-grid>
</template>

<script lang="ts">
import { defineComponent, provide, onMounted, toRefs, PropType } from "vue";
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
import { DetailsSeries } from "@/models/Details";
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

type Options = {
  color: string[];
  legend: {
    data: string[];
  };
  dataset: { source: (number | Date)[][]; dimensions: string[] } | object;
  tooltip: {
    trigger: string;
    axisPointer: {
      animation: boolean;
    };
  };
  xAxis: { type: string } | object;
  yAxis: { type: string } | object;
  series: any[];
};

interface HighlightEvent {
  batch: HighlightBatch[];
}

interface HighlightBatch {
  dataIndex: number;
  dataIndexInside: number;
  seriesIndex: number;
}

type HighlightCallback = (event: HighlightEvent) => void;

export default defineComponent({
  name: "PredictionGraph",
  props: {
    timeseries: {
      type: Array as PropType<DetailsSeries[]>,
      required: true,
    },
    highlight: Object as PropType<HighlightCallback>,
  },
  components: {
    VChart,
  },
  setup(props) {
    const { timeseries: timeseriesRef } = toRefs(props);
    onMounted(() => {
      connect("c");
    });

    return {
      timeseriesRef,
    };
  },
  methods: {
    getOptions(ts: DetailsSeries[], dimension: string) {
      const options: Options = {
        color: [],
        legend: {
          data: [],
        },
        dataset: {},
        tooltip: {
          trigger: "axis",
          axisPointer: {
            animation: false,
          },
        },
        xAxis: { type: "time" },
        yAxis: {},
        series: [],
      };
      const dimensions: string[] = [];
      const colors: string[] = [];
      const source: Array<Array<Date | number>> = [];
      const getDimension = (e: { points: number; users: number }) => {
        if (dimension == "points") {
          return e.points;
        }
        if (dimension == "users") {
          return e.users;
        }
        return 0;
      };
      ts.forEach((e) => {
        dimensions.push(e.details.title);
        colors.push(getColor(e.details.color));
        for (let index = 0; index < e.values.length; index++) {
          const element = e.values[index];
          // if (element.status != "ACTIVE") {
          //   continue;
          // }
          if (source[index]) {
            source[index].push(getDimension(element));
          } else {
            source[index] = [element.timestamp, getDimension(element)];
          }
        }
      });
      options.dataset = {
        source: source,
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
      options.series = outcomeDimensions;
      options.color = colors;
      options.legend.data = dimensions;

      return options;
    },
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
