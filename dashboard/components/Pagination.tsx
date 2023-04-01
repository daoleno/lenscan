import { Card } from "@tremor/react";

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
          <button
            onClick={() => onPageChange(1)}
            className={`${
              currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
            } inline-flex h-8 w-8 items-center justify-center rounded border border-gray-100 px-7 text-sm text-gray-700`}
            disabled={currentPage === 1}
          >
            First
          </button>
          <button
            onClick={() => onPageChange(currentPage - 1)}
            className={`${
              currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
            } inline-flex h-8 w-8 items-center justify-center rounded border border-gray-100`}
            disabled={currentPage === 1}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-3 w-3"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                clip-rule="evenodd"
              />
            </svg>{" "}
          </button>{" "}
          <p className="text-sm">
            {currentPage}
            <span className="mx-1">/</span>
            {totalPages}
          </p>
          <button
            onClick={() => onPageChange(currentPage + 1)}
            className={`${
              currentPage === totalPages ? "opacity-50 cursor-not-allowed" : ""
            } inline-flex h-8 w-8 items-center justify-center rounded border border-gray-100`}
            disabled={currentPage === totalPages}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-3 w-3"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                clipRule="evenodd"
              />
            </svg>
          </button>
          <button
            onClick={() => onPageChange(totalPages)}
            className={`${
              currentPage === totalPages ? "opacity-50 cursor-not-allowed" : ""
            } inline-flex h-8 w-8 items-center justify-center rounded border border-gray-100 px-7 text-sm text-gray-700`}
            disabled={currentPage === totalPages}
          >
            Last
          </button>
        </div>
      </div>
    </Card>
  );
}
