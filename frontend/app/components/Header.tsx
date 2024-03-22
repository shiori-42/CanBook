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
import React from "react";

// export default function ButtonAppBar() {
export default function MenuAppBar() {
  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setAuth(event.target.checked);
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <Box>
      <AppBar color="default" position="static">
        <Toolbar>
          <Typography
            variant="h6"
            component="div"
            fontFamily=""
            letterSpacing={2}
            sx={{ flexGrow: 1 }}
          >
            BookCamp
          </Typography>

          <Button color="inherit" href="/login" sx={{ mr: { sm: 3 } }}>
            <Typography>ログイン</Typography>
          </Button>

          <Button color="inherit" href="/" sx={{ mr: { sm: 3 } }}>
            <Typography>ホーム</Typography>
          </Button>

          {auth && (
            <div>
              <IconButton
                size="large"
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
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

          {/* <IconButton
            size="large"
            color="inherit"
            aria-label="menu"
            sx={{ p: "12px" }}
          >
            <MenuIcon />
          </IconButton> */}
        </Toolbar>
      </AppBar>
    </Box>
  );
}
