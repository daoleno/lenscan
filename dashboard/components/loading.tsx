import { Loader2 } from "lucide-react";

export function Loading() {
  return (
    <div className="flex items-center space-x-2">
      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
      Loading...
    </div>
  );
}
