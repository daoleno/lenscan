import { siteConfig } from "@/config/site";
import { Twitter } from "lucide-react";
import Image from "next/image";
import Link from "next/link";

const navigation = [
  {
    name: "Twitter",
    href: siteConfig.links.twitter,
    icon: <Twitter className="h-5 w-5 fill-current" />,
  },
  {
    name: "Lenster",
    href: siteConfig.links.lenster,
    icon: <Image src="/lenster.svg" alt="Lenster" width={20} height={20} />,
  },
  {
    name: "Telegram",
    href: siteConfig.links.telegram,
    icon: (
      <Image src="/telegram-fill.svg" alt="Telegram" width={20} height={20} />
    ),
  },
];

export default function Footer() {
  return (
    <footer className="container">
      <div className="pt-12 md:flex md:items-center md:justify-between ">
        <div className="flex justify-center space-x-6 md:order-2">
          {navigation.map((item) => (
            <Link key={item.name} href={item.href}>
              <span className="sr-only">{item.name}</span>
              {item.icon}
            </Link>
          ))}
        </div>
        <div className="mt-8 font-medium md:order-1 md:mt-0">
          <p className="text-center leading-5 text-gray-500">
            &copy; 2023 lenscan. Made with ☘️ by
            <Link
              href={siteConfig.links.lenster}
              className="ml-2 font-medium text-gray-700 underline"
            >
              daoleno
            </Link>
          </p>
        </div>
      </div>
    </footer>
  );
}
