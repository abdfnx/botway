export const fetcher = async (...args) => {
  return fetch(...args).then(async (res) => {
    let payload;
    let err;

    try {
      if (res.status === 204) return null; // 204 does not have body

      payload = await res.json();
    } catch (e) {
      err = e;
    }

    if (res.ok) {
      return payload;
    } else {
      return Promise.reject(
        payload || new Error("Something went wrong " + err),
      );
    }
  });
};
