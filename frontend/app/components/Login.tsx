// import {
//   Avatar,
//   Box,
//   Button,
//   Checkbox,
//   FormControlLabel,
//   Grid,
//   Link,
//   Paper,
//   TextField,
//   Typography,
// } from "@mui/material";
// import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
// import { teal } from "@mui/material/colors";

// export const Login = () => {
//   return (
//     <Paper
//       elevation={3}
//       sx={{
//         p: 4,
//         height: { xs: "50vh", sm: "70vh" },
//         width: "280px",
//         m: "20px auto",
//       }}
//     >
//       <Grid
//         container
//         direction="column"
//         justifyContent="flex-start"
//         alignItems="center"
//       >
//         <Avatar sx={{ bgcolor: teal[400] }}>
//           <LockOutlinedIcon />
//         </Avatar>
//         <Typography variant={"h5"} sx={{ m: "30px" }}>
//           Sign In
//         </Typography>
//       </Grid>
//       <TextField label="Username" variant="outlined" fullWidth required />
//       <TextField
//         type="password"
//         label="Password"
//         variant="outlined"
//         fullWidth
//         required
//         sx={{ mt: 2 }}
//       />
//       {/* ラベルとチェックボックス */}
//       {/* <FormControlLabel
//         labelPlacement="end"
//         label="パスワードを忘れました"
//         control={<Checkbox name="checkboxA" size="small" color="primary" />}
//       /> */}
//       <Box mt={5}>
//         <Button type="submit" color="primary" variant="contained" fullWidth>
//           サインイン
//         </Button>

//         {/* <Typography variant="caption">
//             <Link href="#">パスワードを忘れましたか？</Link>
//           </Typography> */}
//         <Typography variant="caption" display="block" mt={1}>
//           アカウントを持っていますか？
//           <Link href="#">アカウントを作成</Link>
//         </Typography>
//       </Box>
//     </Paper>
//   );
// };
