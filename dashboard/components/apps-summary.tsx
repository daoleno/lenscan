import Image from "next/image"
import Link from "next/link"

import { recommendedApps } from "@/config/apps"
import { getTotalApps } from "@/app/api/analystics/topapps/getAppStats"

import { Button } from "./ui/button"
import { Card, CardContent } from "./ui/card"

export default async function AppsSummary() {
  const totalApps = await getTotalApps()
  return (
    <div className="flex flex-col gap-5">
      <div className="flex flex-col gap-3">
        <h2 className="text-3xl font-bold tracking-tight">Discover Apps</h2>
        <span className="tracking-tight text-muted-foreground">
          There are a total of
          <span className="ml-1 text-xl font-bold">{totalApps}</span> apps on
          Lens. Check out some of the most popular apps below.
        </span>
      </div>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
        {recommendedApps.map((app) => (
          <Card key={app.name} className="w-full">
            <CardContent className="flex flex-col p-4 sm:flex-row sm:items-center">
              <div className="flex items-center">
                <Image
                  alt={`${app.name} logo`}
                  className="h-12 w-12 rounded-md object-cover sm:h-15 sm:w-15"
                  height="60"
                  width="60"
                  src={app.icon}
                />
                <div className="ml-4 flex flex-col">
                  <span className="text-base font-bold sm:text-lg">{app.name}</span>
                  <span
                    className="mt-1 text-sm text-muted-foreground"
                    style={{
                      whiteSpace: "pre-line",
                    }}
                  >
                    {app.description}
                  </span>
                </div>
              </div>
              <Link
                href={app.url}
                passHref
                className="mt-4 sm:mt-0 sm:ml-auto"
                target="_blank"
              >
                <Button className="w-full sm:w-auto">VISIT</Button>
              </Link>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  )
}
