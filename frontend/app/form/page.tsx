"use client";
import React, { useState } from "react";
import InputForm from "../_components/InputForm";

const Page = () => {
  const [reload, setReload] = useState(true);
  return <InputForm onListingCompleted={() => setReload(true)} />;
};

export default Page;
