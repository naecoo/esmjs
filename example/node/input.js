const express = require("express");

const app = express();

app.use("/", (req, res) => {
  console.log("new context")
  console.log("method: ", req.method);
  console.log("path: ", req.path);
  console.log();
  res.end("Hello World!");
});

app.listen(8080);
console.log("app is running in http://localhost:8080")
