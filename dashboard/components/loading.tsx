import { cn } from "@/lib/utils";
import { Loader2 } from "lucide-react";

interface LoadingProps {
  fixed?: boolean;
}

export function Loading({ fixed = true }: LoadingProps) {
  return (
    // <div className="flex items-center justify-center space-x-2 fixed top-0 left-0 right-0 bottom-0 bg-opacity-50 z-50">
    <div
      className={cn(
        "flex items-center justify-center space-x-2 top-0 left-0 right-0 bottom-0 bg-opacity-50 z-50",
        fixed ? "fixed" : "my-7"
      )}
    >
      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
      Loading...
    </div>
  );
}
