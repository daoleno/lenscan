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
		<div className={cn("flex", className)} {...props}>
			<Link href="/" className="mr-6 flex items-center space-x-2">
				<Leaf className="h-6 w-6" />
				<span className="flex items-center gap-1 font-bold">
					{siteConfig.name}
					<Badge variant="outline" className="ml-1">
						beta
					</Badge>
				</span>
			</Link>
			<NavigationMenu className="hidden md:flex">
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
							<ul className="grid w-[280px] gap-3 p-4 md:w-[400px] lg:w-[500px]">
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
			
			<div className="flex md:hidden flex-col space-y-4 mt-4">
				{items.map((item) => (
					<Link
						key={item.name}
						href={item.href}
						className="text-muted-foreground hover:text-primary"
					>
						{item.name}
					</Link>
				))}
				<div className="space-y-3">
					<p className="font-medium">Revenue</p>
					{revenueComponents.map((component) => (
						<Link
							key={component.title}
							href={component.href}
							className="block text-muted-foreground hover:text-primary"
						>
							{component.title}
						</Link>
					))}
				</div>
			</div>
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
