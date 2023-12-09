"use client"

import { FC, ReactNode } from "react"
import Link from "next/link"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { dateRange } from "@/app/api/analystics/utils"

import { Button } from "../ui/button"

// ChartCard component
interface AnalyticsProps {
  chartTitle: string
  range: string
  setRange: (range: string) => void
  children: ReactNode
}

export const ChartCard: FC<AnalyticsProps> = ({
  chartTitle,
  range,
  setRange,
  children,
}) => {
  return (
    <Card className="m-7">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <Link href="/analytic/users">{chartTitle}</Link>
          <div className="flex justify-end">
            {Object.keys(dateRange).map((key) => (
              <Button
                size="sm"
                variant={"ghost"}
                key={key}
                className={key === range ? "font-bold" : "font-normal"}
                onClick={() => setRange(key)}
              >
                {key}
              </Button>
            ))}
          </div>
        </CardTitle>
      </CardHeader>
      <CardContent>{children}</CardContent>
    </Card>
  )
}
