import Link from "next/link"

import { siteConfig } from "@/config/site"

export default function NotFound({ type }: { type: string }) {
  return (
    <main className="grid min-h-full place-items-center px-6 py-24 sm:py-32 lg:px-8">
      <div className="text-center">
        <p className="text-base font-semibold text-muted-foreground">404</p>
        <h1 className="mt-4 text-3xl font-bold tracking-tight sm:text-5xl">
          {type} not found
        </h1>
        <p className="mt-6 text-base leading-7 text-muted-foreground">
          Sorry, we couldn’t find the {type.toLowerCase()} you’re looking for.
        </p>
        <div className="mt-10 flex items-center justify-center gap-x-6">
          <Link
            href="/"
            className="border-1 rounded-md border px-3.5 py-2.5 text-sm font-semibold shadow-sm hover:shadow-md focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600"
          >
            Go back home
          </Link>
          <Link
            href={siteConfig.links.telegram}
            target="_blank"
            className="text-sm font-semibold"
          >
            Contact support <span aria-hidden="true">&rarr;</span>
          </Link>
        </div>
      </div>
    </main>
  )
}
