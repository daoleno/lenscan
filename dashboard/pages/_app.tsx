import Layout from "@/components/Layout";
import { LensConfig, LensProvider, production } from "@lens-protocol/react-web";
import { bindings as wagmiBindings } from "@lens-protocol/wagmi";
import PlausibleProvider from "next-plausible";
import type { AppProps } from "next/app";
import "styles/globals.css";
import { WagmiConfig, configureChains, createClient } from "wagmi";
import { polygon } from "wagmi/chains";
import { publicProvider } from "wagmi/providers/public";

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

export default function App({ Component, pageProps }: AppProps) {
  return (
    <PlausibleProvider domain="lenscan.io">
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
