import Image from "next/image"
import { ProfileFragment } from "@lens-protocol/client"

import { getLenny } from "@/lib/lenny"
import { getIPFSURL } from "@/lib/utils"

export async function LensAvatar({ profile }: { profile: ProfileFragment }) {
  const lenny = await getLenny(profile.id)

  return (
    <div>
      {!profile.metadata?.picture ? (
        <Image
          src={lenny.image}
          alt="Lenny Avatar"
          width={200}
          height={200}
          className="h-32 w-32 rounded-full object-cover"
        />
      ) : (
        <img
          className="h-32 w-32 rounded-full object-cover"
          src={getIPFSURL(profile.metadata.picture)}
          alt="profile"
        />
      )}
    </div>
  )
}
