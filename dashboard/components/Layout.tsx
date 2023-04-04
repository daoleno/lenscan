import { WithChildren } from "../types";
import NavBar from "./Navbar";

export default function Layout({ children }: WithChildren) {
  return (
    <main className="bg-slate-50 p-6 sm:p-10">
      <NavBar />
      <section>{children}</section>
    </main>
  );
}
