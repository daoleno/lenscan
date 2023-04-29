import { XOctagon } from "lucide-react";

export function Error({ msg }: { msg: string }) {
  return (
    <div className="flex items-center justify-center space-x-2 fixed top-0 left-0 right-0 bottom-0 bg-opacity-50 text-red-500 text-xl">
      <XOctagon className="mr-2 h-5 w-5" />
      {msg}
    </div>
  );
}
