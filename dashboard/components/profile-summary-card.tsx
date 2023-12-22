/* eslint-disable @next/next/no-img-element */

import Image from "next/image"
import Link from "next/link"
import { ProfileFragment } from "@lens-protocol/client"
import { CheckCircle2, Fingerprint, Tags, XCircle } from "lucide-react"
import Balance from "react-wrap-balancer"

import { getIPFSURL } from "@/lib/utils"

import { Card, CardContent, CardHeader, CardTitle } from "./ui/card"
import { Separator } from "./ui/separator"

export default function ProfileSummaryCard({
  profile,
}: {
  profile: ProfileFragment
}) {
  const socialStats = ["followers", "following"]
  const activityStats = [
    "comments",
    "posts",
    "mirrors",
    "quotes",
    "publications",
    "countOpenActions",
  ]
  const reactionStats = [
    "upvotes",
    "downvotes",
    "upvoted",
    "downvoted",
    "collects",
    // "upvoteReactions",
    // "downvoteReactions",
    // "upvoteReacted",
    // "downvoteReacted",
  ]
  const classifiedStats = [
    { title: "Social", keys: socialStats },
    { title: "Activity", keys: activityStats },
    { title: "Reactions", keys: reactionStats },
  ]

  return (
    <>
      <div>
        <div className="mt-6 h-2/4 overflow-hidden rounded-lg sm:h-64">
          {profile.metadata?.coverPicture ? (
            <img
              className="w-full object-cover"
              src={getIPFSURL(profile.metadata.coverPicture)}
              alt="cover"
            />
          ) : (
            <div className="h-full w-full bg-gradient-to-r from-rose-100 to-teal-100" />
          )}
        </div>
        <div className="flex flex-col sm:flex-row">
          <div className="-mt-16 flex flex-col sm:w-1/2 ">
            <div className="mb-5 flex px-5">
              {profile.metadata?.picture ? (
                <img
                  className="h-32 w-32 rounded-full object-cover"
                  src={getIPFSURL(profile.metadata.picture)}
                  alt="profile"
                />
              ) : (
                <Image
                  src="/images/default-profile.png"
                  alt="profile"
                  width={128}
                  height={128}
                  className="h-32 w-32 rounded-full object-cover"
                />
              )}
            </div>
            <div className="mb-8 flex flex-col px-7">
              <h2 className="text-3xl font-bold text-foreground">
                {profile.metadata?.displayName}
              </h2>
              <span className="mt-2 text-muted-foreground">
                {profile.id} - #{Number(profile.id)}
              </span>
              <Balance className="mt-2 text-muted-foreground">
                @{profile.handle?.fullHandle}
              </Balance>
              <p className="mt-2 text-muted-foreground">
                {profile.metadata?.bio}
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Stats</CardTitle>
          </CardHeader>
          <CardContent className="flex items-center justify-between sm:flex-row sm:space-x-6">
            <div className="flex flex-col">
              {classifiedStats.map(({ title, keys }) => (
                <div key={title}>
                  <Separator className="my-4" />
                  <div className="flex flex-col gap-1" key={title}>
                    <h2 className="text-sm font-semibold uppercase text-muted-foreground">
                      {title}
                    </h2>
                    <div className="grid grid-cols-3 gap-1">
                      {keys.map((key) => (
                        <div className="flex items-center gap-1" key={key}>
                          <p className="font-bold">{profile.stats[key]}</p>
                          <p className="text-xs capitalize text-muted-foreground">
                            {key}
                          </p>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Categories</CardTitle>
            <Tags />
          </CardHeader>

          <CardContent className="flex flex-col space-y-3 overflow-scroll">
            {profile.metadata?.attributes?.map(
              ({ key, value }) =>
                value !== "[]" && (
                  <div
                    key={key}
                    className="flex flex-col space-y-1 font-medium"
                  >
                    <span className="text-sm uppercase text-muted-foreground">
                      {key.replace(/([a-z])([A-Z])/g, "$1 $2")}
                    </span>
                    {key.toLocaleLowerCase().includes("website") ? (
                      <Link href={value}>{value} </Link>
                    ) : key.toLocaleLowerCase().includes("twitter") ? (
                      <Link href={`https://twitter.com/${value}`}>{value}</Link>
                    ) : (
                      <span>{value}</span>
                    )}
                  </div>
                )
            )}
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>On-Chain Identity</CardTitle>
            <Fingerprint />
          </CardHeader>
          <CardContent className="flex flex-col space-y-3">
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium uppercase text-muted-foreground">
                Proof of Humanity
              </span>
              {profile.onchainIdentity.proofOfHumanity ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium uppercase text-muted-foreground">
                ENS Name
              </span>
              <span className="font-medium">
                {profile.onchainIdentity.ens?.name
                  ? String(profile.onchainIdentity.ens.name)
                  : "-"}
              </span>
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium uppercase text-muted-foreground">
                Sybil.org Verified
              </span>
              {profile.onchainIdentity.sybilDotOrg.verified ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium uppercase text-muted-foreground">
                Twitter Handle
              </span>
              <span className="font-medium">
                {profile.onchainIdentity.sybilDotOrg.source?.twitter.handle ||
                  "-"}
              </span>
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  )
}
