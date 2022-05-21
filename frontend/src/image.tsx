import React, { useEffect, useState } from "react";

const Image: React.FC<{ name: string }> = ({ name }) => {
  const [objectURL, setobjectURL] = useState<string>();
  useEffect(() => {
    (async function () {
      const res = await fetch(`/image?name=${name}`);
      if (res.body == null) {
        return;
      }
      const reader = res.body.getReader();
      const stream = new ReadableStream({
        start(controller) {
          return (function pump(): void | PromiseLike<void> {
            return reader.read().then(({ done, value }) => {
              if (done) return controller.close();
              controller.enqueue(value);
              return pump();
            });
          })();
        },
      });
      const data = await new Response(stream).blob();
      const objectURL = URL.createObjectURL(data);
      setobjectURL(objectURL);
    })();

    return () => {
      if (objectURL) URL.revokeObjectURL(objectURL);
    };
  }, []);

  return (
    <span className="pb-1/4 relative m-1 flex w-1/4 max-w-xs flex-col items-center">
      <img
        className="absolute h-full w-full object-cover"
        src={objectURL}
      ></img>
    </span>
  );
};

export default Image;
