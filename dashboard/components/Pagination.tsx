import { Card } from "@tremor/react";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { Button } from "./ui/button";

interface PaginationProps {
  currentPage: number;
  totalResults: number;
  resultsPerPage: number;
  onPageChange: (pageNumber: number) => void;
}

export default function Pagination({
  currentPage,
  totalResults,
  resultsPerPage,
  onPageChange,
}: PaginationProps) {
  const totalPages = Math.ceil(totalResults / resultsPerPage);

  return (
    <Card>
      <div className="flex flex-1 items-center justify-between">
        <div>
          <p className="text-sm text-gray-700">
            Showing{" "}
            <span className="font-medium">
              {(currentPage - 1) * resultsPerPage + 1}
            </span>{" "}
            to{" "}
            <span className="font-medium">
              {Math.min(currentPage * resultsPerPage, totalResults)}
            </span>{" "}
            of <span className="font-medium">{totalResults}</span> results
          </p>
        </div>
        <div className="inline-flex items-center justify-center gap-3">
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() => onPageChange(1)}
            disabled={currentPage === 1}
          >
            First
          </Button>
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() => onPageChange(currentPage - 1)}
            disabled={currentPage === 1}
          >
            <ChevronLeft />
          </Button>
          <p className="text-sm">
            <span className="font-medium">{currentPage}</span>
            <span className="mx-1">/</span>
            {totalPages}
          </p>
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() => onPageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
          >
            <ChevronRight />
          </Button>
          <Button
            variant={"outline"}
            size={"xs"}
            onClick={() => onPageChange(totalPages)}
            disabled={currentPage === totalPages}
          >
            Last
          </Button>
        </div>
      </div>
    </Card>
  );
}
