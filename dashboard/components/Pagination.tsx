import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/20/solid";
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
  const pageNumbers = Array.from({ length: totalPages }, (_, i) => i + 1);

  const handlePreviousClick = () => {
    if (currentPage > 1) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNextClick = () => {
    if (currentPage < totalPages) {
      onPageChange(currentPage + 1);
    }
  };

  const handlePageClick = (pageNumber: number) => {
    onPageChange(pageNumber);
  };

  return (
    <Card>
      <div className="flex flex-1 justify-between sm:hidden">
        <button
          onClick={handlePreviousClick}
          className={`${
            currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
          } relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50`}
          disabled={currentPage === 1}
        >
          Previous
        </button>
        <button
          onClick={handleNextClick}
          className={`${
            currentPage === totalPages ? "opacity-50 cursor-not-allowed" : ""
          } relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50`}
          disabled={currentPage === totalPages}
        >
          Next
        </button>
      </div>

      <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
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

        <div>
          <nav
            className="isolate inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
          >
            <button
              onClick={handlePreviousClick}
              className={`${
                currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
              } relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0`}
              disabled={currentPage === 1}
            >
              <span className="sr-only">Previous</span>
              <ChevronLeftIcon className="h-5 w-5" aria-hidden="true" />
            </button>

            {pageNumbers.map((pageNumber) => (
              <button
                key={pageNumber}
                onClick={() => handlePageClick(pageNumber)}
                className={`${
                  pageNumber === currentPage
                    ? "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    : "text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0"
                } relative inline-flex items-center ${
                  pageNumber >= currentPage + 3 || pageNumber <= currentPage - 3
                    ? "hidden md:inline-flex"
                    : "md:inline-flex"
                } px-4 py-2 text-sm font-semibold`}
              >
                {pageNumber}
              </button>
            ))}

            <button
              onClick={handleNextClick}
              className={`${
                currentPage === totalPages
                  ? "opacity-50 cursor-not-allowed"
                  : ""
              } relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0`}
              disabled={currentPage === totalPages}
            >
              <span className="sr-only">Next</span>
              <ChevronRightIcon className="h-5 w-5" aria-hidden="true" />
            </button>
          </nav>
        </div>
      </div>
    </Card>
  );
}
