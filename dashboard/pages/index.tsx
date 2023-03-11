import ChartView from "@/components/ChartView";
import KpiCardGrid from "@/components/KpiCardGrid";
import TableView from "@/components/TableView";
import { Block, Flex, Tab, TabList, Text, Title } from "@tremor/react";
import { useState } from "react";

export default function Home() {
  const [selectedView, setSelectedView] = useState(1);
  return (
    <main className="bg-slate-50 p-6 sm:p-10">
      <Title>Lenscan</Title>
      <Text>Explorer for the Lens Protocol.</Text>

      <TabList
        defaultValue={1}
        handleSelect={(value) => setSelectedView(value)}
        marginTop="mt-6"
      >
        <Tab value={1} text="Overview" />
        <Tab value={2} text="Detail" />
      </TabList>

      {selectedView === 1 ? (
        <>
          <KpiCardGrid />

          <Flex marginTop="mt-6">
            <ChartView />
          </Flex>
        </>
      ) : (
        <Block marginTop="mt-6">
          <TableView />
        </Block>
      )}
    </main>
  );
}
