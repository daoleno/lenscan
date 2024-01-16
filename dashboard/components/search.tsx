"use client"

import { Input } from "@/components/ui/input"
import { useRouter } from "next/navigation"
import { useState } from "react"

export function Search() {
  const [input, setInput] = useState("")
  const [isLoading, setIsLoading] = useState(false) // New state variable for loading state
  const router = useRouter()

  const handleSearch = async (e: React.SyntheticEvent) => {
    e.preventDefault()
    if (input === "") return

    setIsLoading(true) // Set loading state to true

    try {
      // Simulate an asynchronous operation, e.g., fetching data
      await new Promise((resolve) => setTimeout(resolve, 2000))

      router.push(`/profile/${input}`)
    } catch (error) {
      console.error("Error occurred during search:", error)
    } finally {
      setIsLoading(false) // Set loading state to false after the operation completes
    }
  }

  return (
    <form onSubmit={handleSearch}>
      <Input
        type="search"
        placeholder="Search for a profile"
        className="h-9 md:w-[100px] lg:w-[300px]"
        onInput={(e) => setInput(e.currentTarget.value)}
      />
      {isLoading && <p>Loading...</p>} {/* Show loading state if isLoading is true */}
    </form>
  )
}
