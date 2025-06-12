<template>
  <div ref="chartRef" :style="{ width, height }"></div>
</template>

<script setup lang="ts">
// 引入 echarts 核心模块，核心模块提供了 echarts 使用必须要的接口。
import * as echarts from "echarts/core";
// 引入柱状、折线和饼图常用图表
import { BarChart, LineChart, PieChart } from "echarts/charts";
// 引入标题，提示框，直角坐标系，数据集，内置数据转换器组件，
import { GridComponent, TooltipComponent, LegendComponent } from "echarts/components";
// 引入 Canvas 渲染器，注意引入 CanvasRenderer 或者 SVGRenderer 是必须的一步
import { CanvasRenderer } from "echarts/renderers";

import { useResizeObserver } from "@vueuse/core";

// 按需注册组件
echarts.use([
  CanvasRenderer,
  BarChart,
  LineChart,
  PieChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
]);

const props = defineProps<{
  options: echarts.EChartsCoreOption;
  width?: string;
  height?: string;
}>();

const chartRef = ref<HTMLDivElement | null>(null);
let chartInstance: echarts.ECharts | null = null;

// 初始化图表
const initChart = () => {
  if (chartRef.value) {
    chartInstance = echarts.init(chartRef.value);
    if (props.options) {
      chartInstance.setOption(props.options);
    }
  }
};

// 监听尺寸变化，自动调整
useResizeObserver(chartRef, () => {
  chartInstance?.resize();
});

// 监听 options 变化，更新图表
watch(
  () => props.options,
  (newOptions) => {
    if (chartInstance && newOptions) {
      chartInstance.setOption(newOptions);
    }
  },
  { deep: true }
);

onMounted(() => {
  nextTick(() => initChart());
});

onBeforeUnmount(() => {
  chartInstance?.dispose();
});
</script>
