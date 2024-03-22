"use client"

import { useRouter } from "next/navigation"
import { useState } from "react"

import { Input } from "@/components/ui/input"
import { extractPublicationId } from "@/lib/utils"

export function Search() {
	const [input, setInput] = useState("")
	const router = useRouter()
	const handleSearch = (e: React.SyntheticEvent) => {
		e.preventDefault()
		if (input === "") return

		const pubId = extractPublicationId(input)
		if (pubId) {
			router.push(`/publication/${pubId}`)
			return
		}

		// if it's a number start with #, try to convert it to hex and navigate to profile
		if (input.startsWith("#")) {
			const hexNumber = Number.parseInt(input.slice(1), 10)
			let hex = hexNumber.toString(16)
			// Ensure hex value has even length by padding with a leading zero if necessary
			hex = hex.length % 2 === 0 ? hex : `0${hex}`
			router.push(`/profile/0x${hex}`)
			return
		}

		router.push(`/profile/${input}`)
	}

	return (
		<form onSubmit={handleSearch}>
			<Input
				type="search"
				placeholder="Search for a profile or publication"
				className="h-9 md:w-[100px] lg:w-[300px]"
				onInput={(e) => setInput(e.currentTarget.value)}
			/>
		</form>
	)
}
