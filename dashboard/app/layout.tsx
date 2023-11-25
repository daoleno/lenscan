"use client"

import client from "@/apollo"
import Layout from "@/components/Layout"
import { ApolloProvider } from "@apollo/client"
import { LensConfig, LensProvider, production } from "@lens-protocol/react-web"
import { bindings as wagmiBindings } from "@lens-protocol/wagmi"
import PlausibleProvider from "next-plausible"
import { createPublicClient, http } from "viem"
import { polygon } from "viem/chains"
import { WagmiConfig, createConfig } from "wagmi"

import "styles/globals.css"


const config = createConfig({
  autoConnect: true,
  publicClient: createPublicClient({
    chain: polygon,
    transport: http(),
  }),
})

const lensConfig: LensConfig = {
  bindings: wagmiBindings(),
  environment: production,
}

export default function RootLayout({
  // Layouts must accept a children prop.
  // This will be populated with nested layouts or pages
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <ApolloProvider client={client}>
          <PlausibleProvider
            customDomain="https://analytics.lenscan.io"
            domain="lenscan.io"
            selfHosted
          >
            <WagmiConfig config={config}>
              <LensProvider config={lensConfig}>
                <Layout>{children}</Layout>
              </LensProvider>
            </WagmiConfig>
          </PlausibleProvider>
        </ApolloProvider>
      </body>
    </html>
  )
}
