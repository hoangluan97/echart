import { GradientLineChart } from '../../components/charts/GradientLineChart/gradientline.chart'
import { LineChart } from '../../components/charts/LineChart/line.chart'
import { ShareChart } from '../../components/charts/ShareChart/share.chart'
import "./style.css"
export const Charts = () => {
  return (
    <div className='charts-container'>
        <LineChart />     
        <GradientLineChart />  
        <ShareChart />
    </div>
  )
}