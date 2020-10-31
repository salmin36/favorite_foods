/**
 * http://usejsdoc.org/
 */
const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  watch: false,
  mode: 'development',
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    port: 8080,
    host: "0.0.0.0"
  },
  entry: './src/index.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist')
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loaders: ["react-hot-loader/webpack", "babel-loader"],
      },
      {
        test:/\.css$/,
        use:['style-loader','css-loader']
      },
      {
        test:/\.(jpe?g|png|gif|svg)$/i,
        use:["url-loader?limit=10000","img-loader"]
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: 'src/index.html'
    })
  ]
}