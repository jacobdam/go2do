import React from 'react';
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux';
import List from 'material-ui/lib/lists/list';

import Todo from './Todo';
import * as todoActions from '../actions/todoActions';

function Todos({todos, toggleTodoComplete}) {
  return (
    <List>
      {todos.map(todo => (
        <Todo key={'todo-' + todo.id} todo={todo} onCompleteClick={()=>toggleTodoComplete(todo.id)}/>
      ))}
    </List>
  );
}

function mapStateToProps(state) {
  return { todos: state.todos };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(todoActions, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(Todos);
