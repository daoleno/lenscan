"use client"

import { cn } from "@/lib/utils"
import { Leaf } from "lucide-react"
import Link from "next/link"
import * as React from "react"

import { Badge } from "./ui/badge"

import { siteConfig } from "@/config/site"

import {
	NavigationMenu,
	NavigationMenuContent,
	NavigationMenuItem,
	NavigationMenuLink,
	NavigationMenuList,
	NavigationMenuTrigger,
	navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu"

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
		name: "Profiles",
		href: "/profiles",
	},
	{
		name: "Apps",
		href: "/apps",
	},
	// {
	// 	name: "Revenue",
	// 	href: "/revenue",
	// },
	{
		name: "Analytics",
		href: "/analytics",
	},
]

const revenueComponents: {
	title: string
	href: string
	description: string
}[] = [
	{
		title: "App Revenue",
		href: "/revenue/app",
		description: "Lens Protocol and Lens Apps Revenue Data",
	},
	{
		title: "Profile Revenue",
		href: "/revenue/profile",
		description:
			"Revenue data for profiles are distributed from many Lens-enabled tokens",
	},
]
export function MainNav({
	className,
	...props
}: React.HTMLAttributes<HTMLElement>) {
	return (
		<div className={cn("mr-4 md:flex", className)} {...props}>
			<Link href="/" className="mr-6 flex items-center space-x-2">
				<Leaf className="h-6 w-6" />
				<span className="flex gap-1 font-bold sm:inline-block">
					{siteConfig.name}
					<Badge variant="outline" className="ml-1">
						beta
					</Badge>
				</span>
			</Link>
			<NavigationMenu>
				<NavigationMenuList>
					{items.map((item) => (
						<NavigationMenuItem key={item.name}>
							<Link href={item.href} passHref>
								<NavigationMenuLink className={navigationMenuTriggerStyle()}>
									{item.name}
								</NavigationMenuLink>
							</Link>
						</NavigationMenuItem>
					))}
					<NavigationMenuItem>
						<NavigationMenuTrigger>Revenue</NavigationMenuTrigger>
						<NavigationMenuContent>
							<ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-1 lg:w-[600px] ">
								{revenueComponents.map((component) => (
									<ListItem
										key={component.title}
										title={component.title}
										href={component.href}
									>
										{component.description}
									</ListItem>
								))}
							</ul>
						</NavigationMenuContent>
					</NavigationMenuItem>
				</NavigationMenuList>
			</NavigationMenu>
		</div>
	)
}

const ListItem = React.forwardRef<
	React.ElementRef<"a">,
	React.ComponentPropsWithoutRef<"a">
>(({ className, title, children, ...props }, ref) => {
	return (
		<li>
			<NavigationMenuLink asChild>
				<a
					ref={ref}
					className={cn(
						"block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
						className,
					)}
					{...props}
				>
					<div className="text-sm font-medium leading-none">{title}</div>
					<p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
						{children}
					</p>
				</a>
			</NavigationMenuLink>
		</li>
	)
})
ListItem.displayName = "ListItem"
