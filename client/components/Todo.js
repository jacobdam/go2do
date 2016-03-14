import React from 'react';
import ListItem from 'material-ui/lib/lists/list-item';
import Checkbox from 'material-ui/lib/checkbox';

function Todo({todo, onCompleteClick}) {
  return (
    <ListItem
      primaryText={todo.title}
      leftCheckbox={<Checkbox checked={todo.isCompleted} onCheck={onCompleteClick}/>}
    />
  );
}

export default Todo;
