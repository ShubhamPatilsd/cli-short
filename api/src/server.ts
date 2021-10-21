import express from "express";
import { apiRoutes } from "./router";

import dotenv from "dotenv";
dotenv.config();

const app = express();
app.use(express.json());

app.use("/", apiRoutes);

app.listen(process.env.PORT || 5000, () => {
  console.log("Listening on port");
});
