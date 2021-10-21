import { Router } from "express";
import { router as CreateRoutes } from "./routes/create";

export const apiRoutes = Router();

apiRoutes.use("/create", CreateRoutes);
