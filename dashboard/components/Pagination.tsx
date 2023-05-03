import { Card } from "@tremor/react";
import { Button } from "./ui/button";

interface PaginationProps {
  curCursor: number | null;
  nextCursor: number | null;
  totalResults: number;
  resultsPerPage: number;
  setCursor: any;
}

export default function Pagination({
  curCursor,
  nextCursor,
  totalResults,
  resultsPerPage,
  setCursor,
}: PaginationProps) {
  return (
    <Card className="mt-3">
      <div className="flex flex-1 items-center justify-between">
        <div>
          <p className="text-sm text-gray-700">
            Total <span className="font-medium">{totalResults}</span> results
          </p>
        </div>
        <div className="inline-flex items-center justify-center gap-3">
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() =>
              setCursor(curCursor ? curCursor - resultsPerPage : 0)
            }
            disabled={curCursor ? false : true}
          >
            Prev
          </Button>
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() => setCursor(nextCursor)}
            disabled={nextCursor ? false : true}
          >
            Next
          </Button>
        </div>
      </div>
    </Card>
  );
}
