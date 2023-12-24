"use client"

import { FC, ReactNode } from "react"

import { cn } from "@/lib/utils"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { dateRange } from "@/app/api/analystics/utils"

import { Button } from "../ui/button"

// ChartCard component
interface AnalyticsProps {
  chartTitle: string
  range?: string
  setRange?: (range: string) => void
  children: ReactNode
  className?: string
}

export const ChartCard: FC<AnalyticsProps> = ({
  chartTitle,
  range,
  setRange,
  children,
  className,
}) => {
  return (
    <Card className={cn(className)}>
      <CardHeader>
        <CardTitle className="text-md flex items-center justify-between">
          <div>{chartTitle}</div>
          {setRange && range && (
            <div className="flex justify-end text-sm">
              {Object.keys(dateRange).map((key) => (
                <Button
                  size="sm"
                  variant={"ghost"}
                  key={key}
                  className={cn(
                    "px-1 text-xs font-normal",
                    range === key && "font-bold"
                  )}
                  onClick={() => setRange(key)}
                >
                  {key}
                </Button>
              ))}
            </div>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent>{children}</CardContent>
    </Card>
  )
}
