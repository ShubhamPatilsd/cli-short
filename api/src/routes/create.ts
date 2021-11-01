import { Router } from "express";

import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();

export const router = Router();
router.post("/", async (req, res) => {
  const { code, destinationURL } = req.body;

  const urlRegex =
    /[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)?/gi;
  const regex = new RegExp(urlRegex);

  if (!destinationURL.match(regex)) {
    res.status(400).send("Invalid URL");
  }

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
