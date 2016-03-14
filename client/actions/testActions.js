import { createAction } from 'redux-actions';

function fetchThingImpl() {
  return new Promise((resolve, reject) => {
    console.log('xxxx')
    setTimeout(function() {
      console.log('xxxx2')
      resolve(1234)
    }, 3000);
  });
}

export var fetchThing = createAction('FETCH_THING', (id) => {
  console.log('FETCH_THING action')
  return fetchThingImpl();
});
