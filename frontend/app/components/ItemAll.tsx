import React, { useState } from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Stack, Typography } from "@mui/material";
import Link from "next/link";
import { Items } from "../data/data";
import SearchBox from "./SearchBox";

const ItemAll = () => {
  const [filteredItems, setFilteredItems] = useState(Items);
  const [isLogin, setIsLogin] = useState(true);

  // const handleSearch = (searchQuery: string) => {
  //   const filtered = items.filter((item) =>
  //     item.name.toLowerCase().includes(searchQuery.toLowerCase())
  //   );
  //   setFilteredItems(filtered);
  // };//絞り込みしてる

  //ItemMyListのようにGETメソッド書く！

  return (
    <>
      <Grid container spacing={1} py={2}>
        {filteredItems.map((item, index) => (
          <Grid key={index} item xs={4}>
            <Link href={`/item/${item.id}`} style={{ textDecoration: "none" }}>
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
                    image={item.image_name}
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
                    color={
                      item.sell_type === "レンタル" ? "#009C88" : "#1573FF"
                    }
                    style={{ display: "flex", alignItems: "center" }}
                  >
                    {item.sell_type}
                  </Typography>
                  <Typography
                    fontSize={{ xs: 16, sm: 20 }}
                    alignItems={"center"}
                    justifyContent={"center"}
                    color={
                      item.sell_type === "レンタル" ? "#009C88" : "#1573FF"
                    }
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
