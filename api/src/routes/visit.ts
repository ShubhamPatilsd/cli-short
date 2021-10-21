import { Router } from "express";

import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();

export const router = Router();
router.get("/:code", async (req, res) => {
  const result = await prisma.link.findUnique({
    where: {
      code: req.params.code,
    },
  });

  if (!result) {
    res.status(404).send("That link doesn't exist");
  } else {
    res.redirect(result.destination);
  }
});
