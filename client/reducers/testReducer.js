function test(state = 0, action) {
  console.log('test reducer', state, action.type)
  return state;
}

export default test;
