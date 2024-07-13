"use client";
import {
  Box,
  Button,
  Card,
  CardMedia,
  Grid,
  InputBase,
  Link,
  Stack,
  Typography,
  styled,
} from "@mui/material";
import React, { useState } from "react";
import SearchIcon from "@mui/icons-material/Search";

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

const SearchComponent = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<Item[]>([]);
  const [searchPerformed, setSearchPerformed] = useState(false);
  // const [item, setItems] = useState<Item>();

  // 検索ハンドラ
  const handleSearchCollege = async (e: any) => {
    e.preventDefault(); // フォームのデフォルトの送信を防ぐ

    // ここでバックエンドのURLを適切に設定する
    // const backendUrl = "http://localhost:8080/search"; // バックエンドのアドレスとポートを指定

    const token = localStorage.getItem("token");
    try {
      const response = await fetch(
        `${server}/search?college=${encodeURIComponent(query)}`,
        {
          // fetch(`${server}/search`, {
          method: "GET", // GETリクエストを使用
          headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
            Authorization: `Bearer ${token}`,
            // 必要に応じて認証ヘッダーを追加
            // 'Authorization': Bearer ${yourAuthToken},
          },
        }
      );

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const data = await response.json();
      console.log("Response data:", data); // レスポンスデータをログ出力
      setResults(Array.isArray(data.items) ? data.items : []); // 検索結果を状態に保存
      setSearchPerformed(true);
    } catch (error) {
      console.error("Error fetching data:", error);
      setResults([]); // エラーが発生した場合、結果をクリア
      setSearchPerformed(true);
    }
  };

  return (
    <div>
      <form onSubmit={handleSearchCollege}>
        <Search>
          <SearchIconWrapper>
            <SearchIcon />
          </SearchIconWrapper>
          <StyledInputBase
            placeholder="何をお探しですか"
            inputProps={{ "aria-label": "search" }}
            value={query}
            onChange={(e: any) => setQuery(e.target.value)}
          />
          <Button
            type="submit"
            variant="outlined"
            sx={{ position: "absolute", right: 0 }}
            disabled={!query.trim()}
          >
            Search
          </Button>
        </Search>
      </form>
      <ul>
        {results.length > 0 ? (
          results.map((item, index) => (
            <Grid key={index} item xs={4}>
              <Link
                href={`/item/${item.id}`}
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
                      // sx={{ objectFit: "contain" }}
                      component="img"
                      sx={{
                        height: { xs: 110, sm: 180 },
                        // backgroundImage: `url(${server}/images/${item.image_name})`,
                      }}
                      image={`${server}/images/${item.image_name}`} //ここ
                    />
                  </Box>
                  <Typography
                    fontSize={{ xs: 12, sm: 15 }}
                    height={40}
                    px={0.5}
                  >
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
                      color={
                        item.sell_type === "レンタル" ? "#009C88" : "#1573FF"
                      }
                      style={{ display: "flex", alignItems: "center" }}
                    >
                      {/* {item.sell_type} */}

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
          ))
        ) : searchPerformed ? ( // 検索が行われたが結果がない場合
          <Typography>該当する商品が見つかりませんでした。</Typography>
        ) : null}{" "}
        {/* 検索が行われていない場合は何も表示しない */}
      </ul>
    </div>
  );
};

const Search = styled("div")(({ theme }) => ({
  position: "relative",
  marginRight: theme.spacing(2),
  marginLeft: 0,
  marginTop: 18,
  width: "100%",
  height: 36,
  border: `1px solid #ABABAB`,
}));

const SearchIconWrapper = styled("div")(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: "100%",
  position: "absolute",
  pointerEvents: "none",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: "inherit",
  "& .MuiInputBase-input": {
    padding: theme.spacing(1, 1, 1, 0),
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    transition: theme.transitions.create("width"),
    width: "100%",
  },
}));

export default SearchComponent;
