import * as z from "zod"

export const searchParamsSchema = z.object({
  page: z.string().default("1"),
  per_page: z.string().default("10"),
  sort: z.string().optional(),
  app: z.string().optional(),
  publication_type: z.string().optional(),
  is_momoka: z.string().optional(),
})
