import * as React from "react";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import { Typography } from "@mui/material";

export default function SearchButton() {
  return (
    <Button
      href="/college-search"
      // onClick={() => setIsLogin(!isLogin)} //ログインしたら出品ボタンが出るように！！！！！が難しい
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
        大学検索
      </Typography>
    </Button>
  );
}
