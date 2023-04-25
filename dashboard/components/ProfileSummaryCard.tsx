import { Color } from "@tremor/react";

import { useProfile } from "@lens-protocol/react-web";
import { Loading } from "./loading";

type Categories = {
  title: string;
  metric: string;
  icon: any;
  color: Color;
}[];

export default function ProfileSummaryCard({
  profileId,
}: {
  profileId: string;
}) {
  const { data: profile, loading, error } = useProfile({ profileId });
  console.log("profile", profile);
  if (loading) return <Loading />;
  if (error) return <div>Error: {error.message} </div>;
  if (!profile) return <Loading />;

  return (
    <div className="bg-gray-200 p-4 rounded-lg">
      <div className="flex flex-col md:flex-row md:items-center">
        <img
          src={getIPFSURL(profile.picture)}
          // alt={profile.name}
          className="h-24 w-24 rounded-full md:mr-4"
        />
        <div className="flex flex-col justify-center">
          <h1 className="text-xl font-bold">{profile.name}</h1>
          <p className="text-gray-600">{`@${profile.handle}`}</p>
          <p className="text-gray-600">{profile.bio}</p>
        </div>
      </div>
      <div className="mt-4">
        <img
          src={getIPFSURL(profile.coverPicture)}
          // alt={`Cover picture for ${profile.name}`}
          className="w-full h-32 object-cover rounded-lg"
        />
      </div>
      <div className="mt-4">
        <h2 className="text-lg font-bold mb-2">Profile Attributes</h2>
        <ul className="list-disc pl-4">
          {profile.__attributes!.map(({ key, value }) => (
            <li key={key}>
              <span className="font-bold">{key}:</span> {value}
            </li>
          ))}
        </ul>
      </div>
      <div className="mt-4">
        <h2 className="text-lg font-bold mb-2">Profile Stats</h2>
        <ul className="list-disc pl-4">
          <li>Total Collects: {profile.stats.totalCollects}</li>
          <li>Total Comments: {profile.stats.totalComments}</li>
          <li>Total Followers: {profile.stats.totalFollowers}</li>
          <li>Total Following: {profile.stats.totalFollowing}</li>
          <li>Total Mirrors: {profile.stats.totalMirrors}</li>
          <li>Total Posts: {profile.stats.totalPosts}</li>
          <li>Total Publications: {profile.stats.totalPublications}</li>
        </ul>
      </div>
      <div className="mt-4">
        <h2 className="text-lg font-bold mb-2">On-Chain Identity</h2>
        <ul className="list-disc pl-4">
          <li>
            Proof of Humanity:{" "}
            {profile.onChainIdentity.proofOfHumanity ? "Yes" : "No"}
          </li>
          <li>ENS Name: {String(profile.onChainIdentity.ens!.name) || "-"}</li>
          <li>
            Sybil.org Verified:{" "}
            {profile.onChainIdentity.sybilDotOrg.verified ? "Yes" : "No"}
          </li>
          <li>
            Twitter Handle:{" "}
            {profile.onChainIdentity.sybilDotOrg.source.twitter.handle || "-"}
          </li>
        </ul>
      </div>
      <div className="mt-4">
        <h2 className="text-lg font-bold mb-2">Dispatcher</h2>
        <ul className="list-disc pl-4">
          <li>Address: {profile.dispatcher?.address}</li>
          <li>
            Can Use Relay: {profile.dispatcher?.canUseRelay ? "Yes" : "No"}
          </li>
        </ul>
      </div>
    </div>
  );
}

// url: ipfs://bafybeiewog3iscltj6uvus6iut5kerbbkyxovjhvnikrc4luy5sap6w3zu
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
