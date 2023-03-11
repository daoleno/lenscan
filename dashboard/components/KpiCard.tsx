import {
  BadgeDelta,
  Block,
  Card,
  Flex,
  Metric,
  ProgressBar,
  Text,
} from "@tremor/react";

// Single KPI card in the demo dashboard with sample inputs
export default function KpiCard() {
  return (
    <Card maxWidth="max-w-lg">
      <Flex alignItems="items-start">
        <Block>
          <Text>Sales</Text>
          <Metric>$ 12,699</Metric>
        </Block>
        <BadgeDelta deltaType="moderateIncrease" text="13.2%" />
      </Flex>
      <Flex marginTop="mt-4">
        <Text truncate={true}>68% ($ 149,940)</Text>
        <Text> $ 220,500 </Text>
      </Flex>
      <ProgressBar percentageValue={15.9} marginTop="mt-2" />
    </Card>
  );
}
