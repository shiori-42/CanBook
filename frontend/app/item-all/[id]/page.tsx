"use client";
// import { Items } from "@/app/data/data";
import { Box, Button, Stack, Typography, styled } from "@mui/material";
import React, { useEffect, useState } from "react";

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
  params: { id: string };
}

// const ItemDetail = async ({ params }: { params: { id: string } }) => {
//   const item = Items.find((itemdata) => itemdata.id === parseInt(params.id));

//   if (!item) {
//     return <NotFound />;
//   }

const ItemDetail: React.FC<Prop> = (props) => {
  const { reload = true, onLoadCompleted, params } = props;
  const [item, setItems] = useState<Item>();

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
          const item = data.items.find(
            (itemdata: { id: number }) => itemdata.id === parseInt(params.id)
          );
          setItems(item); // dataからdata.itemsへ変更
        } else {
          console.error("Fetched data is not an array:", data);
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

  if (item == undefined) {
    return null;
  }

  const StyledBox = styled(Box)(({ theme }) => ({
    position: "relative",
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundImage: `url(${server}/images/${item.image_name})`,
    // textAlign: "center",
    width: "100%",
    height: 480,
    zIndex: 0,
    [theme.breakpoints.down("sm")]: {
      height: 400,
    },
  }));

  return (
    <>
      <Box
        width={"100%"}
        height={"100%"}
        borderBottom={1}
        borderColor={"divider"}
        pb={3}
        // sx={backgroundImageStyle}
      >
        <StyledBox />
        {/* <Typography fontSize={20} fontWeight={"bold"}>
          {item.text_name}
        </Typography>
        <Typography fontSize={20}>￥{item.price}</Typography>
        <Typography fontSize={12}>{item.sell_type}</Typography> */}
        <Typography
          fontSize={{ xs: 18, sm: 20 }}
          height={40}
          px={0.5}
          mt={1}
          fontWeight={"bold"}
        >
          {item.name}
        </Typography>
        <Stack direction={"row"} mt={{ xs: 0, sm: 1 }}>
          <Typography
            fontSize={{ xs: 9, sm: 11 }}
            mx={0.5}
            mt={{ xs: 0.7, sm: 0.6 }}
            fontWeight={"bold"}
            textAlign={"center"}
            justifyContent={"center"}
            width={{ xs: 49, sm: 100 }}
            height={{ xs: 15, sm: 23 }}
            border={1.9}
            borderRadius={"20px"}
            color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
            style={{ display: "flex", alignItems: "center" }}
          >
            {item.sell_type}
          </Typography>
          <Typography
            fontSize={{ xs: 20, sm: 25 }}
            width={100}
            alignItems={"center"}
            justifyContent={"center"}
            color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
          >
            ￥{item.price}
          </Typography>
        </Stack>
        <Stack my={2} direction="row" justifyContent="end" spacing={1}>
          <Button
            href="/chat"
            sx={{
              mt: 1,
              border: 1,
              color: "white",
              bgcolor: "#F47381",
              "&:hover": {
                bgcolor: "#ffffff",
                borderColor: "#F47381",
                color: "#F47381",
              },
            }}
          >
            <Typography fontSize={{ xs: 10, sm: 13 }} fontFamily="revert-layer">
              チャットを始める
            </Typography>
          </Button>
        </Stack>
      </Box>
      {/* <Typography fontSize={15} fontWeight={"bold"} mt={1}>
        商品の詳細
      </Typography>
      <ItemInfo /> */}
    </>
  );
};
export default ItemDetail;
