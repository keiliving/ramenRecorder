import React from "react";

export const UploadForm: React.FC = () => {
  return (
    <>
      <div>ファイル</div>
      <input
        type="file"
        className="border-2 bg-slate-400"
        onChange={hundleUpload}
      />
    </>
  );
};

const hundleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
  const files = e.currentTarget.files;
  if (files == null) return;
  const file = files[0];
  const res = await fetch("/health", { method: "POST", body: file });
  console.log(res);
};
