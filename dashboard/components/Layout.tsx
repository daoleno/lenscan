import { WithChildren } from "../types";

import { fontSans } from "@/lib/fonts";
import { cn } from "@/lib/utils";
import Footer from "./footer";
import { SiteHeader } from "./site-header";

export default function Layout({ children }: WithChildren) {
  return (
    <main
      className={cn(
        "relative flex min-h-screen flex-col bg-background pb-7 font-sans antialiased",
        fontSans.variable
      )}
    >
      <SiteHeader />
      <section className="container flex-grow">{children}</section>
      <Footer />
    </main>
  );
}
