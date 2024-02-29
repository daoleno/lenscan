import Image from "next/image"
import Link from "next/link"
import { Hex, hexToNumber } from "viem"

import { cn } from "@/lib/utils"

export default function SocialCard({
  profileHandle,
  address,
  profileId,
}: {
  profileHandle: string | undefined
  address: string | undefined
  profileId: string | undefined
}) {
  const links = [
    {
      appName: "Hey",
      url: `https://hey.xyz/u/${profileHandle}`,
      icon: "/apps/hey.png",
    },
    {
      appName: "Orb",
      url: `https://orb.ac/@${profileHandle}`,
      icon: "/apps/orb.png",
    },
    {
      appName: "Buttrfly",
      url: `https://buttrfly.app/profile/${profileHandle}`,
      icon: "/apps/buttrfly.png",
    },
    {
      appName: "Tape",
      url: `https://tape.xyz/u/${profileHandle}`,
      icon: "/apps/tape.png",
    },
    {
      appName: "Polygon",
      url: `https://polygonscan.com/address/${address}`,
      icon: "/apps/polygon.svg",
    },
    {
      appName: "OpenSea",
      url: `https://opensea.io/assets/matic/0xdb46d1dc155634fbc732f92e853b10b288ad5a1d/${hexToNumber(
        profileId as Hex
      )}`,
      icon: "/apps/opensea.svg",
    },
  ]

  return (
    <div className="flex gap-1">
      {links.map(({ appName, url, icon }) => (
        <Link
          key={appName}
          href={url}
          className="flex items-center gap-2 text-sm font-medium hover:underline"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            src={icon}
            alt={appName}
            width={28}
            height={28}
            className={cn(
              appName === "Polygon" || appName === "OpenSea"
                ? "ml-1.5 h-5 w-5"
                : "",
              "hover:opacity-80"
            )}
          />
          {/* <span>{appName}</span> */}
        </Link>
      ))}
    </div>
  )
}
