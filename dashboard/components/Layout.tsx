import { WithChildren } from "../types";

import { fontSans } from "@/lib/fonts";
import { cn } from "@/lib/utils";
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
      <section className="container space-y-8">{children}</section>
    </main>
  );
}
