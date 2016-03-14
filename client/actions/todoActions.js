import {
  TOGGLE_TODO_COMPLETE,
  ADD_TODO,
  REMOVE_TODO,
} from '../actionTypes';

export function toggleTodoComplete(id) {
  return {
    type: TOGGLE_TODO_COMPLETE,
    payload: {
      id
    }
  }
}

export function setTodoComplete(id, isCompleted) {
  return {
    type: 'SET_TODO_COMPLETE',
    payload: {
      id,
      isCompleted
    }
  }
}

export function addTodo(todo) {
  return {
    type: 'ADD_TODO',
    payload: {
      todo
    }
  }
}

export function removeTodo(id) {
  return {
    type: 'REMOVE_TODO',
    payload: {
      id
    }
  }
}
