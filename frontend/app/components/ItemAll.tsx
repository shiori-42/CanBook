import React, { useEffect, useState } from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Stack, Typography } from "@mui/material";
import Link from "next/link";
import { Items } from "../data/data";
import SearchBox from "./SearchButton";

// const handleSearch = (searchQuery: string) => {
//   const filtered = items.filter((item) =>
//     item.name.toLowerCase().includes(searchQuery.toLowerCase())
//   );
//   setFilteredItems(filtered);
// };//絞り込みしてる

//ItemMyListのようにGETメソッド書く！
//ItemMyListが上手くいったら！！！！

interface Item {
  id: number;
  name: string;
  course_name: string;
  price: number;
  sell_type: string;
  image_name: string;
  created_at: string;
  updated_at: string;
  user_id: number;
}

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:3000";

interface Prop {
  reload?: boolean;
  onLoadCompleted?: () => void;
}

// const ItemAll = () => {
export const ItemAll: React.FC<Prop> = (props) => {
  const { reload = true, onLoadCompleted } = props;
  const [items, setItems] = useState<Item[]>([]);

  const fetchItems = () => {
    const token = localStorage.getItem("token");

    fetch(`${server}/alluseritems`, {
      // Template literalsを使用
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",

        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        if (Array.isArray(data.items)) {
          setItems(data.items); // dataからdata.itemsへ変更
        } else {
          console.error("Fetched data is not an array:", data);
          setItems([]); // データが配列でない場合は、空の配列を設定する
        }
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
    //filteredItems.mapだった
    <>
      <Grid container spacing={1} py={2}>
        {items.map((item, index) => (
          <Grid key={index} item xs={4}>
            <Link
              href={`/item-all/${item.id}`}
              style={{ textDecoration: "none" }}
            >
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
                    // sx={{ objectFit: "contain" }}
                    component="img"
                    sx={{ height: { xs: 110, sm: 180 } }}
                    image={`${process.env.NEXT_PUBLIC_API_URL}/images/${item.image_name}`}
                  />
                </Box>
                <Typography fontSize={{ xs: 12, sm: 15 }} height={40} px={0.5}>
                  {item.name}
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
                    color={item.sell_type === "rental" ? "#009C88" : "#1573FF"}
                    style={{ display: "flex", alignItems: "center" }}
                  >
                    {item.sell_type}
                  </Typography>
                  <Typography
                    fontSize={{ xs: 16, sm: 20 }}
                    alignItems={"center"}
                    justifyContent={"center"}
                    color={item.sell_type === "rental" ? "#009C88" : "#1573FF"}
                  >
                    ￥{item.price}
                  </Typography>
                </Stack>
              </Card>
            </Link>
          </Grid>
        ))}
      </Grid>
    </>
  );
};

export default ItemAll;
