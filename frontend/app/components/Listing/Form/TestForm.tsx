// import { Box, Button, MenuItem, TextField } from "@mui/material";
// import React, { useState } from "react";

// const server = process.env.NEXT_PUBLIC_API_URLL || "http://127.0.0.1:9000";

// interface Prop {
//   onListingCompleted?: () => void;
// }

// type formDataType = {
//   image_name: string | File;
//   text_name: string;
//   class_name: string;
//   price: number;
//   sell_type: number;
// };

// export const Listing: React.FC<Prop> = (props) => {
//   const { onListingCompleted } = props;
//   const initialState = {
//     image_name: "",
//     text_name: "",
//     class_name: "",
//     price: 0,
//     sell_type: 0,
//   };
//   const [values, setValues] = useState<formDataType>(initialState);

//   const onValueChange = (event: React.ChangeEvent<HTMLInputElement>) => {
//     setValues({
//       ...values,
//       [event.target.name]: event.target.value,
//     });
//   };
//   const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
//     setValues({
//       ...values,
//       [event.target.name]: event.target.files![0],
//     });
//   };
//   const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
//     event.preventDefault();
//     const data = new FormData();
//     data.append("image_name", values.image_name);
//     data.append("text_name", values.text_name);
//     data.append("class_name", values.class_name);
//     data.append("price", values.price.toString());
//     data.append("sell_type", values.sell_type.toString());

//     fetch(server.concat("/items"), {
//       //itemsのリンクで入力フォームが出るはず！！！
//       method: "POST",
//       mode: "cors",
//       body: data,
//     })
//       .then((response) => {
//         console.log("POST status:", response.statusText);
//         onListingCompleted && onListingCompleted();
//       })
//       .catch((error) => {
//         console.error("POST error:", error);
//       });
//   };

//   const currencies = [
//     {
//       value: "rental",
//       label: "レンタル",
//     },
//     {
//       value: "sale",
//       label: "売り出し",
//     },
//   ];

//   return (
//     <div className="Listing">
//       <form onSubmit={onSubmit}>
//         <Box mt={5}>
//           <TextField
//             type="file"
//             // sx={{ display: "none" }}
//             name="image_name"
//             id="image_name"
//             onChange={onFileChange}
//             required
//             fullWidth
//             // label="写真"
//             variant="outlined"
//             inputProps={{
//               multiple: true,
//             }}
//             // inputProps={{
//             //   multiple: true,
//             // }}
//             // variant="outlined"
//           />
//           <TextField
//             type="text"
//             name="text_name"
//             id="text_name"
//             onChange={onValueChange}
//             required
//             fullWidth
//             label="教科書名"
//             sx={{ mb: 3, mt: 3 }}
//           />
//           <TextField
//             type="text"
//             name="class_name"
//             id="class_name"
//             onChange={onValueChange}
//             required
//             fullWidth
//             label="講義名"
//             sx={{ mb: 3 }}
//           />
//           <TextField
//             type="text"
//             name="price"
//             id="price"
//             onChange={onValueChange}
//             required
//             fullWidth
//             label="価格"
//             sx={{ mb: 3 }}
//           />
//           <TextField
//             select
//             type="text"
//             name="sell_type"
//             id="sell_type"
//             onChange={onValueChange}
//             required
//             fullWidth
//             label="出品タイプ"
//             helperText="選択してください"
//             sx={{ mb: 3 }}
//           >
//             {currencies.map((option) => (
//               <MenuItem key={option.value} value={option.value}>
//                 {option.label}
//               </MenuItem>
//             ))}
//           </TextField>
//           <Box textAlign="center">
//             <Button type="submit" variant="outlined">
//               List this item
//             </Button>
//           </Box>
//         </Box>
//       </form>
//     </div>
//   );
// };

// // const StyledTextField = styled(TextField)(({ theme }) => ({
// //   width: "100%",
// //   height: 500,
// //   zIndex: 0,
// //   [theme.breakpoints.down("sm")]: {
// //     height: 400,
// //   },
// // }));
