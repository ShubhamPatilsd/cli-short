import { Router } from "express";

export const router = Router();
router.get("/", (req, res) => {
  //req.body stuff
  //do stuff in the database
  res.send("Created");
});
