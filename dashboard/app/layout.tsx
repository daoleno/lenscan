"use client"

import { LensConfig, LensProvider, production } from "@lens-protocol/react-web"
import { bindings as wagmiBindings } from "@lens-protocol/wagmi"
import PlausibleProvider from "next-plausible"
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
          <PlausibleProvider
            customDomain="https://analytics.lenscan.io"
            domain="lenscan.io"
            selfHosted
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
          </PlausibleProvider>
        </ThemeProvider>
      </body>
    </html>
  )
}
