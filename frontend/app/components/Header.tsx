import {
  SignedIn,
  UserButton,
  SignedOut,
  SignInButton,
  SignIn,
  RedirectToSignIn,
} from "@clerk/nextjs";
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

          <SignedIn>
            <UserButton afterSignOutUrl="/" />
          </SignedIn>
          <SignedOut>
            <Button color="inherit" href="/my-page">
              ログイン
            </Button>
          </SignedOut>
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
