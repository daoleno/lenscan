import { PublicationId, usePublication } from "@lens-protocol/react-web";
import NotFound from "./404";
import Comment from "./comment";
import DynamicReactJson from "./dynamic-react-json";
import { Loading } from "./loading";
import Mirror from "./mirror";
import Post from "./post";
("lucide-react");

export default function Publication({ id }: { id: string }) {
  const {
    data: pub,
    loading,
    error,
  } = usePublication({ publicationId: id as PublicationId });

  if (loading) {
    return <Loading />;
  }

  if (error) {
    console.error(error);
    return <NotFound type="Publication" />;
  }

  return (
    <>
      {pub?.__typename === "Post" ? (
        <Post post={pub} />
      ) : pub?.__typename === "Comment" ? (
        <Comment comment={pub} />
      ) : pub?.__typename === "Mirror" ? (
        <Mirror mirror={pub} />
      ) : (
        <DynamicReactJson src={pub} />
      )}
    </>
  );
}
