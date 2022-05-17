import React, { useEffect } from "react";

const Image: React.FC<{ name: string }> = ({ name }) => {
  useEffect(() => {
    (async function () {
      const res = await fetch(`/image?name=${name}`);
      console.log(res);
    })();

    return () => {
      // 現状ない
      console.log("アンマウントされた");
    };
  }, []);

  return (
    <span className="flex flex-col items-center">
      <div className="text-center">{name}</div>
      {/* サムネイルに置き換える。fech するからこれもコンポーネント化した方がよさそう
      <image></image> */}
      <div
        className="m-10 h-32 w-32 bg-red-500"
        onClick={() => console.log("clicked")}
      ></div>
    </span>
  );
};

export default Image;
