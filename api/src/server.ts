import express from "express";
import { apiRoutes } from "./router";

const app = express();

app.use("/", apiRoutes);

app.listen(process.env.PORT || 5000, () => {
  console.log("Listening on port");
});
