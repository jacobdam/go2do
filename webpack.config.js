const path = require('path');
const webpack = require('webpack');

module.exports = {
  entry: ['./client/main.js'],
  output: {
    path: path.join(__dirname, 'static'),
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      {
        test: /\.js/,
        include: [path.join(__dirname, 'client')],
        loader: 'babel',
        query: {
          cacheDirectory: true
        }
      }
    ]
  }
}
