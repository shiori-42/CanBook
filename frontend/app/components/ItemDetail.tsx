"use client";
import {
  Box,
  Button,
  IconButton,
  Stack,
  Typography,
  styled,
} from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import React, { useCallback, useEffect, useState } from "react";

export interface Item {
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

interface ItemDetailProps {
  reload?: boolean;
  onLoadCompleted?: () => void;
  params: { id: string };
  showEditAndDeleteButtons?: boolean;
}

const ItemDetail: React.FC<ItemDetailProps> = (props) => {
  const {
    reload = true,
    onLoadCompleted,
    params,
    showEditAndDeleteButtons = false,
  } = props;
  const [item, setItems] = useState<Item | undefined>();

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
          const item = data.items.find(
            (itemdata: { id: number }) => itemdata.id === parseInt(params.id)
          );
          setItems(item);
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
  }, [params.id, onLoadCompleted]);

  useEffect(() => {
    if (reload) {
      fetchItems();
    }
  }, [reload, fetchItems]);

  if (!item) {
    return null;
  }

  const StyledBox = styled(Box)(({ theme }) => ({
    position: "relative",
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundImage: `url(${server}/images/${item.image_name})`,
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
        mt={2}
        pb={3}
      >
        <StyledBox />
        {showEditAndDeleteButtons && (
          <Stack direction="row" spacing={2} justifyContent="end">
            <IconButton sx={{ color: "#F47381" }} href={`/edit/${item.id}`}>
              <EditIcon />
            </IconButton>
            <IconButton
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              sx={{ color: "#F47381" }}
            >
              <DeleteIcon />
            </IconButton>
          </Stack>
        )}
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
