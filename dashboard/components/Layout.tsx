import { WithChildren } from "../types";

import { fontSans } from "@/lib/fonts";
import { cn } from "@/lib/utils";
import Footer from "./footer";
import { SiteHeader } from "./site-header";

export default function Layout({ children }: WithChildren) {
  return (
    <main
      className={cn(
        "min-h-screen bg-background font-sans antialiased relative pb-10",
        fontSans.variable
      )}
    >
      <SiteHeader />
      <section className="container">{children}</section>
      <Footer />
    </main>
  );
}
