import Image from "next/image"
import Link from "next/link"

import { getTopApps } from "@/app/api/analystics/topapps/getAppStats"

import { Button } from "./ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card"

export default async function DailyApps() {
  const dailyApps = await (await getTopApps("1D")).splice(0, 4)

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <div>Daily Apps</div>
          <Link href="/apps" passHref>
            <Button className="font-semibold" variant={"link"}>
              VIEW ALL
            </Button>
          </Link>
        </CardTitle>
      </CardHeader>
      <div className="w-full whitespace-nowrap">
        <CardContent className="grid grid-cols-2 gap-7">
          {dailyApps.map((app) => (
            <Card key={app.name}>
              <CardContent className="flex items-center p-4 sm:flex-row">
                <Image
                  alt="Orb logo"
                  className="h-15 w-15 aspect-[60/60] rounded-md object-cover"
                  height="60"
                  width="60"
                  src={app.icon || "/images/default-profile.png"}
                />
                <div className="ml-4 mr-8 flex flex-grow flex-col sm:gap-0">
                  <span className="text-lg font-bold">{app.name}</span>
                  <span
                    className="text-sm"
                    style={{
                      whiteSpace: "pre-line",
                    }}
                  >
                    {app.description}
                  </span>
                </div>
                <Link
                  href={app.url || "#"}
                  passHref
                  className="flex flex-grow justify-end"
                  target="_blank"
                >
                  <Button className="hidden md:inline-block">VISIT</Button>
                </Link>
              </CardContent>
            </Card>
          ))}
        </CardContent>
      </div>
    </Card>
  )
}
