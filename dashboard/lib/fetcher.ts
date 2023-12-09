const fetcher = (
  ...args: [input: RequestInfo, init?: RequestInit | undefined]
): Promise<any> => fetch(...args).then((res: Response) => res.json())

export default fetcher
