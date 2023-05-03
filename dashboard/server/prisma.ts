import { Prisma, PrismaClient } from "@prisma/client";

declare global {
  var prisma:
    | PrismaClient<
        Prisma.PrismaClientOptions,
        "query" | "info" | "warn" | "error"
      >
    | undefined;
}

const prisma =
  global.prisma ||
  new PrismaClient({
    log: [
      {
        emit: "event",
        level: "query",
      },
      {
        emit: "stdout",
        level: "error",
      },
      {
        emit: "stdout",
        level: "info",
      },
      {
        emit: "stdout",
        level: "warn",
      },
    ],
  });

if (process.env.NODE_ENV === "development") global.prisma = prisma;

export default prisma;

prisma.$on("query", (e) => {
  const color = e.duration > 5000 ? "\x1b[31m%s\x1b[0m" : "%s";
  console.log(color, "Query: " + e.query);
  console.log(color, "Params: " + e.params);
  console.log(color, "Duration: " + e.duration + "ms");
});
