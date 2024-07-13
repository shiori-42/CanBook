"use client";
import React, { useState } from "react";
import InputForm from "../components/Listing/Form/InputForm";

const Page = () => {
  //form/page.tsxで記述
  const [reload, setReload] = useState(true);
  return (
    <InputForm onListingCompleted={() => setReload(true)} />

    // {/* <ItemMyList reload={reload} onLoadCompleted={() => setReload(false)} /> */}
  );
};

export default Page;
