"use client"

import Link from "next/link"
import { CommentFragment } from "@lens-protocol/client"
import {
  BarChart,
  CheckCircle2,
  ClipboardCheck,
  FileText,
  Focus,
  XCircle,
} from "lucide-react"
import ReactJson from "react-json-view"

import { Badge } from "./ui/badge"
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card"

export default function Comment({ comment }: { comment: CommentFragment }) {
  const overviewItems = [
    { label: "Id", value: comment.id },
    { label: "Title", value: (comment.metadata as any)?.title || "-" },
    { label: "Type", value: comment.__typename },
    { label: "App", value: comment.metadata?.appId },
    {
      label: "On Polygon",
      value: comment.momoka?.__typename !== "MomokaInfo" ? "Yes" : "No",
    },
    { label: "Polygon Tx Hash", value: comment.txHash || "-" },
    {
      label: "On Momoka",
      value: comment.momoka?.__typename == "MomokaInfo" ? "Yes" : "No",
    },
    { label: "Momoka Proof", value: comment.momoka?.proof || "-" },
  ]

  const operationItems = [
    { label: "Can Comment", value: comment.operations.canComment },
    { label: "Can Mirror", value: comment.operations.canMirror },
    { label: "Can Act", value: comment.operations.canAct },
    { label: "Can Decrypt", value: comment.operations.canDecrypt.result },
    { label: "Has Bookmarked ", value: comment.operations.hasBookmarked },
    { label: "Has Reported ", value: comment.operations.hasReported },
    { label: "Has Upvoted ", value: comment.operations.hasUpvoted },
    { label: "Has Downvoted ", value: comment.operations.hasDownvoted },
    { label: "Hidden", value: comment.isHidden },
    { label: "Is Not Interested ", value: comment.operations.isNotInterested },
  ]

  return (
    <div className="flex flex-col space-y-7 py-7">
      <div className="flex flex-col space-y-2">
        <div className="flex items-center text-2xl font-bold">
          <span>Publication</span>
          <span className="ml-2 font-mono">{comment.id}</span>
        </div>
        <div className="text-sm font-bold text-muted-foreground">
          <Badge>{comment.__typename}</Badge>
          <span> @ </span>
          <Link
            href={`/profile/${comment.by.handle?.localName}`}
            className="font-bold underline underline-offset-4"
          >
            {comment.by.handle?.localName}
          </Link>
        </div>
        <div className="ml-1 flex flex-col space-y-1">
          <span className="text-sm">Comment On</span>
          <Link
            className="font-mono underline underline-offset-4"
            href={`/publication/${comment.commentOn?.id}`}
          >
            {comment.commentOn?.id || "-"}
            {(comment.commentOn.metadata as any)?.title &&
              ` (${(comment.commentOn.metadata as any).title})`}
          </Link>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Overview</CardTitle>
            <Focus />
          </CardHeader>
          <CardContent className="grid grid-cols-2 gap-4">
            {overviewItems.map((item, index) => (
              <div key={index} className="flex flex-col space-y-1 font-medium">
                <span className="text-sm uppercase text-muted-foreground">
                  {item.label}
                </span>
                <span>
                  {typeof item.value === "object" ? (
                    <ReactJson
                      name={false}
                      collapsed={true}
                      displayDataTypes={false}
                      src={item.value || {}}
                    />
                  ) : (
                    item.value
                  )}
                </span>
              </div>
            ))}
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Stats</CardTitle>
            <BarChart />
          </CardHeader>
          <CardContent className="grid grid-cols-2 gap-4">
            {Object.entries(comment.stats)
              .filter(([key]) => !key.startsWith("__"))
              .map(([key, value]) => (
                <div key={key} className="flex flex-col space-y-1 font-medium">
                  <span className="text-sm uppercase text-muted-foreground">
                    {key
                      .replace(/^total/, "")
                      .replace(/([a-z])([A-Z])/g, "$1 $2")}
                  </span>
                  <span>{value}</span>
                </div>
              ))}
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Checks</CardTitle>
            <ClipboardCheck />
          </CardHeader>
          <CardContent className="grid grid-cols-2 gap-4">
            {operationItems.map((item, index) => (
              <div key={index} className="flex flex-col space-y-1 font-medium">
                <span className="text-sm uppercase text-muted-foreground">
                  {item.label}
                </span>
                <span className="flex flex-row items-center space-x-2">
                  {typeof item.value === "boolean" ? (
                    item.value ? (
                      <CheckCircle2 className="h-4 w-4" />
                    ) : (
                      <XCircle className="h-4 w-4 text-muted-foreground" />
                    )
                  ) : item.value === "YES" ? (
                    <CheckCircle2 className="h-4 w-4" />
                  ) : item.value === "NO" ? (
                    <XCircle className="h-4 w-4 text-muted-foreground" />
                  ) : (
                    item.value
                  )}
                </span>
              </div>
            ))}
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Metadata</CardTitle>
            <FileText />
          </CardHeader>
          <CardContent className="overflow-auto">
            {/* <Metadata metadata={comment.metadata} /> */}
            <ReactJson
              name={false}
              displayDataTypes={false}
              src={comment.metadata}
            />
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
