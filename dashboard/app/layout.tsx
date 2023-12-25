"use client"

import Script from "next/script"
import { LensConfig, LensProvider, production } from "@lens-protocol/react-web"
import { bindings as wagmiBindings } from "@lens-protocol/wagmi"
import { createPublicClient, http } from "viem"
import { polygon } from "viem/chains"
import { createConfig, WagmiConfig } from "wagmi"

import { fontSans } from "@/lib/fonts"
import { cn } from "@/lib/utils"
import Footer from "@/components/footer"
import { ThemeProvider } from "@/components/providers"
import { SiteHeader } from "@/components/site-header"

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
      <head />
      <body
        className={cn(
          "min-h-screen bg-background font-sans antialiased",
          fontSans.variable
        )}
      >
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <WagmiConfig config={config}>
            <LensProvider config={lensConfig}>
              <div className="relative flex min-h-screen flex-col">
                <SiteHeader />
                <section className="container flex-grow">{children}</section>
                <Footer />
              </div>
            </LensProvider>
          </WagmiConfig>
        </ThemeProvider>

        <Script
          id="goatcounter"
          data-goatcounter="https://lenscan.goatcounter.com/count"
          async
          src="//gc.zgo.at/count.js"
        ></Script>
      </body>
    </html>
  )
}
