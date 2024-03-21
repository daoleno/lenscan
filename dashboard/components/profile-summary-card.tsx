import type { ProfileFragment } from "@lens-protocol/client";
import {
	ActivityIcon,
	CheckCircle2,
	Fingerprint,
	Tags,
	XCircle,
} from "lucide-react";
/* eslint-disable @next/next/no-img-element */
import Link from "next/link";
import Balance from "react-wrap-balancer";

import { formatCryptoValue, getIPFSURL } from "@/lib/utils";

import type { ProfileRevenue } from "@/app/api/analystics/revenue/getProfileRevenue";
import Image from "next/image";
import UserActivity from "./charts/user-activity";
import LennyCard from "./lenny-card";
import { LensAvatar } from "./lens-avatar";
import SocialCard from "./social-card";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Separator } from "./ui/separator";

export default function ProfileSummaryCard({
	profile,
	revenue,
}: {
	profile: ProfileFragment;
	revenue: ProfileRevenue[];
}) {
	const socialStats = ["followers", "following"];
	const activityStats = [
		"publications",
		"posts",
		"comments",
		"mirrors",
		"quotes",
		"countOpenActions",
	];
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
	];
	const classifiedStats = [
		{ title: "Social", keys: socialStats },
		{ title: "Activity", keys: activityStats },
		{ title: "Reactions", keys: reactionStats },
	];

	return (
		<>
			<div>
				<div className="mt-6 h-2/4 overflow-hidden rounded-lg sm:h-64">
					{profile.metadata?.coverPicture ? (
						<img
							className="w-full object-cover"
							src={getIPFSURL(profile.metadata.coverPicture) || ""}
							alt="cover"
						/>
					) : (
						<div className="h-full w-full bg-gradient-to-r from-rose-100 to-teal-100" />
					)}
				</div>
				<div className="flex flex-col items-center justify-between gap-7 sm:mx-32 sm:flex-row">
					<div className="-mt-16 flex w-full flex-col sm:w-2/5">
						<div className="mb-5 flex px-5">
							<LensAvatar
								profileId={profile.id}
								profilePicture={profile.metadata?.picture}
							/>
						</div>
						<h2 className="text-3xl font-bold text-foreground">
							{profile.metadata?.displayName}
						</h2>
						<span className="mt-2 text-muted-foreground">
							{profile.id} - #{Number(profile.id)}
						</span>
						<Balance className="mt-2 text-muted-foreground">
							@{profile.handle?.fullHandle}
						</Balance>
						<Balance className="mt-2 text-muted-foreground">
							{profile.metadata?.bio}
						</Balance>
						<div className="mt-2">
							<SocialCard
								profileHandle={profile.handle?.localName}
								profileId={profile.id}
								address={profile.handle?.ownedBy}
							/>
						</div>
					</div>
					<LennyCard profileId={profile.id} className="mt-3" />
				</div>
				<UserActivity className="my-7" profileId={profile.id} />
			</div>
			<div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-2">
				<Card>
					<CardHeader className="flex flex-row items-center justify-between space-y-0">
						<CardTitle>Stats</CardTitle>
						<ActivityIcon />
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
													<p className="font-bold">
														{(profile.stats as any)[key]}
													</p>
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
						<CardTitle>Revenue</CardTitle>
						<ActivityIcon />
					</CardHeader>
					<CardContent className="relative">
						{revenue.length === 0 ? (
							<div className="flex h-full w-full items-center justify-center min-h-64">
								<p className="text-center text-muted-foreground">No Revenue</p>
							</div>
						) : (
							<div className="flex flex-col gap-3">
								{revenue.map((revenue) => (
									<div
										className="flex items-center gap-1"
										key={revenue.currency}
									>
										<Image
											src={`/tokens/${revenue.currency_symbol.toLowerCase()}.svg`}
											alt={revenue.currency}
											width={20}
											height={20}
										/>
										<p className="font-bold">
											{formatCryptoValue(
												revenue.total_revenue,
												revenue.decimals,
											)}
										</p>
										<a
											className="text-xs text-muted-foreground hover:underline"
											href={`https://polygonscan.com/token/${revenue.currency}`}
											target="_blank"
											rel="noreferrer"
										>
											{revenue.currency_symbol}
										</a>
									</div>
								))}
							</div>
						)}
					</CardContent>
				</Card>

				<Card>
					<CardHeader className="flex flex-row items-center justify-between space-y-0">
						<CardTitle>Categories</CardTitle>
						<Tags />
					</CardHeader>

					<CardContent className="flex flex-col space-y-3 overflow-hidden">
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
								),
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
	);
}
