import Link from "next/link";

import { MainNav } from "@/components/main-nav";
import { buttonVariants } from "@/components/ui/button";
import { siteConfig } from "@/config/site";
import { cn } from "@/lib/utils";
import Image from "next/image";
import { Search } from "./search";

export function SiteHeader() {
  return (
    <header className="supports-backdrop-blur:bg-background/60 sticky top-0 z-40 w-full border-b bg-background/95 shadow-sm backdrop-blur">
      <div className="container flex h-14 items-center">
        <MainNav />

        <div className="flex flex-1 items-center justify-between space-x-2 sm:space-x-4 md:justify-end">
          <div className="ml-auto flex items-center space-x-4">
            <Search />
          </div>
          <nav className="flex items-center space-x-1">
            <Link
              href={siteConfig.links.telegram}
              target="_blank"
              rel="noreferrer"
            >
              <div
                className={cn(
                  buttonVariants({
                    size: "sm",
                    variant: "ghost",
                  }),
                  "w-9 px-0"
                )}
              >
                <Image
                  src="/telegram-fill.svg"
                  alt="Telegram"
                  width={20}
                  height={20}
                />
                <span className="sr-only">Telegram</span>
              </div>
            </Link>
          </nav>
        </div>
      </div>
    </header>
  );
}
