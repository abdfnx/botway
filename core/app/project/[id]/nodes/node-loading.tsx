import React, { memo } from "react";

const NodeLoading = () => {
  return (
    <div className="px-4 py-2 rounded-xl bg-secondary border-2 border-dashed border-gray-800">
      <div className="animate-pulse flex">
        <div className="rounded-full w-12 h-12 flex justify-center items-center bg-bwdefualt" />
        <div className="ml-2 pt-1">
          <div className="h-4 w-20 mb-3 bg-bwdefualt rounded" />
          <div className="h-3 w-10 bg-bwdefualt rounded" />
        </div>
      </div>
    </div>
  );
};

export default memo(NodeLoading);
