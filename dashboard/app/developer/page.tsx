import React from "react"

import LinksGrid from "@/components/links-grid"

const App: React.FC = () => {
  const links = [
    {
      url: "https://docs.lens.xyz/docs/react-hooks-sdk-v2",
      linkText: "React Hooks SDK v2",
    },
    {
      url: "https://lens-protocol.github.io/lens-sdk/modules/_lens_protocol_react_web.html",
      linkText: "@lens-protocol/react-web reference",
    },
    {
      url: "https://lens-protocol.github.io/lens-sdk/modules/_lens_protocol_react.html",
      linkText: "@lens-protocol/react reference",
    },
    {
      url: "https://docs.lens.xyz/docs/lensclient-sdk-1",
      linkText: "LensClient SDK",
    },
  ]

  return (
    <div className="mx-auto p-8">
      <h1 className="text-2xl font-bold">Developer</h1>
      <LinksGrid links={links} />
    </div>
  )
}

export default App
