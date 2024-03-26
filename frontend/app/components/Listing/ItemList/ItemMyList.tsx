// ItemMyList.tsx

"use client";
import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Stack, Typography } from "@mui/material";
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
          <Link href={`/item/${item.id}`} style={{ textDecoration: "none" }}>
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
                  // image={item.image_name}
                  image={item.image_name || "/logo192.png"} // placeholderImageを削除し、デフォルト画像パスを直接記述　ん？？？？？？？？？？
                />
              </Box>
              <Typography fontSize={{ xs: 12, sm: 15 }} height={40} px={0.5}>
                {item.text_name}
              </Typography>
              <Stack direction={"row"} mt={1}>
                <Typography
                  fontSize={{ xs: 7, sm: 9 }}
                  mx={0.5}
                  mt={{ xs: 0.2, sm: 0.5 }}
                  fontWeight={"bold"}
                  textAlign={"center"}
                  justifyContent={"center"}
                  width={{ xs: 45, sm: 53 }}
                  height={{ xs: 15, sm: 18 }}
                  border={1.9}
                  borderRadius={"20px"}
                  color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
                  style={{ display: "flex", alignItems: "center" }}
                >
                  {item.sell_type}
                </Typography>
                <Typography
                  fontSize={{ xs: 16, sm: 20 }}
                  alignItems={"center"}
                  justifyContent={"center"}
                  color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
                >
                  ￥{item.price}
                </Typography>
              </Stack>
            </Card>
          </Link>
        </Grid>
      ))}
    </Grid>
  );
};

export default ItemMyList;
