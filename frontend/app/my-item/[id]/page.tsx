"use client";
import ItemDetail from "@/app/components/ItemDetail";
import { useParams } from "next/navigation";
import React from "react";

const ItemMyListDetailPage: React.FC = () => {
  const params = useParams(); //useParamsは動的なルートパラメータ([id]など)を
  const { id } = params;

  if (!id) {
    return <div>Loading...</div>;
  }

  return (
    <ItemDetail
      params={{
        id: id as string, // idを文字列として渡す
      }}
      showEditAndDeleteButtons={true}
    />
  );
};

export default ItemMyListDetailPage;
