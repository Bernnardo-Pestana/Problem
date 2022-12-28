import './App.css';

import Users from './components/users/user';
import Products from './components/products/products';
import Purchases from './components/purchases/purchase';
import React, {useState,Component } from "react";






function App() {

  let [select,Setselect] = useState('')

  function changeView(val){
    Setselect(val)
  }

  return (
    <div className="App">
        <button onClick={(e) => changeView(1)}>Users</button>
        <button  onClick={(e) => changeView(2)}>Products</button>
        <button onClick={(e) => changeView(3)}>Purchases</button>
        {
         (function(){
          if( select === 1){
            return <div> <Users /></div>
          }else if( select === 2){
            return <div> <Products /></div>
          }else if( select === 3){
            return <div>  <Purchases /></div>
          }
         })
        ()}
    </div>
  );

  //<Users />
}


export default App;
