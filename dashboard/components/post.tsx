import { Post } from "@lens-protocol/react-web";
import {
  BarChart,
  CheckCircle2,
  ClipboardCheck,
  FileText,
  Focus,
  XCircle,
} from "lucide-react";
import dynamic from "next/dynamic";
import Link from "next/link";
import { Badge } from "./ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
("lucide-react");

const DynamicReactJson = dynamic(import("react-json-view"), { ssr: false });

export default function Post({ post }: { post: Post }) {
  const overviewItems = [
    { label: "Id", value: post.id },
    { label: "Name", value: post.metadata.name },
    { label: "Type", value: post.__typename },
    { label: "CreatedAt At", value: post.createdAt },
    { label: "Collect Policy", value: post.collectPolicy || "-" },
    { label: "Reference Policy", value: post.referencePolicy || "-" },
    { label: "Collect Module", value: post.collectModule },
    { label: "Mirrors", value: post.mirrors || "-" },
    { label: "Reaction", value: post.reaction || "-" },
    { label: "Collected By", value: post.collectedBy?.address || "-" },
    { label: "Decryption Criteria", value: post.decryptionCriteria || "-" },
  ];

  const checkItems = [
    { label: "Can Comment", value: post.canComment.result },
    { label: "Can Mirror", value: post.canMirror.result },
    { label: "Can Observer Decrypt", value: post.canObserverDecrypt.result },
    { label: "Hidden", value: post.hidden },
    { label: "Is Gated", value: post.isGated },

    { label: "Has Collected By Me", value: post.hasCollectedByMe },
    {
      label: "Has Optimistic Collected By Me",
      value: post.hasOptimisticCollectedByMe,
    },
    {
      label: "Is Optimistic Mirrored By Me",
      value: post.isOptimisticMirroredByMe,
    },
  ];

  return (
    <div className="flex flex-col space-y-7 py-7">
      <div className="flex flex-col space-y-2">
        <div className="flex items-center text-2xl font-bold text-gray-800">
          <span>Publication</span>
          <span className="ml-2 font-mono">{post.id}</span>
        </div>
        <div className="text-sm font-bold text-gray-600">
          <Badge>{post.__typename}</Badge>
          <span> @ </span>
          <Link
            href={`/profile/${post.profile.handle}`}
            className="font-bold underline underline-offset-4"
          >
            {post.profile.handle}
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
                <span className="text-sm uppercase text-gray-600">
                  {item.label}
                </span>
                <span>
                  {typeof item.value === "object" ? (
                    <DynamicReactJson
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
                  <span className="text-sm uppercase text-gray-600">
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
            {checkItems.map((item, index) => (
              <div key={index} className="flex flex-col space-y-1 font-medium">
                <span className="text-sm uppercase text-gray-600">
                  {item.label}
                </span>
                <span className="flex flex-row items-center space-x-2">
                  {typeof item.value === "boolean" ? (
                    item.value ? (
                      <CheckCircle2 className="h-4 w-4" />
                    ) : (
                      <XCircle className="h-4 w-4 text-muted-foreground" />
                    )
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
            <DynamicReactJson
              name={false}
              displayDataTypes={false}
              src={post.metadata}
            />
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
