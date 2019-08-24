import React from "react";
import PropTypes from "prop-types";

import { scaleTime } from "d3-scale";
import { utcDay, utcMinute, timeMinute } from "d3-time";
import { format } from "d3-format";
import { timeFormat } from "d3-time-format";

import { ChartCanvas, Chart } from "react-stockcharts";
import { LineSeries, CandlestickSeries } from "react-stockcharts/lib/series";
import { XAxis, YAxis } from "react-stockcharts/lib/axes";
import { CrossHairCursor, CurrentCoordinate, MouseCoordinateX, MouseCoordinateY } from "react-stockcharts/lib/coordinates";
import { fitWidth } from "react-stockcharts/lib/helper";
import { last, timeIntervalBarWidth } from "react-stockcharts/lib/utils";
import { discontinuousTimeScaleProvider } from "react-stockcharts/lib/scale";
import { ema, wma, sma, tma } from "react-stockcharts/lib/indicator";
import {
	OHLCTooltip,
	MovingAverageTooltip,
} from "react-stockcharts/lib/tooltip";


window.scaleTime = scaleTime

class CandleStickChart extends React.Component {
	render() {
		const { type, width, data, ratio } = this.props;
		const xAccessor = d => d.date;
		const xExtents = [
			xAccessor(last(data)),
			xAccessor(data[data.length - 200])
		];

    const sma20 = sma()
			.options({ windowSize: 20 })
			.merge((d, c) => {d.sma20 = c;})
			.accessor(d => d.sma20);

    const calculatedData = sma20(data)
		const xScaleProvider = discontinuousTimeScaleProvider.inputDateAccessor(d => d.date);
    /*
		const {
			data,
			xScale,
			xAccessor,
			displayXAccessor,
		} = xScaleProvider(calculatedData);
    */


		return (
			<ChartCanvas height={window.innerHeight - 100}
					ratio={ratio}
					width={width}
					margin={{ left: 50, right: 50, top: 10, bottom: 30 }}
					type={type}
					seriesName="BTC/USDT"
					data={data}
					displayXAccessor={xAccessor}
					xAccessor={xAccessor}
					xScale={scaleTime()}
					xExtents={xExtents}>

				<Chart id={1} yExtents={d => [d.high, d.low]}>
					<XAxis axisAt="bottom" orient="bottom" ticks={10}/>
					<YAxis axisAt="left" orient="left" ticks={10} />
					<CandlestickSeries width={timeIntervalBarWidth(utcMinute)}/>
					<LineSeries yAccessor={sma20.accessor()} stroke={sma20.stroke()}/>
          <CurrentCoordinate yAccessor={sma20.accessor()} fill={sma20.stroke()} />
          <OHLCTooltip origin={[-40, 0]}/>

          <MouseCoordinateX
            at="top"
            orient="bottom"
            displayFormat={timeFormat("%Y-%m-%d %H:%M")} />
          <MouseCoordinateY
            at="right"
            orient="right"
            displayFormat={format(".0f")} />
				</Chart>
        <CrossHairCursor />
			</ChartCanvas>
		);
	}
}

CandleStickChart.propTypes = {
	data: PropTypes.array.isRequired,
	width: PropTypes.number.isRequired,
	ratio: PropTypes.number.isRequired,
	type: PropTypes.oneOf(["svg", "hybrid"]).isRequired,
};

CandleStickChart.defaultProps = {
	type: "svg",
};
CandleStickChart = fitWidth(CandleStickChart);

export default CandleStickChart;
