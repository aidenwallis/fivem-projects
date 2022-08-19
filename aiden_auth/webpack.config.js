const path = require("path");

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
  plugins: [],
  output: {
    filename: "[name].js",
    path: path.join(__dirname, "dist"),
  },
  optimization: {
    minimize: true,
  },
};
