"use client";
import { Box, Typography, styled } from "@mui/material";
import React, { Fragment } from "react";
import ItemInfo from "./ItemInfo";
import NotFound from "../not-found";
import { items } from "../data/data";

const ItemDetail = async ({ params }: { params: { id: string } }) => {
  const item = items.find((itemdata) => itemdata.id === parseInt(params.id));

  if (!item) {
    return <NotFound />;
  }

  // const backgroundImageStyle = {
  //   backgroundImage: `url(${item.imagepath})`,
  //   backgroundSize: "cover",
  //   backgroundPosition: "center",
  //   width: "100%",
  //   height: 500,
  // };

  const StyledBox = styled(Box)(({ theme }) => ({
    position: "relative",
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundImage: `url(${item.imagepath})`,
    // textAlign: "center",
    width: "100%",
    height: 500,
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
        <Typography fontSize={20} fontWeight={"bold"}>
          {item.name}
        </Typography>
        <Typography fontSize={20}>￥{item.price}</Typography>
        <Typography fontSize={12}>{item.sell}</Typography>
      </Box>
      {/* <Typography fontSize={15} fontWeight={"bold"} mt={1}>
        商品の詳細
      </Typography>
      <ItemInfo /> */}
    </>
  );
};

export default ItemDetail;
