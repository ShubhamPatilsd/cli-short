import { Router } from "express";
import { router as CreateRoutes } from "./routes/create";
import { router as VisitRoutes } from "./routes/visit";

export const apiRoutes = Router();

apiRoutes.use("/create", CreateRoutes);
apiRoutes.use("/", VisitRoutes);
