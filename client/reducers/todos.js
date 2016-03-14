import { TOGGLE_TODO_COMPLETE } from '../actionTypes';

function todoReducer(state = {}, action) {
  switch (action.type) {
    case 'SET_TODO_COMPLETE':
      return Object.assign({},
        state,
        { isCompleted: action.payload.isCompleted }
      );

    case TOGGLE_TODO_COMPLETE:
      return Object.assign({},
        state,
        { isCompleted: !state.isCompleted }
      );


    case 'ADD_TODO':
      return Object.assign({}, action.payload.todo);
  }
}

function todosReducer(state = [], action) {
  switch (action.type) {
    case 'SET_TODO_COMPLETE':
    case TOGGLE_TODO_COMPLETE:
      let todoIndex = state.findIndex(todo => (todo.id == action.payload.id));
      return [].concat(
        state.slice(0, todoIndex),
        [todoReducer(state[todoIndex], action)],
        state.slice(todoIndex + 1)
      );

    case 'ADD_TODO':
      return [].concat(
        state,
        todoReducer(null, action)
      );
  }

  return state;
}

export default todosReducer;
