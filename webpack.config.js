const path = require('path');
const webpack = require('webpack');
const postcssImport = require('postcss-import')

module.exports = {
  entry: ['./client/main.js'],
  output: {
    path: path.join(__dirname, 'static'),
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        include: [path.join(__dirname, 'client')],
        loader: 'babel',
        query: {
          cacheDirectory: true
        }
      },
      {
        test: /\.css$/,
        include: [path.join(__dirname, 'client')],
        loader: "style-loader!css-loader?modules&importLoaders=1!postcss-loader"
      }
    ]
  },
  postcss: function (webpack) {
    return [
      postcssImport({
        addDependencyTo: webpack
      })
    ];
  }
}
