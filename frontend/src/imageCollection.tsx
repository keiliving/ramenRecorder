import React, { useEffect } from "react";

// レンダリング時 GET /images
// Image コンポーネント作成し、 GET /image?name="hoge" する。

const ImageCollection: React.FC = () => {
  useEffect(() => {
    a();
    console.log("render");
  });

  const a = async () => {
    const res = await fetch("/images");
    console.log(res);
  };
  return (
    <div className="w-11/12">
      <div className="text-center">title</div>
      <div className="flex flex-wrap justify-center">
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
        <div className="m-10 h-32 w-32 bg-red-500"></div>
        <div className="m-10 h-32 w-32 bg-blue-500"></div>
      </div>
    </div>
  );
};

export default ImageCollection;
