import React, { useEffect, useState } from "react";
import Image from "./image";

// レンダリング時 GET /images
// Image コンポーネント作成し、 GET /image?name="hoge" する。

const ImageCollection: React.FC = () => {
  const [imageAttrs, setImageAttrs] = useState<string[]>([]);
  useEffect(() => {
    (async function () {
      const res = await fetch("/images");
      const json = await res.json();
      setImageAttrs(json);
      console.log(json);
    })();
  }, []);

  return (
    <div className="w-11/12">
      <div className="text-center">title</div>
      <div className="flex flex-wrap justify-center">
        {imageAttrs.map((imageAttr) => (
          <Image name={imageAttr} key={imageAttr} />
        ))}
        {/* <Image />
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
        <Image /> */}
      </div>
    </div>
  );
};

export default ImageCollection;
