import React from 'react';
import { render } from 'react-dom';
import CandleStickChart from './CandleStickChart';
import { getData, dexterToChartCandle } from "./utils"

import { TypeChooser } from "react-stockcharts/lib/helper";

const {Candle, CandlesRequest, CandlesResponse, ExchangesRequest, ExchangesResponse, MarketsRequest, MarketsResponse, DataClient, DataPromiseClient} = require('./data_grpc_web_pb');
class ChartComponent extends React.Component {
	componentDidMount() {
    const c = this.props.dexterDataClient
    const req = new CandlesRequest();
    req.setExchange('binance')
    req.setMarket('BTC/USDT')
    req.setTimeframe('1m')
    c.getCandles(req).then((candles) => {
      window.candles = candles
      console.log(typeof(candles.array), candles.array)
      const data = candles.array[0].map(dexterToChartCandle)
      console.log('data', data)
      this.setState({ data })
    }).then(() => {
      console.log('try to start a stream next')
    })

    /*
		getData().then(data => {
      console.log(data)
			this.setState({ data })
		})
    */
	}
	render() {
		if (this.state == null) {
			return <div>Loading...</div>
		}
		return (
			<TypeChooser>
				{type => <CandleStickChart type={type} data={this.state.data} />}
			</TypeChooser>
		)
	}
}

export default ChartComponent
