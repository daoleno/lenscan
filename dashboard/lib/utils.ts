import {
	NftImageFragment,
	ProfilePictureSetFragment,
} from "@lens-protocol/client"
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import * as timeago from "timeago.js"
import { formatUnits } from "viem"

export function shortHash(hash: string | undefined, length = 10) {
	if (!hash) {
		return "-"
	}
	return `${hash.slice(0, length)}...`
}

export function age(ts: number | undefined | null) {
	if (!ts) {
		return "-"
	}
	return timeago.format(ts * 1000)
}

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs))
}

export function formatNumber(n: number | string) {
	const formatter = new Intl.NumberFormat("en-US", {
		maximumFractionDigits: 2,
	})
	return formatter.format(Number(n))
}

// format any POST_CREATED style to PostCreated style
export function formatEventType(type: string | null) {
	if (!type) {
		return "-"
	}
	return type
		.split("_")
		.map((w) => w[0] + w.slice(1).toLowerCase())
		.join("")
}

const ipfsGateway = "https://lens.infura-ipfs.io"

export function processIPFSURL(uri: string | null | undefined) {
	if (uri && uri.startsWith("ipfs://")) {
		const cid = uri.replace("ipfs://", "")
		return `${ipfsGateway}/ipfs/${cid}`
	}
	return uri
}

export function getIPFSURL(
	picture: ProfilePictureSetFragment | NftImageFragment | string | null,
) {
	if (typeof picture === "string") {
		return processIPFSURL(picture)
	}

	let uri
	if (!picture) {
		return uri
	}

	if (picture.__typename === "ImageSet") {
		uri = picture?.optimized?.uri
	}
	if (picture.__typename === "NftImage") {
		uri = picture.image.optimized?.uri
	}

	return processIPFSURL(uri)
}

export function extractPublicationId(input: string) {
	// Define the regex pattern to match the ID formats within different input types
	const pattern = /0x[0-9A-Fa-f]+-0x[0-9A-Fa-f]+(-[0-9A-Fa-fDA]+-[0-9A-Fa-f]+)?/

	// Use the pattern to search for a match in the input
	const match = input.match(pattern)

	// If a match is found, return the matched publicationId, else return null
	return match ? match[0] : null
}

export function formatCryptoValue(value: bigint, decimals: number) {
	// Assuming formatUnits converts the value based on the provided decimals
	const number = parseFloat(formatUnits(value, decimals))

	// Format the number with thousands separator and fixed decimal places
	// Adjust minimumFractionDigits and maximumFractionDigits as needed
	return new Intl.NumberFormat("en-US", {
		style: "decimal",
		minimumFractionDigits: 0, // Minimum number of decimal places to show
		maximumFractionDigits: 6, // Maximum number of decimal places to show
	}).format(number)
}
