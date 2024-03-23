import * as React from "react";
import Grid from "@mui/material/Grid";
import { Box, Card, CardMedia, Typography } from "@mui/material";
import Link from "next/link";
import { useEffect, useState } from "react";

// const ItemMyList = () => {

interface Item {
  id: number;
  text_name: string;
  class_name: string;
  image_name: string;
  price: number;
  sell_type: number;
}

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:9000";
const placeholderImage = process.env.PUBLIC_URL + "/logo192.png";

interface Prop {
  reload?: boolean;
  onLoadCompleted?: () => void;
}

export const ItemMyList: React.FC<Prop> = (props) => {
  const { reload = true, onLoadCompleted } = props;
  const [items, setItems] = useState<Item[]>([]);
  const fetchItems = () => {
    fetch(server.concat("/items"), {
      method: "GET",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("GET success:", data);
        setItems(data.items);
        onLoadCompleted && onLoadCompleted();
      })
      .catch((error) => {
        console.error("GET error:", error);
      });
  };

  useEffect(() => {
    if (reload) {
      fetchItems();
    }
  }, [reload]);

  console.log(items);

  return (
    <Grid container spacing={1} py={2}>
      {items.map((item, index) => (
        <Grid item xs={4}>
          <Link href={"/ItemDetail"} style={{ textDecoration: "none" }}>
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
              <Typography fontSize={13} height={40}>
                {item.text_name}
              </Typography>
              <Typography>￥{item.price}</Typography>
              <Typography fontSize={9}>{item.sell_type}</Typography>
            </Card>
          </Link>
        </Grid>
      ))}
    </Grid>
  );
};

export default ItemMyList;
