import axios from "axios";
import { useState } from "react";

import { BUNDLR_SPENT_API } from "@/constants";

const useSubmitterSpent = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchData = async (submitters: string[]) => {
    setLoading(true);
    try {
      const response = await axios.post(BUNDLR_SPENT_API, submitters);
      const { sum } = await response.data;
      return sum;
    } catch (error: any) {
      setError(error);
    } finally {
      setLoading(false);
    }
  };

  return { loading, error, fetchData };
};

export default useSubmitterSpent;
