"use client"

import { useRouter } from "next/navigation"
import { useState } from "react"

import { Input } from "@/components/ui/input"
import { isPublicationId } from "@/lib/utils"

export function Search() {
	const [input, setInput] = useState("")
	const router = useRouter()
	const handleSearch = (e: React.SyntheticEvent) => {
		e.preventDefault()
		if (input === "") return

		if (isPublicationId(input)) {
			router.push(`/publication/${input}`)
			return
		}

		router.push(`/profile/${input}`)
	}

	return (
		<form onSubmit={handleSearch}>
			<Input
				type="search"
				placeholder="Search for a profile"
				className="h-9 md:w-[100px] lg:w-[300px]"
				onInput={(e) => setInput(e.currentTarget.value)}
			/>
		</form>
	)
}
