import Image from "next/image"

import { getLenny } from "@/lib/lenny"
import { getIPFSURL } from "@/lib/utils"

export interface LensAvatarProps {
	profileId: string
	profilePicture?: any
}
export async function LensAvatar({
	profileId,
	profilePicture,
}: LensAvatarProps) {
	const lenny = await getLenny(profileId)

	return (
		<div>
			{!profilePicture ? (
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
					src={getIPFSURL(profilePicture) || ""}
					alt="profile"
				/>
			)}
		</div>
	)
}
