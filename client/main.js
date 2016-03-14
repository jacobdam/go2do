import 'babel-polyfill';

import React from 'react';
import ReactDOM from 'react-dom';
import { createStore, combineReducers, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import promiseMiddleware from 'redux-promise';

import TodoApp from './components/TodoApp';
import todos from './reducers/todos'
import testReducer from './reducers/testReducer'


let reducer = combineReducers({ todos, test: testReducer });
let store = createStore(reducer, {}, applyMiddleware(promiseMiddleware));

store.dispatch({ type: 'ADD_TODO', payload: { todo: { id: 1, title: 'Hello 1' } } });
store.dispatch({ type: 'ADD_TODO', payload: { todo: { id: 2, title: 'Hello 2' } } });

ReactDOM.render(
  <Provider store={store}>
    <TodoApp/>
  </Provider>,
  document.getElementById('root')
)
