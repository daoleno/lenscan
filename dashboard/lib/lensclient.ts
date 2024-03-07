import { LensClient, production } from "@lens-protocol/client"

const lensClient = new LensClient({
	environment: production,
})

export default lensClient
