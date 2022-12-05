import { fetcher } from "@/lib/fetch";
import useSWRInfinite from "swr/infinite";

export function useProjectPages({ creatorId, limit = 10 }: any = {}) {
  const { data, error, size, ...props } = useSWRInfinite(
    (index, previousPageData) => {
      // reached the end
      if (previousPageData && previousPageData.projects.length === 0)
        return null;

      const searchParams = new URLSearchParams();

      searchParams.set("limit", limit);

      if (creatorId) searchParams.set("by", creatorId);

      if (index !== 0) {
        // using oldest projects createdAt date as cursor
        // We want to fetch projects which has a date that is
        // before (hence the .getTime()) the last project's createdAt
        const before = new Date(
          new Date(
            previousPageData.projects[
              previousPageData.projects.length - 1
            ].createdAt
          ).getTime()
        );

        searchParams.set("before", before.toJSON());
      }

      return `/api/projects?${searchParams.toString()}`;
    },
    fetcher,
    {
      refreshInterval: 10000,
      revalidateAll: false,
    }
  );

  const isLoadingInitialData = !data && !error;
  const isLoading =
    isLoadingInitialData ||
    (size > 0 && data && typeof data[size - 1] === "undefined");

  const isEmpty = data?.[0]?.length === 0;

  return {
    data,
    error,
    size,
    isLoading,
    isEmpty,
    ...props,
  };
}
