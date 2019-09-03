import React from 'react';
import logo from './logo.svg';
import './App.css';
import ChartComponent from './ChartComponent';

function App(props) {
  return (
    <div className="App">
      <ChartComponent dexterDataClient={props.dexterDataClient} exchange={'binance'} market={'BTC/USDT'} timeframe={'3m'} />
    </div>
  );
}

export default App;
