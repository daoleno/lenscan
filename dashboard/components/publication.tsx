import {
  MetadataAttributeOutputFragment,
  MetadataFragment,
  PostFragment,
} from "@lens-protocol/api-bindings";
import { usePublication } from "@lens-protocol/react-web";
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
import NotFound from "./404";
import { Loading } from "./loading";
import { Badge } from "./ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
("lucide-react");

const DynamicReactJson = dynamic(import("react-json-view"), { ssr: false });

const Metadata: React.FC<{ metadata: MetadataFragment }> = ({ metadata }) => {
  return (
    <div className="mx-auto">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-bold">{metadata.name}</h2>
        {metadata.image && (
          <img
            src={metadata.image}
            alt="metadata"
            className="w-16 h-16 object-contain"
          />
        )}
      </div>
      <p className="text-gray-700 mb-4">{metadata.description}</p>
      <div className="mb-4">
        <h3 className="text-lg font-bold mb-2">Attributes</h3>
        <ul>
          {metadata.attributes.map(
            (attribute: MetadataAttributeOutputFragment, index: number) => (
              <li key={index}>
                <span className="font-bold">{attribute.traitType}: </span>
                <span>{attribute.value}</span>
              </li>
            )
          )}
        </ul>
      </div>
      <div className="mb-4">
        <h3 className="text-lg font-bold mb-2">Media</h3>
        <ul>
          {metadata.media.map((media, index) => (
            <li key={index}>
              {media.original && (
                <img
                  src={media.original.url}
                  // alt={media.description}
                  alt="media"
                  className="w-16 h-16 object-contain"
                />
              )}
            </li>
          ))}
        </ul>
      </div>
      <div className="mb-4">
        <h3 className="text-lg font-bold mb-2">Content</h3>
        <div
          className="prose"
          dangerouslySetInnerHTML={{ __html: metadata.content ?? "" }}
        ></div>
      </div>
    </div>
  );
};

function PostCard({ post }: { post: PostFragment }) {
  const overviewItems = [
    { label: "Id", value: post.id },
    { label: "Name", value: post.metadata.name },
    { label: "Type", value: post.__typename },
    { label: "CreatedAt At", value: post.createdAt },
    { label: "Collect Policy", value: post.collectPolicy || "-" },
    { label: "Reference Policy", value: post.referencePolicy || "-" },
    { label: "Collect Module", value: post.__collectModule },
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
        <div className="text-2xl font-bold text-gray-800 flex items-center">
          <span>Publication</span>
          <span className="ml-2 font-mono">{post.id}</span>
        </div>
        <div className="font-bold text-gray-600 text-sm">
          <Badge>{post.__typename}</Badge>
          <span className=""> @ </span>
          <Link
            href={`/profiles/${post.profile.handle}`}
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
                <span className="text-sm text-gray-600 uppercase">
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
                  <span className="text-sm text-gray-600 uppercase">
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
                <span className="text-sm text-gray-600 uppercase">
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

        <Card className="hover:shadow-lg">
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

export default function Publication({ id }: { id: string }) {
  const { data: post, loading, error } = usePublication({ publicationId: id });

  console.log({ post, loading, error });

  if (loading) {
    return <Loading />;
  }

  if (error) {
    console.error(error);
    return <NotFound type="Publication" />;
  }

  return <PostCard post={post as PostFragment} />;
}
