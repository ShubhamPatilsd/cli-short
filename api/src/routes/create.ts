import { Router } from "express";

import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();

export const router = Router();
router.post("/", async (req, res) => {
  //req.body stuff
  const { code, destinationURL } = req.body;
  //do stuff in the database

  try {
    const result = await prisma.link.create({
      data: {
        code: code,
        destination: destinationURL,
      },
    });

    res.json(result);
  } catch (err) {
    res.status(409).send("Shortened link with given code already exists");
  }
});
