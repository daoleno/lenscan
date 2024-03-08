import Image from "next/image"
import Link from "next/link"

import { cn } from "@/lib/utils"

export default function SocialLink({
	publicationId,
}: {
	publicationId: string | undefined
}) {
	const links = [
		{
			appName: "Hey",
			url: `https://hey.xyz/posts/${publicationId}`,
			icon: "/apps/hey.png",
		},
		{
			appName: "Orb",
			url: `https://orb.ac/p/${publicationId}`,
			icon: "/apps/orb.png",
		},
		{
			appName: "Buttrfly",
			url: `https://buttrfly.app/post/${publicationId}`,
			icon: "/apps/buttrfly.png",
		},
		{
			appName: "Tape",
			url: `https://tape.xyz/watch/${publicationId}`,
			icon: "/apps/tape.png",
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
							"hover:opacity-80",
						)}
					/>
					{/* <span>{appName}</span> */}
				</Link>
			))}
		</div>
	)
}
