import React, { memo } from "react";

const Node = ({ data }: any) => {
  return (
    <div className="px-4 py-2 rounded-xl bg-secondary border-2 border-dashed border-gray-800">
      <div className="flex">
        <div className="rounded-full w-12 h-12 flex justify-center items-center bg-gray-800" />
        <div className="ml-2">
          <div className="text-base text-center text-white font-bold">
            {data.name}
          </div>
        </div>
      </div>
    </div>
  );
};

export default memo(Node);
