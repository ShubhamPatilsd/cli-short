-- CreateTable
CREATE TABLE "Link" (
    "code" TEXT NOT NULL,
    "destination" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Link_code_key" ON "Link"("code");
