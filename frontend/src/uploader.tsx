import React from "react";

const UploadForm: React.FC = () => {
  return (
    <div className="flex space-x-6">
      <div className="">keitaro-m のラーメン記録</div>
      <input
        type="file"
        className="border-2 bg-slate-400"
        onChange={hundleUpload}
      />
    </div>
  );
};

const hundleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
  const formData = new FormData();
  const files = e.currentTarget.files;
  if (files == null) return;
  const file = files[0];
  formData.append("file", file);
  const res = await fetch("/upload", {
    method: "POST",
    body: formData,
  });
  console.log(res);
  const res2 = await fetch("/images");
  console.log(res2);
};

export default UploadForm;
