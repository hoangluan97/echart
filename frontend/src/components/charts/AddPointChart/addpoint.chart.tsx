import { EChartsOption } from "echarts-for-react";
import React from "react";
import { BaseChart } from "../BaseChart/base.chart";

var option: EChartsOption;

export const AddPointChart: React.FC = () => {
  const eChartsRef = React.useRef(null as any);
  
  const symbolSize = 20;
  const data = [
    [15, 0],
    [-50, 10],
    [-56.5, 20],
    [-46.5, 30],
    [-22.1, 40]
  ];

  option = {
    title: {
      text: 'Click to Add Points'
    },
    tooltip: {
      formatter: function (params) {
        var data = params.data || [0, 0];
        return data[0].toFixed(2) + ', ' + data[1].toFixed(2);
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      min: -60,
      max: 20,
      type: 'value',
      axisLine: { onZero: false }
    },
    yAxis: {
      min: 0,
      max: 40,
      type: 'value',
      axisLine: { onZero: false }
    },
    series: [
      {
        id: 'a',
        type: 'line',
        smooth: true,
        symbolSize: symbolSize,
        data: data
      }
    ]
  };

  const handleChartClick = (param) => {
    console.log("click", param);
    let myChart

    if (eChartsRef && eChartsRef.current) {
      myChart = eChartsRef.current?.getEchartsInstance();
      console.log("click", myChart);
      const pointInGrid = myChart.convertFromPixel('grid', [
        param.offsetX,
        param.offsetY
      ]);
      if (myChart.containPixel('grid', [param.offsetX, param.offsetY])) {
        data.push(pointInGrid);
        myChart.setOption({
          series: [
            {
              id: 'a',
              data: data
            }
          ]
        });
      }
    }
  };

  
  const handleChartMouseMove = (param) => {
    let myChart
    const pointInPixel = [param.offsetX, param.offsetY];
    if (eChartsRef && eChartsRef.current) {
      myChart = eChartsRef.current?.getEchartsInstance();
    myChart.getZr().setCursorStyle(
      myChart.containPixel('grid', pointInPixel) ? 'copy' : 'default'
    );
  }};
  const onEvents = {
    click: (param) => handleChartClick(param),
    mousemove: (param) => handleChartMouseMove(param)
  }
  
  return (
    <BaseChart echartRef={eChartsRef} onEvents={onEvents} option={option} />
  );
};
