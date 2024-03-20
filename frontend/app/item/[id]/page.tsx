"use client";
import { items } from "@/app/data/data";
import NotFound from "@/app/not-found";
import { Box, Typography, styled } from "@mui/material";
import React, { Fragment } from "react";

const ItemDetail = async ({ params }: { params: { id: string } }) => {
  const item = items.find((itemdata) => itemdata.id === parseInt(params.id));
  // const member = await getUserData(params.id);

  if (!item) {
    return <NotFound />;
  }

  const backgroundImageStyle = {
    backgroundImage: `url(${item.imagepath})`,
    backgroundSize: "cover",
    backgroundPosition: "center",
    width: "100%",
    height: 500,
  };

  return (
    <>
      <Box
        width={"100%"}
        height={"100%"}
        borderBottom={1}
        borderColor={"divider"}
        pb={3}
      >
        <Box px={{ sm: 10 }}>
          <Box width={"100%"} sx={backgroundImageStyle}></Box>
        </Box>
        <Box mx={{ sm: 10 }}>
          <Typography fontSize={20} fontWeight={"bold"}>
            {item.name}
          </Typography>
          <Typography fontSize={20}>￥{item.price}</Typography>
          <Typography fontSize={12}>{item.sell}</Typography>

          <Typography fontSize={15} fontWeight={"bold"} mt={1}>
            商品の詳細
          </Typography>
        </Box>
      </Box>
      {/* <ItemInfo /> */}
    </>
  );
};

// const StyledBox = styled(Box)(({ theme }) => ({
//   position: "relative",
//   overflow: "hidden",
//   backgroundSize: "cover",
//   backgroundPosition: "center",
//   backgroundImage: `url("/text-o-chem.jpg")`,
//   // textAlign: "center",
//   width: "100%",
//   height: 500,
//   zIndex: 0,
//   [theme.breakpoints.down("sm")]: {
//     height: 400,
//   },
// }));

export default ItemDetail;
