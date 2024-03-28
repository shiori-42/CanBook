import { useState, FormEvent } from "react";
import { useMutateAuth } from "../hooks/useMutateAuth";
import {
  Avatar,
  Box,
  Button,
  Grid,
  Link,
  Paper,
  TextField,
  Typography,
} from "@mui/material";
import { teal } from "@mui/material/colors";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import ImportExportIcon from "@mui/icons-material/ImportExport";

export const Auth = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [pw, setPw] = useState("");
  const [college, setCollege] = useState("");
  const [campus, setCampus] = useState("");
  const [isLogin, setIsLogin] = useState(true);
  const { loginMutation, registerMutation } = useMutateAuth();

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (isLogin) {
      loginMutation.mutate({
        email: email,
        password: pw,
      });
    } else {
      await registerMutation
        .mutateAsync({
          name: name,
          email: email,
          password: pw,
          college: college,
          campus: campus,
        })
        .then(() =>
          loginMutation.mutate({
            email: email,
            password: pw,
          })
        );
    }
  };

  return (
    <Box
      sx={{
        p: 1,
        height: { xs: "50vh", sm: "72vh" },
        width: "280px",
        m: "10px auto",
        mb: 20,
      }}
    >
      <Grid
        container
        direction="column"
        justifyContent="flex-start"
        alignItems="center"
      >
        <Avatar sx={{ bgcolor: teal[400] }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography fontSize={20} sx={{ mt: "15px", mb: "20px" }}>
          {isLogin ? "ログイン" : "会員登録"}
        </Typography>
      </Grid>

      <form onSubmit={submitAuthHandler}>
        {isLogin ? null : (
          <TextField
            name="name"
            type="name"
            label="Username"
            variant="outlined"
            fullWidth
            required
            autoFocus
            onChange={(e) => setName(e.target.value)}
            value={name}
          />
        )}
        <TextField
          name="email"
          type="email"
          label="Email address"
          variant="outlined"
          fullWidth
          required
          onChange={(e) => setEmail(e.target.value)}
          value={email}
          sx={{ mt: 1.5 }}
        />

        <TextField
          name="password"
          type="password"
          label="Password"
          variant="outlined"
          fullWidth
          required
          onChange={(e) => setPw(e.target.value)}
          value={pw}
          sx={{ mt: 1.5 }}
        />
        {isLogin ? null : (
          <>
            <TextField
              name="college"
              type="college"
              label="College"
              variant="outlined"
              fullWidth
              required
              onChange={(e) => setCollege(e.target.value)}
              value={college}
              sx={{ mt: 1.5 }}
            />
            <TextField
              name="campus"
              type="campus"
              label="Campus"
              variant="outlined"
              fullWidth
              required
              onChange={(e) => setCampus(e.target.value)}
              value={campus}
              sx={{ mt: 1.5 }}
            />
          </>
        )}

        <Box mt={4}>
          <Button
            type="submit"
            color="primary"
            variant="contained"
            fullWidth
            disabled={
              isLogin
                ? !email || !pw
                : !name || !email || !pw || !college || !campus
            }
            // href="/profile-form" //ここは普通に飛ばしちゃって大丈夫なのか？
            //ダメでしたuseMutateAuthにかいてる
          >
            {isLogin ? "ログイン" : "会員登録"}
          </Button>
        </Box>
      </form>
      <Box position={"relative"}>
        <Typography
          variant="caption"
          display="block"
          mt={1}
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          アカウントを持っていますか？
        </Typography>
        <Link onClick={() => setIsLogin(!isLogin)}>
          <Typography
            fontSize={12}
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              cursor: "pointer",
            }}
          >
            {isLogin ? "アカウントを作成" : "ログインをする"}
          </Typography>
        </Link>
      </Box>
    </Box>
  );
};

//urlはloginだけでなく、signupのurlも作ってLikタグで飛ばす。また当たらしくsignupのフォルダーを作って、signup/page.tsxを作る。？？？？
//こちら　の部分を onClick={() => setIsLogin(!isLogin)}　にしたら上手く飛ぶかも、、、？
