import { createPublicClient, http } from "viem"
import { polygon } from "viem/chains"

const viemclient = createPublicClient({
  chain: polygon,
  transport: http(),
})

export default viemclient
