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
      const req = new CandlesRequest();
      req.setExchange('binance')
      req.setMarket('BTC/USDT')
      req.setTimeframe('1m')
      const deadline = new Date();
      deadline.setSeconds(deadline.getSeconds() + 100)
      const stream = c.streamCandles(req, { deadline: deadline.getTime() })
      stream.on('data', (response) => {
        console.log('data', response);
        const candle = dexterToChartCandle(response.array)
        const data = this.state.data
        data.push(candle)
        this.setState({ data })
      });
      stream.on('status', function(status) {
        console.log(status.code);
        console.log(status.details);
        console.log(status.metadata);
      });
      stream.on('end', function(end) {
        // stream end signal
      });
      stream.on('error', function(err) {
        console.error(err)
      })
      console.log('stream', stream)
      window.stream = stream

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
