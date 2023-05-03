import Layout from "@/components/Layout";
import { LensConfig, LensProvider, production } from "@lens-protocol/react-web";
import { bindings as wagmiBindings } from "@lens-protocol/wagmi";
import PlausibleProvider from "next-plausible";
import type { AppProps } from "next/app";
import "styles/globals.css";
import { WagmiConfig, configureChains, createClient } from "wagmi";
import { polygon } from "wagmi/chains";
import { publicProvider } from "wagmi/providers/public";

import { trpc } from "@/lib/trpc";

const { provider, webSocketProvider } = configureChains(
  [polygon],
  [publicProvider()]
);

const client = createClient({
  autoConnect: true,
  provider,
  webSocketProvider,
});

const lensConfig: LensConfig = {
  bindings: wagmiBindings(),
  environment: production,
};

export function App({ Component, pageProps }: AppProps) {
  return (
    <PlausibleProvider
      customDomain="https://analytics.lenscan.io"
      domain="lenscan.io"
      selfHosted
    >
      <WagmiConfig client={client}>
        <LensProvider config={lensConfig}>
          <Layout>
            <Component {...pageProps} />
          </Layout>
        </LensProvider>
      </WagmiConfig>
    </PlausibleProvider>
  );
}

export default trpc.withTRPC(App);
