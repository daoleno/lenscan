import Link from "next/link"
import { usePathname } from "next/navigation"
import { Leaf } from "lucide-react"

import { siteConfig } from "@/config/site"
import { cn } from "@/lib/utils"

import { Badge } from "./ui/badge"

const items = [
  {
    name: "Overview",
    href: "/",
  },
  {
    name: "Publications",
    href: "/publications",
  },
  {
    name: "Momoka",
    href: "/momoka-txs",
  },
  {
    name: "Profiles",
    href: "/profiles",
  },
  // {
  //   name: "Analytics",
  //   href: "/analytics",
  // },
]

export function MainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  const pathname = usePathname()
  return (
    <div className={cn("mr-4 md:flex", className)} {...props}>
      <Link href="/" className="mr-6 flex items-center space-x-2">
        <Leaf className="h-6 w-6" />
        <span className="font-bold sm:inline-block">
          {siteConfig.name}
          <Badge variant="outline" className="ml-1">
            beta
          </Badge>
        </span>
      </Link>
      <nav className="hidden items-center space-x-6 text-sm font-medium sm:flex">
        {items.map((item) => (
          <Link
            key={item.name}
            href={item.href}
            className={cn(
              "transition-colors hover:text-foreground/80",
              pathname === item.href ? "text-foreground" : "text-foreground/60"
            )}
          >
            {item.name}
          </Link>
        ))}
      </nav>
    </div>
  )
}
