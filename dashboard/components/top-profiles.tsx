import Link from "next/link"

import { getIPFSURL } from "@/lib/utils"
import { getTopProfiles } from "@/app/api/analystics/getTopProfiles"

import { Card, CardContent, CardHeader, CardTitle } from "./ui/card"
import { ScrollArea, ScrollBar } from "./ui/scroll-area"

export default async function TopProfiles() {
  const topProfiles = await getTopProfiles("1M")

  return (
    <Card>
      <CardHeader>
        <CardTitle>Daily Stars</CardTitle>
      </CardHeader>
      <ScrollArea className="w-full whitespace-nowrap">
        <CardContent className="flex w-max space-x-4 p-4">
          {topProfiles.map((profile) => (
            <Link
              className="shrink-0 rounded-lg border"
              key={profile.id}
              href={`/profile/${profile.id}`}
            >
              <div
                className="relative m-2 h-48 rounded-md bg-cover bg-center"
                style={{
                  backgroundImage: `url(${
                    profile.metadata?.picture
                      ? getIPFSURL(profile.metadata.picture)
                      : null
                  })`,
                }}
              >
                <span className="absolute bottom-0 m-2 inline-block rounded-full bg-muted px-2 py-1 text-sm font-bold text-muted-foreground">
                  {profile.metadata?.displayName}
                </span>
              </div>

              <div className="px-6 py-4">
                <div className="flex space-x-4 text-center text-base ">
                  <div className="flex-1">
                    <p className="text-lg font-bold">
                      {profile.stats.followers}
                    </p>
                    <p>Followers</p>
                  </div>
                  <div className="flex-1">
                    <p className="text-lg font-bold">
                      {profile.stats.following}
                    </p>
                    <p>Following</p>
                  </div>
                  {/* <div className="flex-1">
                <p className="text-lg font-bold text-red-500">{11}</p>
                <p>Score</p>
              </div> */}
                </div>
              </div>
            </Link>
          ))}
        </CardContent>
        <ScrollBar orientation="horizontal" />
      </ScrollArea>
    </Card>
  )
}
