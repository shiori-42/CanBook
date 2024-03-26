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
  const [isLogin, setIsLogin] = useState(true);
  const { loginMutation, registerMutation } = useMutateAuth();

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (isLogin) {
      loginMutation.mutate({
        name: name,
        email: email,
        password: pw,
      });
    } else {
      await registerMutation
        .mutateAsync({
          name: name,
          email: email,
          password: pw,
        })
        .then(() =>
          loginMutation.mutate({
            name: name,
            email: email,
            password: pw,
          })
        );
    }
  };

  return (
    <Paper
      elevation={3}
      sx={{
        p: 4,
        height: { xs: "50vh", sm: "72vh" },
        width: "280px",
        m: "15px auto",
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
        <Typography fontSize={20} sx={{ mt: "18px", mb: "28px" }}>
          {isLogin ? "ログイン" : "会員登録"}
        </Typography>
      </Grid>
      <form onSubmit={submitAuthHandler}>
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

        {/* <div> */}
        {/* //ユーザーがnameを入力するためのところ */}
        {/* <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="name"
            type="name"
            autoFocus
            placeholder="name"
            onChange={(e) => setName(e.target.value)}
            value={name}
          /> */}
        {/* </div> */}
        {/* <div> */}
        {/* //ユーザーがe-mailを入力するためのところ */}
        {/* <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="email"
            type="email"
            // autoFocus
            placeholder="Email address"
            onChange={(e) => setEmail(e.target.value)}
            value={email}
          /> */}
        {/* </div> */}
        {/* <div> */}
        {/* //ユーザーがパスワードを入力するためのところ */}
        {/* <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="password"
            type="password"
            placeholder="Password"
            onChange={(e) => setPw(e.target.value)}
            value={pw}
          />
        </div> */}
        {/* <div className="flex justify-center my-2"> */}
        {/* //submitボタンの作成 */}
        {/* <Typography onClick={() => setIsLogin(!isLogin)}>こちら</Typography> */}
        <Box mt={4}>
          <Button
            type="submit"
            color="primary"
            variant="contained"
            fullWidth
            disabled={!name || !email || !pw}
            // href="/profile-form" //ここは普通に飛ばしちゃって大丈夫なのか？？？？？？？？
            //ダメでしたuseMutateAuthにかいてる
          >
            {isLogin ? "ログイン" : "会員登録"}
          </Button>
          {/* <button
            className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
            disabled={!name || !email || !pw}
            type="submit"
          >
            {isLogin ? "Login" : "Sign Up"}
          </button> */}
          {/* </div> */}
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
        {/* <ImportExportIcon
          onClick={() => setIsLogin(!isLogin)}
          // className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
          width={10}
          height={10}
          // sx={{
          //   display: "flex",
          //   alignItems: "center",
          //   justifyContent: "center",
          //   position: "absolute",
          // }} //きかない //一旦コメントアウト一旦一旦
        /> */}
      </Box>
    </Paper>
  );
};

//urlはloginだけでなく、signupのurlも作ってLikタグで飛ばす。また当たらしくsignupのフォルダーを作って、signup/page.tsxを作る。？？？？
//こちら　の部分を onClick={() => setIsLogin(!isLogin)}　にしたら上手く飛ぶかも、、、？
