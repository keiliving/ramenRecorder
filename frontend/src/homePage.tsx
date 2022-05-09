import React from "react";
import UploadForm from "./uploader";
import ImageCollection from "./imageCollection";

const HomePage: React.FC = () => {
  return (
    <>
      <UploadForm />
      <ImageCollection />
    </>
  );
};

export default HomePage;
