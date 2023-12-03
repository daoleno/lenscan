import lensClient from "@/lib/lensclient"

import Comment from "./comment"
import Mirror from "./mirror"
import Post from "./post"

export default async function Publication({ id }: { id: string }) {
  const pub = await lensClient.publication.fetch({
    forId: id,
  })

  return (
    <>
      {pub?.__typename === "Post" ? (
        <Post post={pub} />
      ) : pub?.__typename === "Comment" ? (
        <Comment comment={pub} />
      ) : pub?.__typename === "Mirror" ? (
        <Mirror mirror={pub} />
      ) : (
        <div>Unknown publication {pub?.__typename}</div>
      )}
    </>
  )
}
