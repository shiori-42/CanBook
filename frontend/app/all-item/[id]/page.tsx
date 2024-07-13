"use client";
import ItemDetail from "@/app/components/ItemDetail";
import { useParams } from "next/navigation";
import React from "react";

const ItemAllDetailPage: React.FC = () => {
  const params = useParams(); //useParamsは動的なルートパラメータ([id]など)を取得するのに使う
  const { id } = params;

  if (!id) {
    return <div>Loading...</div>;
  }

  return (
    <ItemDetail
      params={{
        id: id as string, // idを文字列として渡す
      }}
    />
  );
};

export default ItemAllDetailPage;
