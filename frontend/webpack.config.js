module.exports = {
    mode: "development",
    entry: "./src/main.tsx",
    output: {
        path: `${__dirname}/dist`,
        filename: "main.js"
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                // babel経由でコンパイル
                use: "babel-loader"
            }
        ]
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".json"]
    },
    target: ["web"],
    // webpack-dev-server
    devServer: {
        hot: true,
        contentBase: `${__dirname}`,
        port: 5000
    }
}