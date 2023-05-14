import NotFound from "@/components/404";
import ProfileSummaryCard from "@/components/ProfileSummaryCard";
import ProfileTableView from "@/components/ProfileTableView";
import { Error } from "@/components/error";
import { Loading } from "@/components/loading";
import { ProfileId, useProfile } from "@lens-protocol/react-web";
import { ethers } from "ethers";
import { useRouter } from "next/router";

export default function Profile() {
  const { id } = useRouter().query;
  const q = id?.toString();
  const isProfileId = ethers.utils.isHexString(q);
  const realhandle =
    q == "lensprotocol" ? q : q?.endsWith(".lens") ? q : q + ".lens";
  const handle = isProfileId ? "" : realhandle;
  const {
    data: idProfile,
    loading: idProfileLoading,
    error: idError,
  } = useProfile({
    profileId: q as ProfileId,
  });
  const {
    data: handleProfile,
    loading: handleProfileLoading,
    error: handleError,
  } = useProfile({
    handle: handle,
  });
  const profileId = isProfileId ? q : handleProfile?.id;
  const profile = isProfileId ? idProfile : handleProfile;
  const hasProfile = profileId && profile;
  const isLoading = idProfileLoading || handleProfileLoading;
  const error = idError && handleError;

  if (isLoading) {
    return <Loading />;
  }

  if (!hasProfile) {
    return <NotFound type="Profile" />;
  }

  if (error) {
    return <Error msg={error.message} />;
  }

  return (
    <>
      <ProfileSummaryCard profile={profile} />
      <ProfileTableView profileId={profileId} />
    </>
  );
}
