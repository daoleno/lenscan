import dynamic from "next/dynamic";

const DynamicReactJson = dynamic(import("react-json-view"), {
  ssr: false,
});

export default DynamicReactJson;
