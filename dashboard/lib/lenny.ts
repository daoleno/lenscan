import viemclient from "./viem"

const lennyABI = [
  { inputs: [], stateMutability: "nonpayable", type: "constructor" },
  {
    inputs: [],
    name: "blockSeed",
    outputs: [{ internalType: "bytes32", name: "", type: "bytes32" }],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      { internalType: "uint256", name: "profileId", type: "uint256" },
      { internalType: "uint256", name: "mintTimestamp", type: "uint256" },
    ],
    name: "getTokenURI",
    outputs: [{ internalType: "string", name: "", type: "string" }],
    stateMutability: "view",
    type: "function",
  },
]

export async function getLenny(profileId: string) {
  const data = await viemclient.readContract({
    address: "0xCCF77B802160326282F260bb6e275333fEA9E76C",
    abi: lennyABI,
    functionName: "getTokenURI",
    args: [profileId, 1],
  })

  const dataString = data as string
  const decodedString = decodeBase64DataString(dataString)

  return JSON.parse(decodedString)
}

function decodeBase64DataString(dataString: string): string {
  const base64PrefixIndex = dataString.indexOf("base64,") + "base64,".length
  const base64String = dataString.substring(base64PrefixIndex)
  return atob(base64String)
}
