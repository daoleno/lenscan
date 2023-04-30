import { MirrorFragment } from "@lens-protocol/api-bindings";
import { CheckCircle2, XCircle } from "lucide-react";
import Link from "next/link";
import { Badge } from "./ui/badge";
("lucide-react");

export default function Mirror({ mirror }: { mirror: MirrorFragment }) {
  const overviewItems = [
    {
      label: "Mirror Of",
      value: `${mirror.mirrorOf?.id} (${mirror.mirrorOf?.metadata?.name})`,
      href: `/publication/${mirror.mirrorOf?.id}`,
      type: "link",
    },
    { label: "Hidden", value: mirror.hidden, type: "boolean" },
    { label: "Created At", value: mirror.createdAt, type: "text" },
  ];
  return (
    <div className="flex flex-col space-y-7 py-7">
      <div className="flex flex-col space-y-2">
        <div className="text-2xl font-bold text-gray-800 flex items-center">
          <span>Publication</span>
          <span className="ml-2 font-mono">{mirror.id}</span>
        </div>
        <div className="font-bold text-gray-600 text-sm">
          <Badge>{mirror.__typename}</Badge>
          <span> @ </span>
          <Link
            href={`/profile/${mirror.profile.handle}`}
            className="font-bold underline underline-offset-4"
          >
            {mirror.profile.handle}
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
  );
}
