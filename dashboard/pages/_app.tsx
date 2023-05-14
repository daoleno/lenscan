import client from "@/apollo";
import Layout from "@/components/Layout";
import { trpc } from "@/lib/trpc";
import { ApolloProvider } from "@apollo/client";
import { LensConfig, LensProvider, production } from "@lens-protocol/react-web";
import { bindings as wagmiBindings } from "@lens-protocol/wagmi";
import PlausibleProvider from "next-plausible";
import type { AppProps } from "next/app";
import "styles/globals.css";
import { createPublicClient, http } from "viem";
import { polygon } from "viem/chains";
import { WagmiConfig, createConfig } from "wagmi";

const config = createConfig({
  autoConnect: true,
  publicClient: createPublicClient({
    chain: polygon,
    transport: http(),
  }),
});

const lensConfig: LensConfig = {
  bindings: wagmiBindings(),
  environment: production,
};

export function App({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={client}>
      <PlausibleProvider
        customDomain="https://analytics.lenscan.io"
        domain="lenscan.io"
        selfHosted
      >
        <WagmiConfig config={config}>
          <LensProvider config={lensConfig}>
            <Layout>
              <Component {...pageProps} />
            </Layout>
          </LensProvider>
        </WagmiConfig>
      </PlausibleProvider>
    </ApolloProvider>
  );
}

export default trpc.withTRPC(App);
