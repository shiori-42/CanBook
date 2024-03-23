"use client";
import { Box, Typography, styled } from "@mui/material";
import React, { Fragment } from "react";
import { Items } from "@/app/data/data";
import NotFound from "@/app/not-found";

const ItemDetail = async ({ params }: { params: { id: string } }) => {
  const item = Items.find((itemdata) => itemdata.id === parseInt(params.id));

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
    backgroundImage: `url(${item.image_name})`,
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
          {item.text_name}
        </Typography>
        <Typography fontSize={20}>￥{item.price}</Typography>
        <Typography fontSize={12}>{item.sell_type}</Typography>
      </Box>
      {/* <Typography fontSize={15} fontWeight={"bold"} mt={1}>
        商品の詳細
      </Typography>
      <ItemInfo /> */}
    </>
  );
};

export default ItemDetail;
