import Link from "next/link"
import { MirrorFragment } from "@lens-protocol/client"
import { CheckCircle2, XCircle } from "lucide-react"

import { Badge } from "./ui/badge"

export default function Mirror({ mirror }: { mirror: MirrorFragment }) {
  const overviewItems = [
    {
      label: "Mirror On",
      value: `${mirror.mirrorOn.id}} (by ${mirror.mirrorOn.by.handle?.localName})`,
      href: `/publication/${mirror.mirrorOn.id}`,
      type: "link",
    },
    { label: "Published On", value: mirror.publishedOn?.id, type: "text" },
    { label: "Created At", value: mirror.createdAt, type: "text" },
  ]
  return (
    <div className="flex flex-col space-y-7 py-7">
      <div className="flex flex-col space-y-2">
        <div className="flex items-center text-2xl font-bold text-gray-800">
          <span>Publication</span>
          <span className="ml-2 font-mono">{mirror.id}</span>
        </div>
        <div className="text-sm font-bold text-gray-600">
          <Badge>{mirror.__typename}</Badge>
          <span> @ </span>
          <Link
            href={`/profile/${mirror.by.handle?.id}`}
            className="font-bold underline underline-offset-4"
          >
            {mirror.by.handle?.localName}
          </Link>
        </div>
      </div>
      <div className="flex flex-col space-y-2">
        {overviewItems.map((item, index) => (
          <div key={index} className="flex flex-col space-y-2">
            <span className="text-sm">{item.label}</span>
            {item.type === "link" ? (
              <Link
                className="font-mono underline underline-offset-4"
                href={item.href || "#"}
              >
                {item.value || "-"}
              </Link>
            ) : item.type === "boolean" ? (
              item.value ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )
            ) : (
              <span className="font-mono">{item.value || "-"}</span>
            )}
          </div>
        ))}
      </div>
      {/* <DynamicReactJson src={mirror.mirrorOf} /> */}
    </div>
  )
}
