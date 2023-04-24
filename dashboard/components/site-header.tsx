import Link from "next/link";

import { MainNav } from "@/components/main-nav";
import { buttonVariants } from "@/components/ui/button";
import { siteConfig } from "@/config/site";
import { cn } from "@/lib/utils";
import { Twitter } from "lucide-react";
import Image from "next/image";
import { Search } from "./search";

export function SiteHeader() {
  return (
    <header className="supports-backdrop-blur:bg-background/60 sticky top-0 z-40 w-full border-b bg-background/95 shadow-sm backdrop-blur">
      <div className="container flex h-14 items-center">
        <MainNav />

        {/* <MobileNav /> */}
        <div className="flex flex-1 items-center justify-between space-x-2 sm:space-x-4 md:justify-end">
          <div className="ml-auto flex items-center space-x-4">
            <Search />
          </div>
          <nav className="flex items-center space-x-1">
            <Link
              href={siteConfig.links.lenster}
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
                {/* <Github className="h-5 w-5" /> */}
                <Image
                  src="/lenster.svg"
                  alt="Lenster"
                  width={20}
                  height={20}
                />
                <span className="sr-only">Lenster</span>
              </div>
            </Link>
            <Link
              href={siteConfig.links.twitter}
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
                <Twitter className="h-5 w-5 fill-current" />
                <span className="sr-only">Twitter</span>
              </div>
            </Link>
            {/* <ModeToggle /> */}
          </nav>
        </div>
      </div>
    </header>
  );
}
