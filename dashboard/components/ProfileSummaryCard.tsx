/* eslint-disable @next/next/no-img-element */

import { formatNumber, shortHash } from "@/lib/utils";
import { ProfileFragment } from "@lens-protocol/api-bindings";
import {
  CheckCircle2,
  Fingerprint,
  Tags,
  Webhook,
  XCircle,
} from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import Balance from "react-wrap-balancer";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

export default function ProfileSummaryCard({
  profile,
}: {
  profile: ProfileFragment;
}) {
  if (!profile) return null;
  return (
    <>
      <div>
        <div className="h-2/4 sm:h-64 overflow-hidden mt-6 rounded-lg">
          {profile.coverPicture ? (
            <img
              className="w-full object-cover"
              src={getIPFSURL(profile.coverPicture)}
              alt="cover"
            />
          ) : (
            <div className="w-full h-full bg-gradient-to-r from-rose-100 to-teal-100" />
          )}
        </div>
        <div className="flex sm:flex-row flex-col">
          <div className="flex flex-col -mt-16 sm:w-1/2 ">
            <div className="flex px-5 mb-5">
              {profile.picture ? (
                <img
                  className="object-cover rounded-full h-32 w-32"
                  src={getIPFSURL(profile.picture)}
                  alt="profile"
                />
              ) : (
                <Image
                  src="/images/default-profile.png"
                  alt="profile"
                  width={128}
                  height={128}
                  className="object-cover rounded-full h-32 w-32"
                />
              )}
            </div>
            <div className="flex flex-col px-7 mb-8">
              <h2 className="text-3xl font-bold text-gray-900">
                {profile.name}
              </h2>
              <Balance className="text-gray-400 mt-2 dark:text-gray-400">
                @{profile.handle}
              </Balance>
              <p className="mt-2 text-gray-600">{profile.bio}</p>
            </div>
          </div>
          <div className="grid grid-cols-3 gap-6 sm:grid-cols-2 lg:grid-cols-4 py-7 sm:w-1/2">
            {Object.entries(profile.stats)
              .filter(
                ([key]) =>
                  !key.startsWith("__") &&
                  key !== "commentsCount" &&
                  key !== "postsCount" &&
                  key !== "mirrorsCount"
              )
              .map(([key, value]) => (
                <Card key={key} className="rounded-xl overflow-scroll">
                  <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                    <CardTitle className="text-sm font-medium">
                      {key.replace(/^total/, "")}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="text-2xl font-bold">
                      {formatNumber(value)}
                    </div>
                  </CardContent>
                </Card>
              ))}
          </div>
        </div>
      </div>
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Categories</CardTitle>
            <Tags />
          </CardHeader>

          <CardContent className="flex flex-col space-y-3 overflow-scroll">
            {profile.__attributes!.map(
              ({ key, value }) =>
                value !== "[]" && (
                  <div
                    key={key}
                    className="flex flex-col space-y-1 font-medium"
                  >
                    <span className="text-sm text-gray-600 uppercase">
                      {key.replace(/([a-z])([A-Z])/g, "$1 $2")}
                    </span>
                    {key.toLocaleLowerCase().includes("website") ? (
                      <Link href={value}>{value} </Link>
                    ) : key.toLocaleLowerCase().includes("twitter") ? (
                      <Link href={`https://twitter.com/${value}`}>{value}</Link>
                    ) : (
                      <span>{value}</span>
                    )}
                  </div>
                )
            )}
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>On-Chain Identity</CardTitle>
            <Fingerprint />
          </CardHeader>
          <CardContent className="flex flex-col space-y-3">
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                Proof of Humanity
              </span>
              {profile.onChainIdentity.proofOfHumanity ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                ENS Name
              </span>
              <span className="font-medium">
                {profile.onChainIdentity.ens?.name
                  ? String(profile.onChainIdentity.ens.name)
                  : "-"}
              </span>
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                Sybil.org Verified
              </span>
              {profile.onChainIdentity.sybilDotOrg.verified ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                Twitter Handle
              </span>
              <span className="font-medium">
                {profile.onChainIdentity.sybilDotOrg.source.twitter.handle ||
                  "-"}
              </span>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle>Dispatcher</CardTitle>
            <Webhook />
          </CardHeader>
          <CardContent className="flex flex-col space-y-3">
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                Address
              </span>
              <Link
                href={`https://polygonscan.com/address/${profile.dispatcher?.address}`}
                target="_blank"
                className="font-medium underline underline-offset-4"
              >
                {shortHash(profile.dispatcher?.address)}
              </Link>
            </div>
            <div className="flex flex-col space-y-1">
              <span className="text-sm font-medium text-gray-600 uppercase">
                Can Use Relay
              </span>
              {profile.dispatcher?.canUseRelay ? (
                <CheckCircle2 className="h-4 w-4" />
              ) : (
                <XCircle className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  );
}

const ipfsGateway = "https://lens.infura-ipfs.io";
function getIPFSURL(picture: any) {
  let url = "";
  if (!picture) {
    return url;
  }
  if (picture.__typename === "MediaSet") {
    url = picture.original.url;
  }
  if (picture.__typename === "NftImage") {
    url = picture.uri;
  }

  if (url && url.startsWith("ipfs://")) {
    const cid = url.replace("ipfs://", "");
    return `${ipfsGateway}/ipfs/${cid}`;
  }
  return url;
}
