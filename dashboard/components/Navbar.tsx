import { MagnifyingGlassIcon } from "@heroicons/react/20/solid";
import Link from "next/link";

function classNames(...classes: any) {
  return classes.filter(Boolean).join(" ");
}

export default function NavBar() {
  return (
    <div className="flex h-16 justify-between">
      <div className="flex px-2 lg:px-0">
        <div className="flex flex-shrink-0 items-center text-md font-bold text-gray-900">
          🎄 Lenscan
        </div>
        <div className="hidden lg:ml-6 lg:flex lg:space-x-8">
          {/* Current: "border-indigo-500 text-gray-900", Default: "border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700" */}
          <Link
            href="/"
            className="inline-flex items-center border-b-2 border-indigo-500 px-1 pt-1 text-sm font-medium text-gray-900"
          >
            Dashboard
          </Link>
          <Link
            href="/events"
            className="inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700"
          >
            Events
          </Link>
          <Link
            href="/profiles"
            className="inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700"
          >
            Profiles
          </Link>
          <Link
            href="/publications"
            className="inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium text-gray-500 hover:border-gray-300 hover:text-gray-700"
          >
            Publications
          </Link>
        </div>
      </div>
      <div className="flex flex-1 items-center justify-center px-2 lg:ml-6 lg:justify-end">
        <div className="w-full max-w-lg lg:max-w-xs">
          <label htmlFor="search" className="sr-only">
            Search
          </label>
          <div className="relative">
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              <MagnifyingGlassIcon
                className="h-5 w-5 text-gray-400"
                aria-hidden="true"
              />
            </div>
            <input
              id="search"
              name="search"
              className="block w-full rounded-md border-0 bg-white py-1.5 pl-10 pr-3 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              placeholder="Search"
              type="search"
            />
          </div>
        </div>
      </div>
    </div>
  );
}
