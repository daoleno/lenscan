import { Input } from "@/components/ui/input";
import { useRouter } from "next/router";
import { useState } from "react";

export function Search() {
  const [input, setInput] = useState("");
  const router = useRouter();
  const handleSearch = () => {
    if (input === "") return;
    router.push(`/profiles/${input}`);
  };

  return (
    <form onSubmit={handleSearch}>
      <Input
        type="search"
        placeholder="Search for a profile"
        className="h-9 md:w-[100px] lg:w-[300px]"
        onInput={(e) => setInput(e.currentTarget.value)}
      />
    </form>
  );
}
