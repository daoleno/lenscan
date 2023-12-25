"use client"

import Link from "next/link"
import { PostFragment } from "@lens-protocol/client"
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

export default function Post({ post }: { post: PostFragment }) {
  const overviewItems = [
    { label: "Id", value: post.id },
    // @ts-ignore
    { label: "Title", value: post.metadata?.title },
    { label: "Type", value: post.__typename },
    { label: "App", value: post.metadata?.appId },
    {
      label: "On Polygon",
      value: post.momoka?.__typename !== "MomokaInfo" ? "Yes" : "No",
    },
    { label: "Polygon Tx Hash", value: post.txHash || "-" },
    {
      label: "On Momoka",
      value: post.momoka?.__typename == "MomokaInfo" ? "Yes" : "No",
    },
    { label: "Momoka Proof", value: post.momoka?.proof || "-" },
  ]

  const operationItems = [
    { label: "Can Comment", value: post.operations.canComment },
    { label: "Can Mirror", value: post.operations.canMirror },
    { label: "Can Act", value: post.operations.canAct },
    { label: "Can Decrypt", value: post.operations.canDecrypt.result },
    { label: "Has Bookmarked ", value: post.operations.hasBookmarked },
    { label: "Has Reported ", value: post.operations.hasReported },
    { label: "Has Upvoted ", value: post.operations.hasUpvoted },
    { label: "Has Downvoted ", value: post.operations.hasDownvoted },
    { label: "Hidden", value: post.isHidden },
    { label: "Is Not Interested ", value: post.operations.isNotInterested },
  ]

  return (
    <div className="flex flex-col space-y-7 py-7">
      <div className="flex flex-col space-y-2">
        <div className="flex items-center text-2xl font-bold">
          <span>Publication</span>
          <span className="ml-2 font-mono">{post.id}</span>
        </div>
        <div className="text-sm font-bold text-muted-foreground">
          <Badge>{post.__typename}</Badge>
          <span> @ </span>
          <Link
            href={`/profile/${post.by.handle?.localName}`}
            className="font-bold underline underline-offset-4"
          >
            {post.by.handle?.localName}
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
                <span className="overflow-auto">
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
            {Object.entries(post.stats)
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
            <CardTitle>Operations</CardTitle>
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
            {/* <Metadata metadata={post.metadata} /> */}
            {typeof window !== "undefined" && (
              <ReactJson
                name={false}
                displayDataTypes={false}
                collapsed={true}
                src={post.metadata}
              />
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
