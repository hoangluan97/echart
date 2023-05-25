import ReactECharts, { EChartsOption } from 'echarts-for-react';
import React, { CSSProperties } from 'react';
import * as echarts from 'echarts';

export interface BaseChartProps {
  option: EChartsOption;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  onEvents?: Record<string, (e: any) => void>;
  width?: string | number;
  height?: string | number;
  style?: CSSProperties;
  classname?: string;
  echartRef?: any;
}


export const BaseChart: React.FC<BaseChartProps> = ({ option, width, height, onEvents, style, echartRef, ...props }) => {

  const chartHeight = height || '400px';
  const chartWidth = width || '100%';

  return  (
    <ReactECharts
      {...props}
      option={option }
      style={{ ...style, height: chartHeight, width: chartWidth , minWidth: '100%', maxWidth: '100%', zIndex: 0 }}
      onEvents={onEvents}
      ref={echartRef}
      echarts={echarts}
    />
  );
};