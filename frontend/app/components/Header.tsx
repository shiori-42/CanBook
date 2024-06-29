"use client";
import { AccountCircle } from "@mui/icons-material";
import {
  Box,
  AppBar,
  Toolbar,
  IconButton,
  Typography,
  Button,
  Menu,
  MenuItem,
  Link,
} from "@mui/material";
import React, { useState } from "react";

// export default function ButtonAppBar() {
export default function MenuAppBar() {
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const [isLogin, setIsLogin] = useState(true);

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <Box>
      <AppBar position="static" sx={{ bgcolor: "#ffffff" }}>
        <Toolbar>
          <Typography
            fontSize={23}
            component="div"
            fontFamily="revert-layer"
            letterSpacing={1.5}
            fontWeight={"bold"}
            color={"#009C88"}
            sx={{ flexGrow: 1 }}
            ml={1}
          >
            CanBook
          </Typography>

          <Button
            color="inherit"
            href="/home"
            sx={{
              mr: { sm: 2 },
              border: 1,
              borderColor: "#ffffff",
              color: "#009C88",
              "&:hover": {
                borderColor: "#009C88",
                bgcolor: "white",
              },
            }}
          >
            <Typography fontSize={{ xs: 10, sm: 13 }} fontFamily="revert-layer">
              ホーム
            </Typography>
          </Button>

          <Button
            href="/login"
            sx={{
              mr: { sm: 2 },
              border: 1,
              color: "white",
              bgcolor: "#009C88",
              "&:hover": {
                bgcolor: "#ffffff",
                borderColor: "#009C88",
                color: "#009C88",
              },
            }}
          >
            <Typography fontSize={{ xs: 10, sm: 13 }} fontFamily="revert-layer">
              ログイン
            </Typography>
          </Button>

          <Button
            href="/form"
            onClick={() => setIsLogin(!isLogin)} //ログインしたら出品ボタンが出るようにしたい
            sx={{
              mr: { sm: 2 },
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
              出品
            </Typography>
          </Button>

          {auth && (
            <div>
              <IconButton
                size="large"
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                sx={{ color: "#c6c6c6" }}
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: "top",
                  horizontal: "right",
                }}
                keepMounted
                transformOrigin={{
                  vertical: "top",
                  horizontal: "right",
                }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                <MenuItem role="none" onClick={handleClose}>
                  <Link
                    role="menuitem"
                    href="/profile"
                    sx={{ textDecoration: "none" }}
                    width={"100%"}
                  >
                    <Typography color={""}>Profile</Typography>
                  </Link>
                </MenuItem>
                {/* <MenuItem onClick={handleClose}>My account</MenuItem> */}
                <MenuItem role="none" onClick={handleClose}>
                  <Link
                    role="menuitem"
                    href="/my-account"
                    sx={{ textDecoration: "none" }}
                    width={"100%"}
                  >
                    <Typography color={""}>My account</Typography>
                  </Link>
                </MenuItem>
              </Menu>
            </div>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}
