import NotFound from "@/components/404";
import Layout from "@/components/Layout";
import ProfileSummaryCard from "@/components/ProfileSummaryCard";
import ProfileTableView from "@/components/ProfileTableView";
import { Error } from "@/components/error";
import { Loading } from "@/components/loading";
import { useProfile } from "@lens-protocol/react-web";
import { useRouter } from "next/router";

export default function Profile() {
  const { id } = useRouter().query;
  const q = id?.toString();
  const isProfileId = q?.startsWith("0x");
  const realhandle =
    q == "lensprotocol" ? q : q?.endsWith(".lens") ? q : q + ".lens";
  const handle = isProfileId ? "" : realhandle;
  const {
    data: idProfile,
    loading: idProfileLoading,
    error: idError,
  } = useProfile({
    profileId: q || "",
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
    return (
      <Layout>
        <Loading />
      </Layout>
    );
  }

  if (!hasProfile) {
    return (
      <Layout>
        <NotFound type="Profile" />
      </Layout>
    );
  }

  if (error) {
    return (
      <Layout>
        <Error msg={error.message} />
      </Layout>
    );
  }

  return (
    <Layout>
      <ProfileSummaryCard profile={profile} />
      <ProfileTableView profileId={profileId} />
    </Layout>
  );
}
