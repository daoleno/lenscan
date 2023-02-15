import { Block, Card, ColGrid, Tab, TabList, Text, Title } from "@tremor/react";

import { useState } from "react";

export default function Home() {
  const [selectedView, setSelectedView] = useState(1);
  return (
    <main>
      <Title>Dashboard</Title>
      <Text>Lorem ipsum dolor sit amet, consetetur sadipscing elitr.</Text>

      <TabList
        defaultValue={1}
        handleSelect={(value) => setSelectedView(value)}
        marginTop="mt-6"
      >
        <Tab value={1} text="Page 1" />
        <Tab value={2} text="Page 2" />
      </TabList>

      {selectedView === 1 ? (
        <>
          <ColGrid
            numColsMd={2}
            numColsLg={3}
            gapX="gap-x-6"
            gapY="gap-y-6"
            marginTop="mt-6"
          >
            <Card>
              {/* Placeholder to set height */}
              <div className="h-28" />
            </Card>
            <Card>
              {/* Placeholder to set height */}
              <div className="h-28" />
            </Card>
            <Card>
              {/* Placeholder to set height */}
              <div className="h-28" />
            </Card>
          </ColGrid>

          <Block marginTop="mt-6">
            <Card>
              <div className="h-80" />
            </Card>
          </Block>
        </>
      ) : (
        <Block marginTop="mt-6">
          <Card>
            <div className="h-96" />
          </Card>
        </Block>
      )}
    </main>
  );
}
