import Link from "next/link"
import { FaTelegram } from "react-icons/fa"

import { siteConfig } from "@/config/site"
import { cn } from "@/lib/utils"
import { MainNav } from "@/components/main-nav"
import { buttonVariants } from "@/components/ui/button"

import { Icons } from "./icons"
import { Search } from "./search"
import { ThemeToggle } from "./theme-toggle"

export function SiteHeader() {
  return (
    <header className="supports-backdrop-blur:bg-background/60 sticky top-0 z-40 w-full border-b bg-background/95 shadow-sm backdrop-blur">
      <div className="container flex h-14 items-center">
        <MainNav />

        <div className="flex flex-1 items-center justify-between space-x-2 sm:space-x-4 md:justify-end">
          <div className="ml-auto flex items-center space-x-4">
            <Search />
          </div>
          <nav className="flex items-center">
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
                  })
                )}
              >
                <FaTelegram className="h-5 w-5" />
                <span className="sr-only">Telegram</span>
              </div>
            </Link>
            <Link
              href={siteConfig.links.github}
              target="_blank"
              rel="noreferrer"
            >
              <div
                className={cn(
                  buttonVariants({
                    size: "sm",
                    variant: "ghost",
                  })
                )}
              >
                <Icons.gitHub className="h-5 w-5" />
                <span className="sr-only">Github</span>
              </div>
            </Link>
            <ThemeToggle />
          </nav>
        </div>
      </div>
    </header>
  )
}
