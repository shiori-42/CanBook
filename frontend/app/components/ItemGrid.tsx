import * as React from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Typography } from "@mui/material";
import { items } from "../data/data";
import Link from "next/link";

const ItemGrid = () => {
  return (
    <Grid container spacing={1} py={2}>
      {items.map((item, index) => (
        <Grid key={index} item xs={4}>
          <Link href={"/aaa"} style={{ textDecoration: "none" }}>
            <Card
              sx={{
                height: 200,
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
                  sx={{ objectFit: "contain" }}
                  component="img"
                  height={110}
                  image={item.imagepath}
                />
              </Box>
              <Typography fontSize={13} height={40}>
                {item.name}
              </Typography>
              <Typography>￥{item.price}</Typography>
              <Typography fontSize={9}>{item.sell}</Typography>
            </Card>
          </Link>
        </Grid>
      ))}
    </Grid>
  );
};

export default ItemGrid;
