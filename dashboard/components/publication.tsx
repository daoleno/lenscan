import { usePublication } from "@lens-protocol/react-web";
import NotFound from "./404";
import Comment from "./comment";
import { Loading } from "./loading";
import Post from "./post";
("lucide-react");

export default function Publication({ id }: { id: string }) {
  const { data: pub, loading, error } = usePublication({ publicationId: id });

  console.log({ pub, loading, error });

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
      ) : null}
    </>
  );
}
