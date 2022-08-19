const path = require("path");
const HTMLWebpackPlugin = require("html-webpack-plugin");
const webpack = require("webpack");
const appConfig = require("./config.json");

module.exports = {
  entry: path.join(__dirname, "./src/client/index.tsx"),
  output: {
    filename: "[name].js",
    path: path.join(__dirname, "dist/client"),
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: ["ts-loader"],
      },
      {
        test: /\.s[ca]ss$/,
        use: ["style-loader", "css-loader", "sass-loader"],
      },
      {
        test: /\.(jpe?g|png|gif|svg)$/i,
        use: ["file-loader"],
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"],
  },
  plugins: [
    new HTMLWebpackPlugin({ template: path.join(__dirname, "./src/client/index.html") }),
    new webpack.DefinePlugin({
      "process.env": JSON.stringify({
        NODE_ENV: process.env.NODE_ENV || "production",
        TRPC_HOST: appConfig.trpc.host,
      }),
    }),
  ],
  devServer: {
    compress: true,
    port: 3403,
  },
};
