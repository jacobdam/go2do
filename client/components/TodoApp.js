import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import AppBar from 'material-ui/lib/app-bar';
import FloatingActionButton from 'material-ui/lib/floating-action-button';
import ContentAdd from 'material-ui/lib/svg-icons/content/add';

import Todos from './Todos.js';
import styles from '../styles/styles.css';

import * as testActions from '../actions/testActions';

function TodoApp({fetchThing}) {
  debugger
  return (
    <div>
      <AppBar title="Todo App"/>
      <Todos/>
      <FloatingActionButton className={styles.addTodoButton}>
        <ContentAdd onClick={fetchThing}/>
      </FloatingActionButton>
    </div>
  );
}

function mapStateToProps(state) {
  return state => state;
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(testActions, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(TodoApp);
