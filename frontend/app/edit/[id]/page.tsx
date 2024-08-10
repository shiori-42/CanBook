"use client";
import React from "react";
import { useParams } from "next/navigation";
import EditForm from "@/app/_components/EditMyItem";

const EditPage: React.FC = () => {
  const params = useParams();
  const { id } = params;

  return (
    <EditForm
      params={{
        id: id as string,
      }}
      itemId={id as string}
    />
  );
};

export default EditPage;
