import LensClient, { polygon } from "@lens-protocol/client";

const lensClient = new LensClient({
  environment: polygon,
});

export default lensClient;
