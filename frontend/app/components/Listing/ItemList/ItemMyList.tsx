// ItemMyList.tsx

"use client";
import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Typography } from "@mui/material";
import Link from "next/link";
import { Items } from "@/app/data/data";

interface Item {
  id: number;
  text_name: string;
  class_name: string;
  image_name: string;
  price: number;
  sell_type: number;
}

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:3000";

interface Prop {
  reload?: boolean;
  onLoadCompleted?: () => void;
}

export const ItemMyList: React.FC<Prop> = (props) => {
  const { reload = true, onLoadCompleted } = props;
  const [items, setItems] = useState<Item[]>([]);

  const fetchItems = () => {
    fetch(`${server}/items`, {
      // Template literalsを使用
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        // ここで取得したデータがItem[]型であると仮定
        setItems(data); // data.itemsからdataへ変更
        if (onLoadCompleted) {
          onLoadCompleted();
        }
      })
      .catch((error) => {
        console.error("Error fetching data: ", error);
      });
  };

  useEffect(() => {
    if (reload) {
      fetchItems();
    }
  }, [reload]);

  return (
    <Grid container spacing={1.2} py={2}>
      {Items.map((item, index) => (
        <Grid key={index} item xs={4} pb={3}>
          <Link
            href={`/ItemDetail/${item.id}`}
            style={{ textDecoration: "none" }}
          >
            {" "}
            {/* リンクの修正 */}
            <Card
              sx={{
                height: { xs: 200, sm: 270 },
                transition: "0.3s",
                "&:hover": {
                  bgcolor: "#ffffff",
                  opacity: 0.5,
                  transition: "0.3s",
                },
              }}
            >
              <Box bgcolor={"#ededed"}>
                <CardMedia
                  component="img"
                  sx={{ height: { xs: 110, sm: 180 } }}
                  image={item.image_name || "/logo192.png"} // placeholderImageを削除し、デフォルト画像パスを直接記述
                />
              </Box>
              <Typography fontSize={13} height={40}>
                {item.text_name}
              </Typography>
              <Typography>￥{item.price}</Typography>
              <Typography fontSize={9}>{item.sell_type}</Typography>
            </Card>
          </Link>
        </Grid>
      ))}
    </Grid>
  );
};

export default ItemMyList;
