import React, { useCallback, useEffect, useState } from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Stack, Typography } from "@mui/material";
import Link from "next/link";

interface Item {
  id: number;
  name: string;
  class_name: string;
  price: number;
  sell_type: string;
  image_name: string;
  created_at: string;
  updated_at: string;
  user_id: number;
}

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:3000";

interface ItemAllProps {
  reload?: boolean;
  onLoadCompleted?: () => void;
}

const ItemAll: React.FC<ItemAllProps> = (props) => {
  const { reload = true, onLoadCompleted } = props;
  const [items, setItems] = useState<Item[]>([]);

  const fetchItems = useCallback(() => {
    const token = localStorage.getItem("token");

    fetch(`${server}/alluseritems`, {
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
          setItems(data.items);
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
  }, [onLoadCompleted]);

  useEffect(() => {
    if (reload) {
      fetchItems();
    }
  }, [reload, fetchItems]);

  return (
    <>
      <Grid container spacing={1} py={2}>
        {items.map((item, index) => (
          <Grid key={index} item xs={4}>
            <Link
              href={`/all-item/${item.id}`}
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
                    component="img"
                    sx={{
                      height: { xs: 110, sm: 180 },
                      objectFit: "contain",
                    }}
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
