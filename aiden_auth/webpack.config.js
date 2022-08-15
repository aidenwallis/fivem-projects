const path = require("path");
const ESLintPlugin = require("eslint-webpack-plugin");

module.exports = {
  target: "node",
  entry: {
    client: path.join(__dirname, "./src/fivem-client/index.ts"),
    server: path.join(__dirname, "./src/fivem-server/index.ts"),
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: ["ts-loader"],
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: [".ts"],
  },
  plugins: [new ESLintPlugin()],
  output: {
    filename: "[name].js",
    path: path.join(__dirname, "dist"),
  },
  optimization: {
    minimize: true,
  },
};
