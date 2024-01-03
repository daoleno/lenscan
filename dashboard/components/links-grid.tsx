import React from "react"
import Link from "next/link"

type LinksProps = {
  links: { url: string; linkText: string }[]
}

const LinksGrid: React.FC<LinksProps> = ({ links }) => (
  <div className="flex flex-col gap-3 py-4">
    {links.map((link, idx) => (
      <Link
        key={idx}
        href={link.url}
        passHref
        className="flex flex-col gap-1 hover:underline"
        target="_blank"
      >
        <span className="font-bold">{link.linkText}</span>
        <span className="mx-3 text-muted-foreground">{link.url}</span>
      </Link>
    ))}
  </div>
)

export default LinksGrid
