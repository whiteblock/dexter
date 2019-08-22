import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import * as grpcWeb from 'grpc-web';
//import {DataServiceClient} from './DexterData';
const {Candle, CandlesRequest, CandlesResponse, ExchangesRequest, ExchangesResponse, MarketsRequest, MarketsResponse, DataClient, DataPromiseClient} = require('./data_grpc_web_pb');

window.Candle = Candle
window.CandlesRequest = CandlesRequest
window.CandlesResponse = CandlesResponse
window.ExchangesRequest = ExchangesRequest
window.ExchangesResponse = ExchangesResponse
window.MarketsRequest = MarketsRequest
window.MarketsResponse = MarketsResponse
window.DataClient = DataClient
window.DataPromiseClient = DataPromiseClient

window.c = new DataPromiseClient('http://34.94.208.167:8080')

ReactDOM.render(<App dexterDataClient={window.c} />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

