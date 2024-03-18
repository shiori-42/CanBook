"use client";
import {
  Box,
  Typography,
  styled,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import React, { Fragment } from "react";
import ItemInfo from "./ItemInfo";

const Subheader = () => {
  return (
    <>
      <Box
        width={"100%"}
        height={"100%"}
        borderBottom={1}
        borderColor={"divider"}
        pb={3}
      >
        <StyledBox />
        <Typography fontSize={20} fontWeight={"bold"}>
          マクマリー有機化学概説 第7版
        </Typography>
        <Typography fontSize={20}>￥900</Typography>
        <Typography fontSize={12}>レンタル</Typography>
      </Box>
      <Typography fontSize={15} fontWeight={"bold"} mt={1}>
        商品の詳細
      </Typography>
      <ItemInfo />
    </>
  );
};

const StyledBox = styled(Box)(({ theme }) => ({
  position: "relative",
  overflow: "hidden",
  backgroundSize: "cover",
  backgroundPosition: "center",
  backgroundImage: `url("/text-o-chem.jpg")`,
  // textAlign: "center",
  width: "100%",
  height: 500,
  zIndex: 0,
  [theme.breakpoints.down("sm")]: {
    height: 400,
  },
}));

export default Subheader;
