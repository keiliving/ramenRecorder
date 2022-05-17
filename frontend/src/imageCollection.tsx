import React, { useEffect, useState } from "react";
import Image from "./image";

// レンダリング時 GET /images
// Image コンポーネント作成し、 GET /image?name="hoge" する。

const ImageCollection: React.FC = () => {
  const [imageNames, setImageNames] = useState<string[]>([]);
  useEffect(() => {
    (async function () {
      const res = await fetch("/images");
      setImageNames(await res.json());
    })();
  }, []);

  return (
    <div className="w-11/12">
      <div className="text-center">title</div>
      <div className="flex flex-wrap justify-center">
        {imageNames.map((imageName, i) => (
          <div key={i}>aaa</div>
        ))}
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
        <Image />
      </div>
    </div>
  );
};

export default ImageCollection;
