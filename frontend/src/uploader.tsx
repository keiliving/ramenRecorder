import React from "react";

export const UploadForm: React.FC = () => {
  return (
    <>
      <div>ファイル</div>
      <input
        type="file"
        className="border-2 bg-slate-400"
        onChange={hundleUpload}
        name="test"
      />
    </>
  );
};

const hundleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
  const formData = new FormData();

  const files = e.currentTarget.files;
  if (files == null) return;
  const file = files[0];
  formData.append("userfile", file);
  const res = await fetch("/upload", {
    method: "POST",
    body: formData,
  });
  console.log(res);
};