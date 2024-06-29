// import * as React from "react";
// import Grid from "@mui/material/Grid";
// import { Box, Card, CardMedia, Typography } from "@mui/material";
// import Link from "next/link";
// import { items } from "../data/data";

// const ItemKeep = () => {
//   return (
//     <Grid container spacing={1} py={2}>
//       {items.map((item, index) => (
//         <Grid key={index} item xs={4}>
//           <Link href={"/ItemDetail"} style={{ textDecoration: "none" }}>
//             <Card
//               sx={{
//                 height: { xs: 200, sm: 270 },
//                 transition: "0.3s",
//                 "&:hover": {
//                   bgcolor: "#ffffff",
//                   opacity: 0.5,
//                   transition: "0.3s",
//                 },
//               }}
//             >
//               <Box bgcolor={"#ededed"}>
//                 <CardMedia
//                   // sx={{ objectFit: "contain" }}
//                   component="img"
//                   sx={{ height: { xs: 110, sm: 180 } }}
//                   image={item.imagepath}
//                 />
//               </Box>
//               <Typography fontSize={13} height={40}>
//                 {item.name}
//               </Typography>
//               <Typography>￥{item.price}</Typography>
//               <Typography fontSize={9}>{item.sell}</Typography>
//             </Card>
//           </Link>
//         </Grid>
//       ))}
//     </Grid>
//   );
// };

// export default ItemKeep;

//保存したリスト
