import React, { useEffect } from "react";

const Image: React.FC = () => {
  useEffect(() => {
    console.log("render img");
  });

  return (
    <div
      className="m-10 h-32 w-32 bg-red-500"
      onClick={() => console.log("clicked")}
    ></div>
  );
};

export default Image;
