import { fetcher } from "@/lib/fetch";
import useSWR from "swr";

export const useCurrentUser = () => {
  return useSWR("/api/user", fetcher);
};

export const useUser = (id: any) => {
  return useSWR(`/api/users/${id}`, fetcher);
};
