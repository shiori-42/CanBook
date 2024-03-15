import MenuIcon from "@mui/icons-material/Menu";
import {
  Box,
  AppBar,
  Toolbar,
  IconButton,
  Typography,
  Button,
} from "@mui/material";

export default function ButtonAppBar() {
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
          <Button color="inherit" href="/login">
            ログイン
          </Button>
          <IconButton
            size="large"
            color="inherit"
            aria-label="menu"
            sx={{ p: "12px" }}
          >
            <MenuIcon />
          </IconButton>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
